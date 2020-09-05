package configs

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	localCertFile = "C:\\Users\\Tu`ng-Chann\\Downloads\\cacert.der"
)

func BurpProxy() *http.Client {
	//Set up Proxy and Certificate
	caCert, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	//proxy, err := url.Parse("http://127.0.0.1:8080")
	//if err != nil {
	//	log.Fatal(err)
	//}

	config := &tls.Config{
		InsecureSkipVerify: true,
		//RootCAs:            caCertPool,
	}
	tr := &http.Transport{
		TLSClientConfig: config,
		//Proxy:           http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(20 * time.Second),
	}

	return client
}