package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/kafka-bridge/internal/model"
	"github.com/ujunglangit-id/kafka-bridge/internal/repository"
	"os"
)

func main() {
	//log setup
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	//load config
	cfg, err := model.InitConfig()
	if err != nil {
		log.Fatalf("err load config, %#v", err)
	}

	kfka := repository.New(cfg)
	err = kfka.InitBrokers()
	if err != nil {
		log.Fatalf("err init broker, %#v", err)
	}

	log.Infof("consumer started")
	os.Exit(0)
}
