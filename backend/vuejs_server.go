package main

import (
	"awesomeProject/driver"
	"awesomeProject/repository/request"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type header struct{
	Host  string
	Status string
}

func requestList(w http.ResponseWriter, r *http.Request) {
	var arr []*url.URL
	conn := driver.ConnectMongoDB()
	listreq, err := request.NewMongoDBRepo(conn).GetAllRequests()
	if err != nil {
		log.Fatalln(err)
	}
	for _, j := range listreq {
		readRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(j.Raw)))
		if err != nil {
			log.Fatalln(err)
		}
		readRequest.RequestURI = ""
		u, err := url.Parse(j.Protocol + "://" + j.Host + ":" + strconv.Itoa(j.Port) + readRequest.URL.Path)
		if err != nil {
			log.Fatalln(err)
		}
		readRequest.URL = u
		log.Println(readRequest.URL)
		arr = append(arr, readRequest.URL)
	}
	driver.DisconnectMongo(conn, context.Background())
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	jsonData, _ := json.Marshal(arr)
	w.Write(jsonData)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	conn := driver.ConnectMongoDB()
	listreq, err := request.NewMongoDBRepo(conn).GetAllRequests()
	if err != nil {
		log.Fatalln(err)
	}
	var datas []header
	for _, j := range listreq {
		readRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(j.Raw)))
		if err != nil {
			log.Fatalln(err)
		}
		readRequest.RequestURI = ""
		u, err := url.Parse(j.Protocol + "://" + j.Host + ":" + strconv.Itoa(j.Port) + readRequest.URL.Path)
		if err != nil {
			log.Fatalln(err)
		}
		readRequest.URL = u

		//AnalyzeRequest(readRequest)
		resp := SendRequest(readRequest)
		//AnalyzeResponse(resp)

		d := header{
			Host:  readRequest.Host,
			Status: resp.Status,
		}
		log.Println(d.Host)
		log.Println(d.Status)
		datas = append(datas, d)
		resp.Body.Close()
	}
	driver.DisconnectMongo(conn, context.Background())
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	log.Println(datas)
	jsonData, _ := json.Marshal(datas)
	log.Println(jsonData)
	// json.NewEncoder(w).Encode(datas)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/reqlist", requestList)
	http.ListenAndServe("localhost:8081", nil)
}
