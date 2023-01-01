package CastHunter

import (
	"bytes"
	"fmt"
	"net/http"
)

func NewPostNoData(uri string, roku bool) {
	data := []byte(``)
	request, x := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if x != nil {
		fmt.Println("Could not make a request -> ", x)
		return
	}
	client := &http.Client{}
	response, x := client.Do(request)
	if x != nil {
		fmt.Println("Could not fufil the request -> ", x)
		return
	}
	if response.StatusCode == 200 {
		OK[100]()
	}
	defer response.Body.Close()
}
