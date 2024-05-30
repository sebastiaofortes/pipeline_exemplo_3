package main

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello sexta"))
		if err != nil{
			fmt.Println(err)
		}
	})

	num := 12

	r, _ := json.Marshal(num)

	log.Println(r)
  
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println(err)
	}
}

func Soma(a, b int) int {
	return a + b
}