package main

import (
  "bytes"
  "encoding/csv"
  "fmt"
  "flag"
  pathlib "path"
  "html/template"
  "path/filepath"
  "io"
  "io/ioutil"
  "log"
  "os"
  "strings"
  "net/http"
  //urllib "net/url"

  "gopkg.in/russross/blackfriday.v2"
  "github.com/Depado/bfchroma"
  "github.com/alecthomas/chroma/lexers"
  "github.com/alecthomas/chroma/formatters/html"
  "github.com/elazarl/go-bindata-assetfs"
)

var basePath = ""

func init() {
  flag.StringVar(&basePath, "base", basePath, "base path")
}

var errUnidentifiedLanguage = fmt.Errorf("source language cannot be identified")

func highlight(w io.Writer, filename string) error {
  lexer := lexers.Match(filename)
  if lexer == nil {
    return errUnidentifiedLanguage
  }

  formatter := html.New(html.WithClasses())

  f, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer f.Close()

  content, err := ioutil.ReadAll(f)
  if err != nil {
    return err
  }

  iterator, err := lexer.Tokenise(nil, string(content))
  if err != nil {
    return err
  }

  buf := &bytes.Buffer{}
  err = formatter.Format(buf, Style, iterator)
  if err != nil {
    return err
  }

  pageTplBytes := MustAsset("file.html")
  pageTpl := template.Must(template.New("page").Parse(string(pageTplBytes)))
  title := pathlib.Base(filename)

  return pageTpl.Execute(w, map[string]interface{}{
    "Title": title,
    "Parts": crumbs(filename, title),
    "Styles": stylePaths("chroma.css", "style.css"),
    "Content": template.HTML(buf.Bytes()),
  })
}

func stylePaths(names ...string) []string {
  var out []string
  for _, name := range names {
    out = append(out, pathlib.Clean("/" + pathlib.Join(basePath, "_templates", name)))
  }
  return out
}

func isIndex(path string) (string, bool) {
  potential := []string{"index.html", "index.md"}

  for _, poten := range potential {
    indexPath := filepath.Join(path, poten)
    ok, _ := exists(indexPath)
    if ok {
      return indexPath, true
    }
  }
  return "", false
}

func markdownPage(w io.Writer, path string) error {

  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()

  fm, body, err := frontmatter(f)
  if err != nil {
    return err
  }

  b := blackfriday.Run(body, blackfriday.WithRenderer(bfchroma.NewRenderer(
    bfchroma.ChromaStyle(Style),
  )))

  pageTplBytes := MustAsset("page.html")
  pageTpl := template.Must(template.New("page").Parse(string(pageTplBytes)))

  var title string
  if t, ok := fm["title"].(string); ok {
    title = t
  } else {
    title = pathlib.Base(path)
    title = strings.Replace(title, ".md", "", -1)
    title = strings.Replace(title, ".html", "", -1)
  }

  return pageTpl.Execute(w, map[string]interface{}{
    "Title": title,
    "Parts": crumbs(path, title),
    "Styles": stylePaths("chroma.css", "style.css"),
    "Content": template.HTML(b),
  })
}

func renderTable(w http.ResponseWriter, path string) error {

  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()

  r := csv.NewReader(f)
  r.Comma = '\t'
  r.ReuseRecord = true

  styleBytes := MustAsset("style.css")
  fmt.Fprintf(w, "<style>")
  w.Write(styleBytes)
  fmt.Fprintf(w, "</style>")

  fmt.Fprintf(w, "<table>")
  for {
    rec, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      return err
    }
    fmt.Fprintf(w, "<tr>")
    for _, col := range rec {
      fmt.Fprintf(w, "<td>%s</td>", col)
    }
    fmt.Fprintf(w, "</tr>")
  }
  fmt.Fprintf(w, "</table>")
  return nil
}

func sendfile(w http.ResponseWriter, path string) {

  var err error
  ext := filepath.Ext(path)

  switch ext {
  case ".css":
    w.Header().Set("Content-Type", "text/css")
    err = sendRaw(w, path)
  case ".html":
    w.Header().Set("Content-Type", "text/html")
    err = sendRaw(w, path)
  case ".md":
    w.Header().Set("Content-Type", "text/html")
    err = markdownPage(w, path)
  case ".ipynb":
    w.Header().Set("Content-Type", "text/html")
    err = notebookPage(w, path)
  case ".tsv":
    w.Header().Set("Content-Type", "text/html")
    err = renderTable(w, path)
  default:
    err = highlight(w, path)
    if err == nil {
      w.Header().Set("Content-Type", "text/html")
    }
    if err == errUnidentifiedLanguage {
      err = sendRaw(w, path)
    }
  }

  if err != nil {
    fmt.Fprintln(w, err)
  }
}

func sendRaw(w io.Writer, filename string) error {
  f, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer f.Close()
  _, err = io.Copy(w, f)
  return err
}

func newurl(p string) string {
  return pathlib.Clean(pathlib.Join("/", basePath, p))
}

type dirent struct {
  Name string
  HREF string
  Type string
}

func listdir(w http.ResponseWriter, req *http.Request, path string) {
  w.Header().Set("Content-Type", "text/html")
  dirListTplBytes := MustAsset("dirlist.html")
  dirListTpl := template.Must(template.New("dirlist").Parse(string(dirListTplBytes)))

  files, err := ioutil.ReadDir(path)
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  files = filterPrefix(files, ".")
  files = filterPrefix(files, "venv")
  files = filterPrefix(files, "_")
  files = filterSuffix(files, ".pyc")

  var list []dirent
  var readme string

  for _, file := range files {
    name := file.Name()

    var type_ string
    var p string

    if strings.ToLower(name) == "readme.md" {
      inb, err := ioutil.ReadFile(pathlib.Join(path, name))
      if err == nil {
        b := blackfriday.Run(inb)
        readme = string(b)
      }
    }

    if strings.HasSuffix(name, ".href") {
      b, _ := ioutil.ReadFile(filepath.Join(path, name))
      p = strings.TrimSpace(string(b))
      name = strings.TrimSuffix(name, ".href")
      type_ = "external-link-alt"
    } else {
      p = newurl(pathlib.Clean("/" + pathlib.Join(path, name) + "/"))
    }

    if info, err := os.Stat(filepath.Join(path, name)); err == nil && info.IsDir() {
      type_ = "folder"
    }

    list = append(list, dirent{Name: name, HREF: p, Type: type_})
  }

  title := pathlib.Base(path)
  if title == "." || title == "" {
    title = "home"
  }

  err = dirListTpl.Execute(w, map[string]interface{}{
    "Title": title,
    "Parts": crumbs(path, title),
    "List": list,
    "Styles": stylePaths("chroma.css", "style.css"),
    "Readme": template.HTML(readme),
  })
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }
}

type crumb struct {
  Name string
  Path string
}

func crumbs(path, title string) []crumb {
  var p string
  out := []crumb{
    {Name: "home", Path: pathlib.Clean(pathlib.Join("/", basePath))},
  }
  parts := strings.Split(path, "/")
  for _, part := range parts[:len(parts)-1] {
    if part == "." {
      continue
    }
    if part != "" {
      p = pathlib.Join(p, part)
      out = append(out, crumb{Name: part, Path: pathlib.Clean(pathlib.Join("/", basePath, p, "/")) + "/"})
    }
  }
  if len(out) > 1 {
    out[len(out)-1].Path = ""
  }
  return out
}

// exists returns whether the given file or directory exists or not
func exists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func filterPrefix(files []os.FileInfo, prefix string) []os.FileInfo {
  var out []os.FileInfo
  for _, f := range files {
    name := f.Name()
    if strings.HasPrefix(name, prefix) {
      continue
    }
    out = append(out, f)
  }
  return out
}
func filterSuffix(files []os.FileInfo, suffix string) []os.FileInfo {
  var out []os.FileInfo
  for _, f := range files {
    name := f.Name()
    if strings.HasSuffix(name, suffix) {
      continue
    }
    out = append(out, f)
  }
  return out
}

func main() {
  log.Println("Hello, Maude.")
  flag.Parse()

  handle := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    path := strings.TrimPrefix(req.URL.Path, "/")

    if path == "" {
      path = "."
    }
    log.Printf("PATH %q", path)

    info, err := os.Stat(path)
    if err == nil {

      if info.IsDir() {

        // TODO pull more canonical path redirects from net/http/fs.go
        // TODO correctly handle file mime types, e.g. for .css files
        //      fall back to net/http/FileServer?

        // TODO things are very broken here
        url := req.URL.Path
		    if len(url) != 0 && url[len(url)-1] != '/' {
          target := newurl("/"+url+"/") + "/"
          log.Println("redirect", target)
          localRedirect(w, req, target)
          return
        }

        if indexPath, ok := isIndex(path); ok {
          sendfile(w, indexPath)
          return
        }

        listdir(w, req, path)
        return

      } else {
        sendfile(w, path)
        return
      }

    } else {
      fmt.Fprintln(w, err)
    }
  })

  mux := http.NewServeMux()

  var fs = http.FileServer(&assetfs.AssetFS{
    Asset:     Asset,
    AssetDir:  AssetDir,
    AssetInfo: AssetInfo,
  })
  mux.HandleFunc("/_templates/chroma.css", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/css")
    formatter := html.New()
    formatter.WriteCSS(w, Style)
  })
  mux.Handle("/_templates/", http.StripPrefix("/_templates/", fs))

  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    log.Println(req.URL)
    handle(w, req)
  })

  log.Println("Listening on :7279")
  http.Handle("/", http.StripPrefix(basePath, mux))
  http.ListenAndServe(":7279", nil)
}

// localRedirect gives a Moved Permanently response.
// It does not convert relative paths to absolute paths like Redirect does.
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}
