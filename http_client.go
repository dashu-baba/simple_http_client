package httpclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

type Request struct {
	Url     string
	Method  string
	Headers map[string]interface{}
}

func Execute(request Request) {

	switch request.Method {
	case "GET":
		{
			res, err := http.Get("http://www.google.com/robots.txt")
			if err != nil {
				log.Fatal(err)
			}
			robots, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)
			break
		}
	default:
		{
			panic("Not supported")
		}
	}
}
