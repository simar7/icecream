package main

import (
	"sync"

	"github.com/simar7/icecream/pkg/publish"
	"github.com/simar7/icecream/pkg/subscribe"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go publish.Pub()
	go subscribe.Sub()
	wg.Wait()
}
