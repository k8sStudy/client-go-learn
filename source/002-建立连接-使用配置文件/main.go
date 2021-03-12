package main

import (
	"../../config"
	"fmt"
	"log"
)

func main() {
	clientSet, err := config.InitClientConfigFile()
	if err != nil {
		log.Fatal("err:\t", err)
		return
	}
	fmt.Printf("%#v\n", clientSet)
}
