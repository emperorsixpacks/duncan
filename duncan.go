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

var duncan *Duncan

const DEFAULT_PORT = 5000
const DEFAULT_HOST = "127.0.0.1"

type Context map[string]any
type RedisConnetion struct {
	Addr     string
	Password string
	DB       uint
}

type Duncan struct {
	Name        string
	Host        string
	Port        uint
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
	return fmt.Sprintf("%v:%v", this.Host, this.Port)
}

func (this *Duncan) AddRouter(router *routers.Router) {
	this.router = router
}

func (this *Duncan) initHTTPserver() {
	this.server = &http.Server{
		//		Handler:      this.router.GetHandler(),
		Addr:         this.getServerAddress(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func FromConfig(configPath string) (*Duncan, error) {
	if duncan == nil {
		var duncanConfig DuncanConfig
		err := validPath(configPath)
		config, err := loadConfig(configPath)
		err = yaml.Unmarshal(config, &duncanConfig)
		if err != nil {
			return nil, err
		}
		newDuncan(duncanConfig)
	}
	return duncan, nil
}

func Defualt() *Duncan {
	return &Duncan{
		Name: "MeetUps Guru",
		Host: DEFAULT_HOST,
		Port: DEFAULT_PORT,
	}
}

func newDuncan(config DuncanConfig) error {
	duncan = &Duncan{
		Name: config.App.Name,
		Host: config.App.Host,
		Port: config.App.Port,
	}
	return nil
}

// TODO do not know if it will work, but how about using factory here
