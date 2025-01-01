package duncan

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

var duncan *Duncan

// TODO get dir name and use as default name
// TODO we can use a default sqlite databse, so we can ceate one if none is provided in config or if as default
const DEFAULT_PORT = "5000" // NOTE this should be a unit, unable to unmartal from env
const DEFAULT_HOST = "127.0.0.1"

var DEFAULT_CONFIG = DuncanConfig{
	App: Appconfig{
		Name: "dunan",
		Port: DEFAULT_PORT,
		Host: DEFAULT_HOST,
	},
}

type Context map[string]any

// TODO move out of here
type RedisConnetion struct {
	Addr     string
	Password string
	DB       string
}

func (this *RedisConnetion) GetDBVal() int {
	uintval, err := strconv.ParseUint(this.DB, 10, 8)
	if err != nil {
		panic(fmt.Errorf("Redisconnnection.DB: Could not parse string value %v to unit", this.DB))
	}
	return int(uintval)
}

type Duncan struct {
	Name        string
	Host        string
	Port        string
	server      *http.Server
	router      *Router
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

func New() *Duncan {
	if duncan == nil {
		err := newDuncan(DEFAULT_CONFIG)
		if err != nil {
			panic(errors.New("Could not create server"))
		}
	}
	return duncan
}

func newDuncan(config DuncanConfig) error {
	duncan = &Duncan{
		Name: config.App.Name,
		Host: config.App.Host,
		Port: config.App.Port,
	}
	duncan.initHTTPserver()
	return nil
}

// You know, factory can be used here
// TODO uint unmartial fail, from str port
// Routers manages handling, routes are the routes registered in that router. so router.addRoute() will add a new route to that router.
// This means that in router, all we need to do is just manages routing between routes. This means that the router is a tree and the routes are nodes. The router could also be an edge that is wen we create a base router or group
