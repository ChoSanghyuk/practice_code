package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestCron(t *testing.T) {

	c := cron.New()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Printf("HELLO, %s\n", time.Now().Format("150405"))
	})
	c.AddFunc("10,20,30,40,50,0 * * * * *", func() {
		fmt.Printf("HELLO2, %s\n", time.Now().Format("150405"))
	})
	c.Start()
	fmt.Printf("START %s\n\n", time.Now().Format("150405"))

	time.Sleep(1 * time.Minute)
}
