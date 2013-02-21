package main

import (
	"fmt"
	"net/http"
        "net/url"
	)

func main() {
  resp, _ := http.PostForm("http://localhost:3000",
	url.Values{"key": {"Value"}, "id": {"123"}})
	fmt.Println(resp)
}
