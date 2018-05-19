package main

import (
  "io"
  "io/ioutil"
  "regexp"
  "gopkg.in/yaml.v2"
)

var fmrx = regexp.MustCompile("(?m)^---$")

func frontmatter(r io.Reader) (map[string]interface{}, []byte, error) {
  b, err := ioutil.ReadAll(r)
  if err != nil {
    return nil, nil, err
  }
  d := map[string]interface{}{}
  m := fmrx.FindIndex(b)
  if len(m) == 0 {
    return d, b, nil
  }
  start, end := m[0], m[1]
  raw := b[:start]
  body := b[end:]

  err = yaml.Unmarshal(raw, &d)
  return d, body, err
}
