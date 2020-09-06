package main

import (
	"fmt"
	"net/http"
)

func AnalyzeResponse(res *http.Response){
	//Header Analysis
	if len(res.Header) != 0 {
		fmt.Println("Header - H:")
		for k, v := range res.Header {
			for _, vv := range v {
				fmt.Println(k + " > " + vv)
			}
		}
	}

	//Cookie Analysis
	if len(res.Cookies()) != 0 {
		fmt.Println("Cookie - C:")
		for _,i := range res.Cookies() {
			fmt.Println(i.Name + " > " + i.Value)
		}
	}
}