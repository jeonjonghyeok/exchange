package main

import (
	//"bitbucket.org/dream_yun/moduleapi/app"
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {
	fpLog, err := os.OpenFile("logfile.txt",os.O_CREATE |os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	log.SetOutput(fpLog)
	log.SetFormatter(&log.JSONFormatter{})

	/*coreFields:=log.Fields{"gopher_lagos": "staging-1","meetup": "foo-TestApp","session": "1ce3f6v",}
	log.WithFields(coreFields).WithFields(log.Fields{"product_type": "ticket","quantity": 3, "price":100.0}).Info("start ms")
	*/
	log.WithFields(log.Fields{
		"Application":	"TestApp",
		"Version":			"v0.1.1",
	}).Info("Application start...")

	log.WithFields(log.Fields{
		"ID":	"jjh",
	}).Error("Not fond user")
	log.Info("this iserr")

	//myapp := app.New("jjh")
	//myapp.Run()
}
