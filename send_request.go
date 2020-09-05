package main

import (
	"awesomeProject/configs"
	"log"
	"net/http"
)

func SendRequest(req *http.Request) {
	client := configs.BurpProxy()
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)

	defer resp.Body.Close()
}