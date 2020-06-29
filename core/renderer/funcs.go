package renderer

import (
	"bytes"
	"strconv"
	"text/template"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"ucFirst":       UcFirst,
		"lcFirst":       LcFirst,
		"quote":         strconv.Quote,
		"rawQuote":      rawQuote,
		"dump":          Dump,
		"ref":           ref,
		"ts":            TypeIdentifier,
		"call":          Call,
		"prefixLines":   prefixLines,
		"notNil":        notNil,
		"reserveImport": CurrentImports.Reserve,
		"lookupImport":  CurrentImports.Lookup,
		"go":            ToGo,
		"goPrivate":     ToGoPrivate,
		"add": func(a, b int) int {
			return a + b
		},
		"render": func(filename string, tpldata interface{}) (*bytes.Buffer, error) {
			return render(resolveName(filename, 0), tpldata)
		},
	}
}
