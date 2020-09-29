//@Description 速率限制器
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/17 8:48 下午

package tools

import (
	"bilibili-spirder/utils"
	"math/rand"
	"time"
)

type Rate struct {
	c chan int
}

func NewRate(c int) *Rate {
	rate := &Rate{
		c: make(chan int, c),
	}
	rate.Start()
	return rate
}

func (r *Rate) Start() {
	go func(r *Rate) {
		for {
			rate := utils.Config.Rate
			t := rand.Intn(rate)
			r.c <- t
			time.Sleep(time.Duration(t) * time.Second)
		}
	}(r)
}

func (r *Rate) Get() chan int {
	return r.c
}
