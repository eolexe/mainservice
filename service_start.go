package mainservice

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func ServiceStart(service Service, configMethod ConfigMethod, waitGroup *sync.WaitGroup) (chanStop chan int) {
	var err error
	
	if waitGroup != nil {
		waitGroup.Add(1)
	}

	config := service.NewConfig()
	flConfigPath := flag.String("c", "", "config path")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	if *flConfigPath == "" {
		log.Fatalf("need config file")
	}

	if err = configMethod(config, *flConfigPath); err != nil {
		log.Fatalf("load config err: %s", err)
	}

	if err = service.Init(); err != nil {
		log.Fatalf("init err: %s", err)
	}
	chSignals := make(chan os.Signal)
	chStop := make(chan int)
	go func() {
		if err = service.Run(); err != nil {
			log.Fatalf("run err: %s", err)
		}
		chStop <- 1
	}()
	signal.Notify(chSignals, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		select {
		case sig := <-chSignals:
			log.Printf("received signal %d", sig)
		case <-chStop:
			log.Printf("received command stop")
		}
		log.Printf("stopping")
		if err = service.Stop(); err != nil {
			log.Fatalf("stop err: %s", err)
		}
		if waitGroup != nil {
			waitGroup.Done()
		}
	}()
	return chStop
}
