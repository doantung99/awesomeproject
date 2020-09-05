package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func JSONAnalysis(value interface{}) interface{} {
	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Bool:
			return value.(bool)
		case reflect.String:
			return value.(string)
		case reflect.Slice:
			fmt.Println(value)
			d := value.([]interface{})
			for _, j := range d {
				fmt.Print("j la ")
				fmt.Println(JSONAnalysis(j))
			}
		case reflect.Map:
			fmt.Println(value)
			d := value.(map[string]interface{})
			for i, j := range d {
				fmt.Println("i la " + i)
				fmt.Print("j la ")
				fmt.Println(JSONAnalysis(j))
			}
		case reflect.Float64:
			return value.(float64)
		default:
			return "Another Type"
		}
	} else {
		return nil
	}
	return ""
}

func AnalyzeRequest(request *http.Request) {
	//body, _ := ioutil.ReadAll(request.Body)

	//Path Analysis
	fmt.Println("Path - U:")
	fmt.Println(request.URL.Path)
	for _, path := range strings.Split(request.URL.Path, "/") {
		if path != "" {
			fmt.Println("(path) > " + path)
		}
	}

	//Header Analysis
	if len(request.Header) != 0 {
		fmt.Println("Header - H:")
		for k, v := range request.Header {
			for _, vv := range v {
				fmt.Println(k + " > " + vv)
			}
		}
	}

	//Cookie Analysis
	if len(request.Cookies()) != 0 {
		fmt.Println("Cookie - C:")
		for _, i := range request.Cookies() {
			fmt.Println(i.Name + " > " + i.Value)
		}
	}

	//GET Data Analysis
	if len(request.URL.Query()) != 0 {
		fmt.Println("GET Data - G:")
		for name, query := range request.URL.Query() {
			for _, value := range query {
				fmt.Println(name + " > " + value)
			}
		}
	}

	//POST Data Analysis
	var result map[string]interface{}
	json.NewDecoder(request.Body).Decode(&result)
	if len(result) != 0 {
		fmt.Println("POST Data - P:")
		fmt.Println(result)
		for i, j := range result {
			fmt.Println("---------------------")
			fmt.Println("i la " + i)
			fmt.Print("j la ")
			fmt.Println(JSONAnalysis(j))
		}
	}
}
