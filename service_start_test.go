package mainservice

import (
	"testing"
	"sync"
	"time"
	"os"
)

type testApp struct {
	needStop bool
}

func (a *testApp) Init() error {
	return nil
}

func (a *testApp) NewConfig() interface{} {
	return &testConfigINI{}
}

func (a *testApp) Run() error {
	for !a.needStop {
		time.Sleep(time.Second)
	}
	return nil
}

func (a *testApp) Stop() error {
	a.needStop = true
	return nil
}

func TestRun(t *testing.T) {
	os.Args = append(os.Args, "-c", "test/config.ini")
	app := testApp{}
	waitGroup := sync.WaitGroup{}
	chStop := ServiceStart(&app, ConfigMethodINI, &waitGroup)
	chStop <- 1
	waitGroup.Wait()
}