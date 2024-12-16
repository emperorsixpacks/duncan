package duncan

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/emperorsixpacks/duncan/routers"
	"gopkg.in/yaml.v3"
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
	name        string
	host        string
	port        int
	server      *http.Server
	router      *routers.Router
	template    *template.Template
	middlewares []MiddleWare
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

func (this *Duncan) initHTTPserver() {
	this.server = &http.Server{
		Handler:      this.router.GetHandler(),
		Addr:         this.getServerAddress(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func NewFromConfig(configPath string) error {
	var duncanConfig DuncanConfig
	err := validPath(configPath)
	config, err := loadConfig(configPath)
	err = yaml.Unmarshal(config, &duncanConfig)
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
