package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/router"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	if err:=config {
		
	}

	r := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(r, middlewares...)

	//ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	r.Run()
}

func pingServer() error {
	for i := 0; i < 3; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/monitor/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
