package main

import (
	"go-weather/cmd"
	"log"
)

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
