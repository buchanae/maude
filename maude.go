package main

import (
  "fmt"
  pathlib "path"
  "html/template"
  "path/filepath"
  "io"
  "io/ioutil"
  "log"
  "os"
  "strings"
  "net/http"
  urllib "net/url"

  "gopkg.in/russross/blackfriday.v2"
  "github.com/alecthomas/chroma/quick"
)

const basePath = "/buchanan"

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

func sendfile(w http.ResponseWriter, path string) {

  ext := filepath.Ext(path)
  switch ext {
  case ".html":
    w.Header().Set("Content-Type", "text/html")
  case ".css":
    w.Header().Set("Content-Type", "text/css")
  case ".md":
    w.Header().Set("Content-Type", "text/html")
  case ".go":
    w.Header().Set("Content-Type", "text/html")
  default:
    w.Header().Set("Content-Type", "text/plain")
  }

  f, err := os.Open(path)
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }
  defer f.Close()

  switch ext {
  case ".go":
    inb, err := ioutil.ReadAll(f)
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }
    err = quick.Highlight(w, string(inb), "go", "html", "monokai")
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }

  case ".md":
    inb, err := ioutil.ReadAll(f)
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }
    b := blackfriday.Run(inb)
    _, err = w.Write(b)
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }

  default:
    _, err = io.Copy(w, f)
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }
  }
}

func newurl(p string) string {
  return pathlib.Clean(pathlib.Join("/", basePath, p))
}

type dirent struct {
  Name string
  HREF string
}

func listdir(w http.ResponseWriter, req *http.Request, path string) {
  w.Header().Set("Content-Type", "text/html")
  dirListTpl := template.Must(template.ParseGlob("_templates/*.html"))

  files, err := ioutil.ReadDir(path)
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  files = filterPrefix(files, ".")
  files = filterPrefix(files, "_")

  var list []dirent

  if req.URL.Path != "/" {
    up, _ := urllib.Parse("./..")
    u := req.URL.ResolveReference(up)
    list = append(list, dirent{Name: "..", HREF: newurl(u.String())})
  }

  for _, file := range files {
    name := file.Name()
    var p string
    if strings.HasSuffix(name, ".href") {
      b, _ := ioutil.ReadFile(filepath.Join(path, name))
      p = strings.TrimSpace(string(b))
      name = strings.TrimSuffix(name, ".href")
    } else {
      p = newurl(pathlib.Clean("/" + pathlib.Join(path, name) + "/"))
    }
    list = append(list, dirent{Name: name, HREF: p})
  }

  err = dirListTpl.Execute(w, map[string]interface{}{
    "List": list,
    "Styles": pathlib.Join(basePath, "_templates", "style.css"),
  })
  if err != nil {
    fmt.Fprintln(w, err)
    return
  }
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

func main() {
  log.Println("Hello, Maude.")

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

        /*
        url := req.URL.Path
		    if len(url) == 0 || url[len(url)-1] != '/' {
          localRedirect(w, req, newurl("/"+pathlib.Base(url)+"/"))
          return
        }
        */

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

  http.Handle("/", http.StripPrefix(basePath, handle))

  log.Println("Listening on :7279")
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
