package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"net/http"
	"time"

	v "github.com/hyper-xx/mygin/pkg/version"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/config"
	"github.com/hyper-xx/mygin/model"
	"github.com/hyper-xx/mygin/router"
	"github.com/hyper-xx/mygin/router/middleware"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show verison info.")
)

func main() {
	pflag.Parse()
	//shell verion -v
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}
	//Init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//Init DB
	model.DB.Init()
	defer model.DB.Close()

	//Set gin runmode
	gin.SetMode(viper.GetString("runmode"))

	//Create gin engine
	r := gin.New()

	//middlewares := []gin.HandlerFunc{}

	router.Load(
		r,
		//middlewares...
		middleware.RequestId(),
		middleware.Logging(),
	)

	//ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	//Start to listening
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, r).Error())
		}()
	}
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
