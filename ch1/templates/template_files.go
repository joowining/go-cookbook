package templates

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func CreateTemplate(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data),
		os.FileMode(0755))
}

func InitTemplates() error {
	tempdir, err := ioutil.TempDir("", "temp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(tempdir)
	err = CreateTemplate(filepath.Join(tempdir, "t1.tmpl"),
		`Template 1! {{ .Var1 }}
		{{ block "template2" . }} {{end}}
		{{ block "template3" . }} {{end}}
		`)
	if err != nil {
		return nil
	}
	err = CreateTemplate(filepath.Join(tempdir, "t2.tmpl"),
		`{{ define "template2" }} Template2! {{ .Var2 }} {{ end }}`)
	if err != nil {
		return err
	}
	err = CreateTemplate(filepath.Join(tempdir, "t3.tmpl"),
		`{{ define "template3" }} TEmplate 3! {{ .Var3 }} {{ end }}`)
	if err != nil {
		return err
	}

	pattern := filepath.Join(tempdir, "*.tmpl")
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	tmpl.Execute(os.Stdout, map[string]string{
		"Var1": "Var1!",
		"Var2": "Var2!!",
		"Var3": "Var3!!!",
	})
	return nil
}
