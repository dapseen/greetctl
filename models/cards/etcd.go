package cards

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

//get etcd connection

//create and persist to etcd
func CreateAndPersist(config Config) {
	var (
		timeout        = 2 * time.Second
		requestTimeout = 10 * time.Second
	)

	cardID := config.CardID
	occasion := config.Occasion
	lang := config.Language
	name := config.UserName

	var greeting string

	if occasion == newYear && lang == langEN {
		greeting = newYearGreetingENG + name
	} else if occasion == newYear && lang == langFR {
		greeting = newYearGreetingFR + name
	} else if occasion == txGiving && lang == langEN {
		greeting = txDayGreetingENG + name
	} else if occasion == txGiving && lang == langFR {
		greeting = txDayGreetingFR + name
	} else if occasion == bDay && lang == langEN {
		greeting = bDayGreetingENG + name
	} else {
		greeting = bDayGreetingFR + name
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: timeout,
	})
	if err != nil {
		println(err)
	}

	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, _ := cli.Put(ctx, cardID, greeting)

	if cancel != nil {
		println(cancel)
	}

	log.Fatal(resp)

}

//fetch single record

///fetch all records
