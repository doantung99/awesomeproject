package main

import (
	"awesomeProject/driver"
	"awesomeProject/repository/request"
	"bufio"
	"bytes"
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main()  {
	conn := driver.ConnectMongoDB()
	//reqList, err := request.NewMongoDBRepo(conn).GetAllRequests()
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//for _, j := range reqList {
	//	fmt.Println(j.Id)
	//}

	listreq, err := request.NewMongoDBRepo(conn).GetAllRequests()
	if err != nil{
		log.Println(err)
	}

	for _, j := range listreq {
		readRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(j.Raw)))
		if err != nil{
			log.Println(err)
		}
		//log.Printf("%v", readRequest)
		//log.Println(readRequest)
		//log.Println(j.Protocol + "://" + j.Host + ":" + strconv.Itoa(j.Port) + readRequest.URL.Path)
		//log.Println(reflect.TypeOf(readRequest))
		readRequest.RequestURI = ""
		u, err := url.Parse(j.Protocol + "://" + j.Host + ":" + strconv.Itoa(j.Port) + readRequest.URL.Path)
		if err != nil {
			log.Println(err)
		}
		readRequest.URL = u

		AnalyzeRequest(readRequest)
		resp := SendRequest(readRequest)
		AnalyzeResponse(resp)
		resp.Body.Close()
	}
	driver.DisconnectMongo(conn, context.Background())
}
