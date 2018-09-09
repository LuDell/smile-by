package model

import (
	"os"
	"encoding/json"
	"fmt"
	"time"
)

var SeeLogConfig *Config

func init()  {
	SeeLogConfig = initConfig()
}

func initConfig() *Config {
	file,err1 := os.Open("config/config.json")
	defer file.Close()
	if err1 !=nil {
		fmt.Println("读取配置文件错误", err1)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err2:= decoder.Decode(&config)
	if err2 !=nil {
		fmt.Println("数据绑定错误",err2)
	}
	fmt.Println("config参数",config)
	return &config
}

type Config struct {
	Redis redis
	Mongo mongo
}

type redis struct {
	Tcp string
	Port string
	Password string
	Max_idle int
	Max_active int
}

type mongo struct {
	Tcp string
	Port string
	Database string
	User_name string
	Password string
	Timeout time.Duration
}
