package main

import (
	"os"

	"./optionsHandler"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func main() {
	opts := optionsHandler.Options{
		AccountSid: os.Getenv("SID"),
		AuthToken:  os.Getenv("TOKEN"),
		Receiver:   os.Getenv("RECEIVER"),
		Sender:     os.Getenv("SENDER"),
	}

	if opts.AccountSid == "" || opts.AuthToken == "" || opts.Sender == "" {
		log.Fatal("'SID', 'TOKEN' and 'SENDER' environment variables need to be set")
	}

	o := optionsHandler.NewMOptionsWithHandler(&opts)
	err := fasthttp.ListenAndServe(":9090", o.HandleFastHTTP)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
