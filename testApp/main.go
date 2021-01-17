package main

import (
	"bitbucket.org/dream_yun/moduleapi/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Application":	"TestApp",
		"Version":			"v0.1.1",
	}).Info("Application start...")
	myapp := app.New("jjh")
	myapp.Run()


}
