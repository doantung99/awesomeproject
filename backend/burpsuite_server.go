package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func test(rw http.ResponseWriter, req *http.Request) {
	x := struct {
		Host string     `json:"host"`
		Port int        `json:"port"`
		Https bool 	    `json:"https"`
		Request []byte  `json:"rawRequest"`
		Path string		`json:"path"`
	}{}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &x)
	log.Println(string(x.Request))
	log.Println(x.Port)
	log.Println(x.Host)
	log.Println(x.Https)
	log.Println(x.Path)
	//conn := driver.ConnectMongoDB()
	//add := models.Request{
	//	Raw: x.Request,
	//	Host: x.Host,
	//	Port: x.Port,
	//	Protocol: x.Protocol,
	//}
	//id, err := request.NewMongoDBRepo(conn).CreateRequest(&add)
	//if err != nil{
	//	log.Println(err)
	//}
	//log.Println(id)
	//driver.DisconnectMongo(conn, context.Background())
}

func main() {
	http.HandleFunc("/api", test)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}