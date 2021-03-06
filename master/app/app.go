package app

import (
	"flag"
	"fmt"
	"mcrontab/master/api"
	"mcrontab/master/config"
	"mcrontab/master/job"
	"runtime"
)

type App struct {
}

var (
	configFile string
)

func (this *App) Run() {
	var (
		err error
	)

	if err = this.initApp(); err != nil {
		goto Err
	}
Err:
	fmt.Println(err)
}

func (this *App) initApp() (err error) {
	initArgs()
	initEnv()
	if err = initConfig(); err != nil {
		return
	}
	if err = initJobManager(); err != nil {
		return
	}
	if err = initApiServer(); err != nil {
		return
	}

	return
}

//初始化命令行参数
func initArgs() {
	flag.StringVar(&configFile, "config", "src/mcrontab/master/config/config.json", "master配置文件")
	flag.Parse()
}

//初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

//加载配置文件
func initConfig() (err error) {
	if err = config.LoadConfig(configFile); err != nil {
		return
	}
	return
}

func initApiServer() (err error) {
	err = api.InitApiServer()
	return
}

func initJobManager() (err error) {
	err = job.InitJobManager()
	return
}

func NewApp() *App {
	return &App{}
}
