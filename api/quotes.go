package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Quote, Author, Category string
}

func GetRequest(url string) *http.Response {
	res, err := http.Get("https://talaikis.com/api/quotes/random")

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func ReadBody(res *http.Response) string {
	byteArr, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(byteArr)
}

func ReadJSON(s string) Message {
	var m Message

	dec := json.NewDecoder(strings.NewReader(s))

	err := dec.Decode(&m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}

func GetQuotes() (string, string) {

	var quotesBody string
	var quotesAuthor string
	var aQuote string
	res := GetRequest("https://talaikis.com/api/quotes/")

	s := ReadBody(res)

	m := ReadJSON(s)

	quotesBody = m.Quote
	quotesAuthor = m.Author

	aQuote = string("\n" + quotesBody + "\n" + quotesAuthor)

	return "", aQuote

}

/*
func GetQuotesBody() string {
	var quotesBody string
	var quotesAuthor string
	quotesBody, quotesAuthor := GetQuotes()
	res := GetRequest("https://talaikis.com/api/quotes/")
	s := ReadBody(res)
	m := ReadJSON(s)
	quotesBody, quotesAuthor := GetQuotes()

	return string(quotesBody + quotesAuthor)
}
*/
/*
func GetQuotesAuthor() string {

	quotesBody, quotesAuthor := GetQuotes()

	return quotesAuthor + quotesBody
}
*/
