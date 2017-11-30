package main

import (
	"github.com/gin-gonic/gin"
	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"

	"github.com/Depado/opensirene-ui/conf"
	"github.com/Depado/opensirene-ui/views"
)

func main() {
	var err error
	var config string
	var api string

	// Flag parsing
	flag.StringVarP(&config, "config", "c", "conf.yml", "Path to the configuration file")
	flag.StringVarP(&api, "opensirene-api", "a", "", "URL of a deployed OpenSirene API")
	flag.Parse()

	// Configuration
	if err = conf.Load(config); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	if api != "" {
		conf.C.OpenSireneAPI = api
	}

	// Setup debug mode or not in Gin
	if !conf.C.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create the router
	r := gin.Default()

	// Static and template loading
	r.Static("/static/", "./assets")
	r.LoadHTMLGlob("./templates/*.tmpl")

	// Route to access the UI
	r.GET("/ui", views.GetUI)
	r.GET("/ui/:data", views.GetUIData)

	// Run the server
	logrus.WithFields(logrus.Fields{"port": conf.C.Server.Port, "host": conf.C.Server.Host}).Info("Starting server")
	if err = r.Run(conf.C.Server.ListenString()); err != nil {
		logrus.WithError(err).Fatal(err)
	}
}
