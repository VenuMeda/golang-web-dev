package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}