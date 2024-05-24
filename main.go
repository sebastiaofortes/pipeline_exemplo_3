package main

import (
	"encoding/json"
	"log"
)

func main(){

	num := 12

	r, _ := json.Marshal(num)

	log.Println(r)

}

func Soma(a, b int) int {
	return a + b
}