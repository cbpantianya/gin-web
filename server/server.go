package server

import (
	"gin-template/server/config"
	"gin-template/server/cron"
	"gin-template/server/database"
	"gin-template/server/router"
	"gin-template/server/webEngine"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Instance *Server

type Server struct {
	Engine *gin.Engine
	DB     *database.Database
}

func Init() {
	logrus.WithField("server", "Global").Info("Init Global server...")

	// Init config
	config.InitConfig()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// init WebEngine and Database
	Instance = &Server{
		Engine: webEngine.InitWebEngine(),
		DB:     database.InitDB(),
	}
	// init router
	err := router.InitRouter(Instance.Engine)
	if err != nil {
		return
	}
	// init cron
	cron.InitCron()
	// init handler

}

func Start() {
	// start server
	logrus.WithField("server", "Global").Info("All Server start success!")
	err := Instance.Engine.Run(config.Cfg.Server.Host + ":" + config.Cfg.Server.Port)
	if err != nil {
		panic(err)
		return
	}

}
