package util

import (
	"fmt"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		log.Fatal(err)
	}
}
