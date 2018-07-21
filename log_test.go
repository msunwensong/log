package logservice

import (
	"testing"
	"time"
)

//测试整个流程
func TestLog(t *testing.T) {
	g := GoLog{}
	err := g.Init("d.xml")
	if err != nil {
		return
	}
	for i := 1; i <= 50; i++ {
		go g.Info(i)

	}
	time.Sleep(10 * time.Second)
}
