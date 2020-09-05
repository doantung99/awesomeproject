package main

import (
	"awesomeProject/driver"
	"awesomeProject/repository/request"
	"bufio"
	"bytes"
	"context"
	"log"
	"net/http"
	"reflect"
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
		log.Println(readRequest)
		log.Println(reflect.TypeOf(readRequest))

		//AnalyzeRequest(readRequest)
		SendRequest(readRequest)
	}
	driver.DisconnectMongo(conn, context.Background())
}
