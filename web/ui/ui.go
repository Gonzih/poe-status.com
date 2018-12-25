//go:generate go run -tags=dev template_generate.go

package ui

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"gitlab.com/Gonzih/poe-status.com/db"
)

type indexTemplateData struct {
	Results []*db.ScanResult
}

// Index renders index page of the web ui
func Index(res http.ResponseWriter) error {
	tmpl, err := parseTemplate("index.html")
	if err != nil {
		return err
	}
	results, err := db.AllScanResults(db.DB())
	if err != nil {
		return err
	}

	data := &indexTemplateData{
		Results: results,
	}

	tmpl.Execute(res, data)
	// res.Write([]byte("hello world!"))
	return nil
}

func parseTemplate(fname string) (*template.Template, error) {
	file, err := TemplateDir.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fdata, err := ioutil.ReadAll(file)
	tmpl := template.New(fname)
	tmpl, err = tmpl.Parse(string(fdata))
	if err != nil {
		return nil, err
	}

	return tmpl, err
}
