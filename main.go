package main

import (
	"fmt"
	"time"

	"github.com/rancher/go-rancher-metadata/metadata"
)

func main() {
	client := metadata.NewClient("http://rancher-metadata/latest")
	var err error
	for err == nil {
		var c metadata.Container
		c, err = client.GetSelfContainer()
		if err == nil {
			fmt.Println(c.Name)
			time.Sleep(1 * time.Second)
		}
	}
	panic(err.Error())
}
