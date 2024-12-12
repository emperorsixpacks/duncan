package duncan

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/emperorsixpacks/duncan/routers"
)

const DEFAULT_PORT = 5000
const DEFAULT_HOST = "127.0.0.1"

type Context map[string]any
type RedisConnetion struct {
	Addr     string
	Password string
	DB       int
}

type Duncan struct {
	name     string
	host     string
	port     int
	server   *http.Server
	router   *routers.Router
	template *template.Template
}

func (this *Duncan) Start() {
	log.Print("Starting Duncan Server")
	log.Print("Server has started on : ", this.getServerAddress())
	this.initHTTPserver()
	err := this.server.ListenAndServe()
	if err != nil {
		log.Fatal("could not start server : ", err)
		return
	}
}

func (this *Duncan) Stop() {
	return
}

func (this *Duncan) getServerAddress() string {
	return fmt.Sprintf("%v:%v", this.host, this.port)
}

func (this *Duncan) AddRouter(router *routers.Router) {
	this.router = router
}

func (this Duncan) readHtmlFileString(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func (this *Duncan) findTemplates(cleanRoot string) (int, []string, error) {
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
func (this *Duncan) parseTemplatetoRoot(rootTemplate *template.Template, name string, html_path string) error {
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

func (this *Duncan) loadTemplates(template_path string) error {
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
	this.template = rootTemaplate
	return nil
}

// Look into moving all these template stuff to another place

func (this *Duncan) LoadTemplates(template_path string) error {
	return this.loadTemplates(template_path)
}

func (this Duncan) RenderHtml(w http.ResponseWriter, name string, data interface{}) error {
	err := this.template.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}

func (this *Duncan) initHTTPserver() {
	this.server = &http.Server{
		Handler:      this.router.R,
		Addr:         this.getServerAddress(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func NewFromConfig(configPath string) error {
	if err := validPath(configPath); err != nil {
		return nil
	}
	file, err := loadConfig(configPath)
	if err != nil {
		return err
	}
	o, err := ymltoMap(file)
	fmt.Println(o)
	if err != nil {
		return err
	}
	return nil
}

func Defualt() *Duncan {
	return &Duncan{
		name: "MeetUps Guru",
		host: DEFAULT_HOST,
		port: DEFAULT_PORT,
	}
}

// TODO do not know if it will work, but how about using factory here
