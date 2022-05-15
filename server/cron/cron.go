package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

// Cron is the cron job runner.

type HandleCron interface {
	Register() map[string]func()
}

type config []struct {
	Name     string `yaml:"name"`
	MainFunc string `yaml:"main_func"`
	Schedule string `yaml:"schedule"`
}

var HandleCronMap = map[string]func(){}

func RegisterCronHandler(handler HandleCron) {
	for k, v := range handler.Register() {
		HandleCronMap[k] = v
	}
}

// InitCron initializes the cron job runner.
func InitCron() {
	var config config
	logrus.WithField("server", "cron").Info("Start cron...")
	c := cron.New(
		cron.WithSeconds(),
	)
	rawRouterYaml, err := os.ReadFile("./server/cron/cron.yaml")
	if err != nil {
		panic(err)
	} else {
		err := yaml.Unmarshal(rawRouterYaml, &config)
		if err != nil {
			panic(err)
		}
	}
	for _, v := range config {
		_, err := c.AddFunc(v.Schedule, HandleCronMap[v.MainFunc])
		if err != nil {
			panic(err)
		}
		logrus.WithField("server", "cron").Info("Register cron job:", v.Name)
	}
	_, err = c.AddFunc("@every 10s", func() {})
	if err != nil {
		return
	}
	c.Start()
}
