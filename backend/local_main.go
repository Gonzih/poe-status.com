package main

import (
	"context"
	"encoding/json"
	"log"
)

func localMain() {
	data, err := HandleRequest(context.TODO())
	log.Println(err)
	b, _ := json.Marshal(data)
	log.Println(string(b))
}
