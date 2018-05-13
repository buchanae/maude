package main

import (
  "encoding/json"
  "log"
  "bytes"
  pathlib "path"
  "io"
  "html/template"
  "strings"

  "github.com/alecthomas/chroma/lexers"
  chromaHTML "github.com/alecthomas/chroma/formatters/html"
  "github.com/alecthomas/chroma/styles"
  "io/ioutil"
)

func highlightNotebookSource(src string) interface{} {
  lexer := lexers.Get("python")
  /* TODO don't understand lexer analysis
  lexer := lexers.Analyse(src)
  if lexer == nil {
    log.Println("no lexer", src)
    return src
  }
  */

  style := styles.Get("monokai")
  formatter := chromaHTML.New(chromaHTML.WithClasses())

  iterator, err := lexer.Tokenise(nil, src)
  if err != nil {
    log.Println("token error", err)
    return src
  }

  w := &bytes.Buffer{}
  //formatter.WriteCSS(w, style)

  err = formatter.Format(w, style, iterator)
  if err != nil {
    log.Println("format err", err)
    return src
  }
  return template.HTML(w.String())
}

type Output struct {
  OutputType string `json:"output_type"`
  Name string `json:"name"`
  Data map[string]interface{} `json:"data"`

  //Data map[string][]string `json:"data"`
  Text []string `json:"text"`
}

func (o *Output) Formatted() interface{} {
  switch o.OutputType {
  case "stream":
    return strings.Join(o.Text, "")
  case "execute_result", "display_data":
    if d, ok := o.Data["text/html"]; ok {
      var str string
      for _, row := range d.([]interface{}) {
        str += row.(string)
      }
      return template.HTML(str)
    }
    if d, ok := o.Data["text/plain"]; ok {
      var str string
      for _, row := range d.([]interface{}) {
        str += row.(string)
      }
      return str
    }
  }
  return ""
}

type Cell struct {
  CellType string `json:"cell_type"`
  Metadata struct{} `json:"metadata"`
  Source []string `json:"source"`
  Outputs []Output `json:"outputs"`
  ExecutionCount int `json:"execution_count"`
}
func (cell *Cell) Highlighted() interface{} {
  src := strings.Join(cell.Source, "")
  if cell.CellType != "code" {
    return src
  }
  return highlightNotebookSource(src)
}

type Notebook struct {
  Cells []Cell `json:"cells"`
}

func notebookPage(w io.Writer, path string) error {
  b, err := ioutil.ReadFile(path)
  if err != nil {
    return err
  }
  nb := &Notebook{}
  err = json.Unmarshal(b, nb)
  if err != nil {
    return err
  }


  tplBytes := MustAsset("notebook.html")
  tpl := template.Must(template.New("notebook").Parse(string(tplBytes)))

  err = tpl.Execute(w, map[string]interface{}{
    "NB": nb,
    "Styles": pathlib.Clean("/" + pathlib.Join(basePath, "_templates", "notebook.css")),
  })
  if err != nil {
    return err
  }
  return nil
}
