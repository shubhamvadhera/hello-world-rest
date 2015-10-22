package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	defer req.Body.Close()
	jsonIn, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Panic@helloPost.ioutil.ReadAll")
		panic(err)
	}

	var reqi interface{}
	json.Unmarshal([]byte(jsonIn), &reqi)
	reqs := reqi.(map[string]interface{})
	//fmt.Println(reqs["name"])

	type outInterface struct {
		Greeting string `json:"greeting"`
	}
	var foo outInterface
	foo.Greeting = "Hello, " + reqs["name"].(string) + "!"

	jsonOut, _ := json.Marshal(foo)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", jsonOut)
	//fmt.Println("Response:", string(jsonOut), " 201 OK")
}
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", helloPost)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
