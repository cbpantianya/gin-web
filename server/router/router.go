package router

import (
	"gin-template/server/gateway"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type HandleRequest interface {
	Register() map[string]gin.HandlerFunc
}

var HandleRequestMap = map[string]gin.HandlerFunc{}

func RegisterRequestHandler(handler HandleRequest) {
	for k, v := range handler.Register() {
		HandleRequestMap[k] = v
	}
}

// TreeRouter is a model of tree of routers.
type TreeRouter struct {
	Router []struct {
		PreRouter   string         `yaml:"pre_router"`
		Title       string         `yaml:"title"`
		Version     string         `yaml:"version"`
		Paths       []ChildrenPath `yaml:"paths"`
		Description string         `yaml:"description"`
	} `yaml:"router"`
}

// ChildrenPath is a model of children path.
type ChildrenPath struct {
	Router   string                `yaml:"router"`
	Gateway  string                `yaml:"gateway"`
	Method   []map[string][]string `yaml:"method"`
	Children []ChildrenPath        `yaml:"children"`
}

func InitRouter(engine *gin.Engine) error {
	// Load router tree from yaml file.
	logrus.WithField("server", "Router").Info("Start add router...")
	tree := TreeRouter{}
	err := tree.analyzeRouter()
	if err != nil {
		return err
	} else {
		// Add router tree to gin engine.
		// Can you see the router tree? :)
		for _, router := range tree.Router {
			routerGroup := engine.Group(router.PreRouter)
			for _, path := range router.Paths {
				// From there to calculate children path
				childRouterRegister(routerGroup, path, "/")
			}
		}
	}

	return nil
}

// analyzeRouter is a function to analyze router and return struct.
func (t *TreeRouter) analyzeRouter() error {
	rawRouterYaml, err := os.ReadFile("./server/router/router.yaml")
	if err != nil {
		return err
	} else {
		err := yaml.Unmarshal(rawRouterYaml, &t)
		if err != nil {
			return err
		}
		return nil
	}
}

// childRouterRegister is a function to register the router tree.
func childRouterRegister(engine *gin.RouterGroup, router ChildrenPath, fatherPath string) {
	if fatherPath == "/" {
		fatherPath = ""
	}
	if router.Children != nil {
		// Register current router
		for _, requestMethod := range router.Method {
			for httpMethod, handlerMethod := range requestMethod {
				var handlerList []gin.HandlerFunc
				if router.Gateway == "base" {
					handlerList = append(handlerList, gateway.BaseCheckInGateway)
				} else if router.Gateway == "auth" {
					handlerList = append(handlerList, gateway.AuthCheckInGateway)
				} else if router.Gateway == "special" {
					handlerList = append(handlerList, gateway.SpecialCheckInGateway)
				} else if router.Gateway == "hard" {
					handlerList = append(handlerList, gateway.HardCheckInGateway)
				} else {
					logrus.WithField("server", "Router").Info("No gateway setting, use base gateway.")
					handlerList = append(handlerList, gateway.BaseCheckInGateway)
				}

				for _, handlerName := range handlerMethod {
					handlerList = append(handlerList, HandleRequestMap[handlerName])
				}
				logrus.WithField("server", "Router").Infof("Register %s...", fatherPath+router.Router)
				engine.Handle(httpMethod, fatherPath+router.Router, handlerList...)
			}
		}
		// Register children router
		for _, child := range router.Children {
			childRouterRegister(engine, child, router.Router)
		}
	} else {
		// Register current router
		for _, requestMethod := range router.Method {
			for httpMethod, handlerMethod := range requestMethod {
				var handlerList []gin.HandlerFunc
				if router.Gateway == "base" {
					handlerList = append(handlerList, gateway.BaseCheckInGateway)
				} else if router.Gateway == "auth" {
					handlerList = append(handlerList, gateway.AuthCheckInGateway)
				} else if router.Gateway == "special" {
					handlerList = append(handlerList, gateway.SpecialCheckInGateway)
				} else if router.Gateway == "hard" {
					handlerList = append(handlerList, gateway.HardCheckInGateway)
				} else {
					logrus.WithField("server", "Router").Info("No gateway setting, use base gateway.")
					handlerList = append(handlerList, gateway.BaseCheckInGateway)
				}

				for _, handlerName := range handlerMethod {
					handlerList = append(handlerList, HandleRequestMap[handlerName])
				}
				logrus.WithField("server", "Router").Infof("Register %s...", fatherPath+router.Router)
				engine.Handle(httpMethod, fatherPath+router.Router, handlerList...)
			}
		}
	}
}
