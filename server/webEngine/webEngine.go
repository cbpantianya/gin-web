package webEngine

import (
	"gin-template/server/slogger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitWebEngine() *gin.Engine {
	//Release Mode
	logrus.WithField("server", "Gin").Info("Start webEngine...")
	gin.SetMode(gin.ReleaseMode)
	webEngine := gin.New()
	webEngine.Use(slogger.SLogger())
	return webEngine
}
