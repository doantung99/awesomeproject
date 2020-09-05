package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	data, err := json.Marshal(map[string]interface{}{
		"id":   1,
		"name": "Doan Le Manh Tung",
		"detail": map[string]interface{}{
			"major": map[string]interface{}{
				"subject1": "IAO",
				"subject2": "IAA",
			},
			"class": "1301",
		},
		"check": true,
		"array1": []map[string]interface{}{
			{"a":1},
			{"b":2},
			{"c":3},
		},
		"array2": []int{1, 2, 3},
		"test": nil,
		"float": 1.32,
	})

	if err != nil {
		log.Fatalln(err)
	}
	//data := []byte("")

	//Create new Request
	request, err := http.NewRequest("GET", "https://api.unsplash.com/photos?data=test1&utf-8=%E2%9C%93&commit=%E6%96%B0%E8%A6%8F%E7%99%BB%E9%8C%B2", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println(string(body))
}
