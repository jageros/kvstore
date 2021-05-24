package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strsvc/apps/strsrv/service"
	"strsvc/internal/attribute"
	"strsvc/internal/conf"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	conf.Initialize("config.toml")
	attribute.InitializeDB(conf.CFG.DBAddr, conf.CFG.DBName, conf.CFG.DBUser, conf.CFG.DBPassword)
	defer attribute.Stop()

	gin.SetMode(conf.CFG.GinModel)
	router := gin.Default()
	router.Use(cors.Default())
	r := router.Group("/api")

	// ======== RegisterHandle =========
	service.RegisterHandle(r)
	// =================================
	//r.Run() // listen and serve on 0.0.0.0:8080

	err := endless.ListenAndServe(conf.CFG.ListenAddr, router)
	if err != nil {
		log.Printf("ListenAndServe err=%v", err)
	}
}
