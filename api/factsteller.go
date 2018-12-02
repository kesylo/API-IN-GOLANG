package api

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Getfact() string {

	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(3)
	var body string
	// fmt.Println(random)

	switch random {
	case 0:
		body = RandomFact("http://numbersapi.com/random/trivia")
	case 1:
		body = RandomFact("http://numbersapi.com/random/math")
	case 2:
		body = RandomFact("http://numbersapi.com/random/date")
	}

	return body
}

func RandomFact(link string) string {

	resp, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
