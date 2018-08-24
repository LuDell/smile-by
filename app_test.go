package main

import (
	"testing"
	"strconv"
	"smile-by/config"
	"time"
)

var logger = config.Logger

func Test_run(t *testing.T)  {
	var chan1 = make(chan string,3)
	for i :=0; i<10; i++ {
		go func(num int) {
			var temp= "chan" + strconv.FormatInt(int64(num),10)
			chan1 <- temp
		}(i)
	}

	for  i := 0 ; i<10; i++ {
		go func(chanTemp chan string) {
			var temp = <-chan1
			logger.Info("读取",temp)
		}(chan1)
	}

	time.Sleep(1 *time.Second)
}

func Test_log(test *testing.T)  {
	chanTemp := make(chan string,1)
	chanTemp <- "start"
	logger.Info(<- chanTemp)
	time.Sleep(1 *time.Second)
}


