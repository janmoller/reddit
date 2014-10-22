package main

import (
	"fmt"
	"github.com/janmoller/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}

}
