package main

import (
	"fmt"
	"github.com/sternix/kurlar"
	"log"
)

func main() {
	kur, err := kurlar.Today()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(kur)
}
