package CastHunter

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		if roku {
			fmt.Println("[Information] Success: Roku box at ( ", TargetMain, " ) has been sent the buffer | RESP 200")
		} else {
			fmt.Println("[Information] Success: Hostname at ( ", TargetMain, " ) has been sent the buffer | RESP 200")
		}
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
