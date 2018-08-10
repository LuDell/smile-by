package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_run(t *testing.T)  {
	var i = 0
	for  i = 0 ;i<10;i++ {
		go func (i int)  {
			i = i + 100
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1*time.Second)
}


