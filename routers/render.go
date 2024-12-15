package routers

import (
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Define how we want our html render to look like
// Be able to load temples from a directory
// Be able to load a single template
// Be able to pass arguments, using structs
// Be able to call a Render method to send the new html to the page

// Based off the HTML struct from gin
type HTML struct {
	Template *template.Template
}

// this will be called from requests, like do r.Render
func (this HTML) Render(w http.ResponseWriter, t string, data interface{}) error {
	return this.Template.ExecuteTemplate(w, t, data)
}
func (this HTML) readHtmlFileString(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func (this *HTML) findTemplates(cleanRoot string) (int, []string, error) {
	last_index := len(cleanRoot) + 1
	html_files := []string{}
	err := filepath.Walk(cleanRoot, func(path string, info fs.FileInfo, file_err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if file_err != nil {
				return file_err
			}
			html_files = append(html_files, path)
		}
		return nil
	})
	return last_index, html_files, err
}
func (this *HTML) parseTemplatetoRoot(rootTemplate *template.Template, name string, html_path string) error {
	new_template := rootTemplate.New(name)
	html_file, err := os.ReadFile(html_path)
	if err != nil {
		return err
	}
	_, err = new_template.Parse(string(html_file))
	if err != nil {
		return err
	}
	return nil
}

func (this *HTML) loadTemplates(template_path string) error {
	rootTemaplate := template.New("")
	//	cleanRoot := filepath.Clean(template_path)
	last_index, html_files, err := this.findTemplates(template_path)
	if err != nil {
		return err
	}
	for _, html_file := range html_files {

		name := html_file[last_index:]
		err := this.parseTemplatetoRoot(rootTemaplate, name, html_file)
		if err != nil {
			return err
		}
	}
	this.Template = rootTemaplate
	return nil
}

// Look into moving all these template stuff to another place

func (this *HTML) LoadTemplates(template_path string) error {
	return this.loadTemplates(template_path)
}

// Now we need to create functions and methods to load the passd html
