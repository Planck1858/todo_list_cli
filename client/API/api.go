package API

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ToDo struct {
	Index     int       `json:"index"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"created"`
}

type List struct {
	Tasks []ToDo `json:"tasks"`
}

func TryConnection() string {
	body := getData(makeURL("/"))

	var data string
	data = string(body)

	return data
}

func GetList() List {
	body := getData(makeURL("/getList"))

	var data List
	json.Unmarshal(body, &data.Tasks)

	return data
}

func NewTask(data string) {
	text := []byte(data)
	_, err := http.Post(makeURL("/newTask"), "application/json", bytes.NewBuffer(text))
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteList() {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", makeURL("/clearList"), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func DeleteTask(index string) {
	client := &http.Client{}
	text := "/deleteTask/" + index

	req, err := http.NewRequest("DELETE", makeURL(text), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func ChangeTask(index string) {
	client := &http.Client{}
	text := "/changeTask/" + index

	req, err := http.NewRequest("PUT", makeURL(text), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func makeURL(text string) string {
	return fmt.Sprintf("http://localhost:8080%s", text)
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
