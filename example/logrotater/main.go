package main

import (
	"github.com/xavierzho/pkg/logger"
	"go.uber.org/zap"
	"sync"
)

func main() {
	log, err := logger.New(
		logger.WithFileRotation("log/test.log"),
		//logger.WithDisableConsole(),
	)
	if err != nil {
		panic(err)
	}
	//var ch = make(chan struct{}, 1000)
	defer log.Sync()
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			log.Error("err occurs", zap.String("para1", "value1"), zap.String("para2", "value2"))
			wg.Done()
		}()
	}
	wg.Wait()

}
