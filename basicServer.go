package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main(){
	fmt.Println("Hello world");
	// getRequest()
	postRequest()
}

func getRequest(){
	url := ""
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	// var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// byteCount, _ := responseString.Write(content)

	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())
	fmt.Println(string(content))
}

func postRequest(){
	url := ""

	requestBody := strings.NewReader(`
			{
				"email":"yuvraj@gmail.com",
				"first_name":"yuvraj",
				"last_name":"rathva",
				"avatar":"https://sample.in/img/faces/2-image.jpg",
			}
		`)

	response, err := http.Post(url, "application/json", requestBody)
	if err!=nil{panic(err)}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}