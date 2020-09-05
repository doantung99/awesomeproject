package main

import (
	"awesomeProject/driver"
	"awesomeProject/models"
	"awesomeProject/repository/request"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func test(rw http.ResponseWriter, req *http.Request) {
	x := struct {
		Data []byte `json:"data"`
	}{}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &x)
	log.Println(string(x.Data))
	conn := driver.ConnectMongoDB()
	add := models.Request{
		Raw: x.Data,
	}
	id, err := request.NewMongoDBRepo(conn).CreateRequest(&add)
	if err != nil{
		log.Println(err)
	}
	log.Println(id)
	driver.DisconnectMongo(conn, context.Background())
}

func main() {
	http.HandleFunc("/api", test)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}