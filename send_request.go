package main

import (
	"awesomeProject/configs"
	"log"
	"net/http"
)

func SendRequest(req *http.Request) *http.Response {
	client := configs.BurpProxy()
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	//log.Println(resp)
	//AnalyzeResponse(resp)
	return resp
}