package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:8080/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := []byte(`{"name": "John", "age": 30}`)
	http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))

	fmt.Println(resp.Status)

	//fmt.Println(string(body))
}
