package main

import (
	"fmt"
	"testing"
	"strconv"
	log "github.com/Sirupsen/logrus"
)

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
			fmt.Println("读取",temp)
		}(chan1)
	}
}

func Test_log(test *testing.T)  {

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}


