package cards

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

//get etcd connection

//create and persist to etcd
func CreateAndPersist(config Config, key clientv3.KV, ctx context.Context) {

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

	resp, _ := key.Put(ctx, cardID, greeting)

	fmt.Printf("updated records: %v", resp.Header.Revision)

}

//fetch single record

func FetchCard(id string, key clientv3.KV, ctx context.Context) {

	resp, _ := key.Get(ctx, id)

	fetch := string(resp.Kvs[0].Value)

	fmt.Println("Fetch Cards succefully: ", fetch)

}

///fetch all records
