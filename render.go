package duncan

import (
  "html/template"
	"net/http"
)

// Define how we want our html render to look like
// Be able to load temples from a directory
// Be able to load a single template
// Be able to pass arguments, using structs
// Be able to call a Render method to send the new html to the page

// Based off the HTML struct from gin
type HTML struct {
	Name     string
	Data     any
	Template *template.Template
}

func (this HTML) Render(w http.ResponseWriter) error {
	return this.Template.ExecuteTemplate(w, this.Name, this.Data)

}

// Now we need to create functions and methods to load the passd html
