package main

import (
	"errors"

	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/config"
	"github.com/hyper-xx/mygin/model"
	"github.com/hyper-xx/mygin/router"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	//Init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//Init DB
	model.DB.Init()

	//Set gin runmode
	gin.SetMode(viper.GetString("runmode"))

	//Create gin engine
	r := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(r, middlewares...)

	//ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), r).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/monitor/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
