package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ChunHou23/catalogue-service/internal/base/setup"
	"github.com/ChunHou23/catalogue-service/internal/config"
)

var appConfig config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	appConfig.InfoLog = infoLog
	appConfig.ErrorLog = errorLog
	appConfig.InProduction = false

	srv := &http.Server{
		Addr:    ":8080",
		Handler: setup.SetupControllers(&appConfig),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
