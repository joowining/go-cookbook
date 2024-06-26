package templates

import (
	"html/template"
	"os"
	"strings"
)

const sampleTemplate = `
	This template demonstrates printing a {{ .Variable | printf "%#v" }}.
	{{ if .Condition }}
	IF condition is set, we'll print this
	{{ else }}
	Otherwise, we'll print this instead
	{{ end }} 
	next we'll iterate over an array of strings :
	{{ range $index, $item := .Items }}
		{{ $index }} : {{ $item}}
	{{ end }}
	We can also easily import other functions like strings.Split
	then immediately used the array created as a result :
	{{ range $index, $item := split .Words "," }}
		{{ $index }}: {{ $item }}
	{{ end }}
	Blocks are a way to embed templates into one another
	{{ block "block_example" .}}
		No Block defined!
	{{ end }}
	{{/* 여러 줄 주석
		처리하는 방법 
		*/}}
`

const secondTemplate = `
	{{ define "block_example" }}
		{{ .OtherVariable }}
	{{ end }}
`

func RunTemplate() error {
	data := struct {
		Condition     bool
		Variable      string
		Items         []string
		Words         string
		OtherVariable string
	}{
		Condition:     true,
		Variable:      "variable",
		Items:         []string{"item1", "item2", "itme3"},
		Words:         "another_item1, another_item2, another_item3",
		OtherVariable: "I'm defined in a second template!",
	}
	funcmap := template.FuncMap{
		"split": strings.Split,
	}

	t := template.New("example")
	t = t.Funcs(funcmap)

	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}

	t2, err := t.Clone()
	if err != nil {
		return err
	}

	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}

	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}
	return nil
}
