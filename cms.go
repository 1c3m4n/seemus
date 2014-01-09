package main

import (
	"code.google.com/p/gorest"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Article struct {
	Title   string
	Date    string
	Content []string
}

func main() {
	fmt.Println("Starting up")
	fileContent, fileErr := ioutil.ReadFile("test.json")
	if fileErr != nil {
		fmt.Println("Error reading file : " + fileErr.Error())
	}
	var m Article

	err := json.Unmarshal(fileContent, &m)
	if err != nil {
		fmt.Println("Error unmarshalling file contents to JSON : " + err.Error())
	}
	fmt.Println(m.Content)
	gorest.RegisterService(new(ContentService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8787", nil)
	fmt.Println("Shutting Down")
}

type ContentService struct {
	gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`
	getArticle         gorest.EndPoint `method:"GET" path:"/articles/{ArticleName:string}" output:"string"`
}

func (serv ContentService) GetArticle(ArticleName string) (out string) {
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}
