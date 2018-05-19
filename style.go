package main

import (
  "github.com/alecthomas/chroma"
)

var kwd = "#5d86ab"
var plain = "#dedede"

var Style = chroma.MustNewStyle("maude", chroma.StyleEntries{
	chroma.Text:                "#F8F8F2",
	chroma.Error:               "#bbb57f bg:#1e0010",
	chroma.Comment:             "#636363",
	chroma.Keyword:             kwd,
	chroma.KeywordNamespace:    kwd,
	chroma.Operator:            "#a2a2a2",
	chroma.Punctuation:         plain,

	chroma.Name:                plain,
	chroma.NameAttribute:       plain,
	chroma.NameClass:           plain,
	chroma.NameConstant:        "#66d9ef",
	chroma.NameDecorator:       plain,
	chroma.NameException:       plain,
	chroma.NameFunction:        plain,
	chroma.NameOther:           plain,
	chroma.NameTag:             kwd,

	chroma.LiteralNumber:       "#ae81ff",
	chroma.Literal:             "#ae81ff",
	chroma.LiteralDate:         "#bbb57f",
	chroma.LiteralString:       "#bbb57f",
	chroma.LiteralStringEscape: "#ae81ff",

	chroma.GenericDeleted:      kwd,
	chroma.GenericEmph:         "italic",
	chroma.GenericInserted:     plain,
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   "#75715e",
	chroma.Background:          "bg:#171717",
})
