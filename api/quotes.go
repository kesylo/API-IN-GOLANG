package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetQuotes() string {
	var quotesBody string
	var quotesAuthor string
	//var aQuote string
	res := GetRequest("https://talaikis.com/api/quotes/")
	s := ReadBody(res)
	m := ReadJSON(s)
	quotesBody = m.Quote
	quotesAuthor = m.Author
	aQuote := "\n" + quotesBody + "\n\n\n " + quotesAuthor
	//return quotesBody + quotesAuthora
	return aQuote
}

func GetQuotesB(n int) string {

	//var quotesBody string
	//var quotesAuthor string
	//var aQuote string
	var author string
	var quote string
	res := GetRequest("https://talaikis.com/api/quotes/")
	s := ReadBody(res)
	m := ReadJSON(s)
	if n == 1 {
		author = m.Quote
		return author
	}
	if n == 2 {
		quote = m.Author
		return quote
	}
	//quotesBody = m.Quote
	//quotesAuthor = m.Author
	//quote = m.Quote
	//author = m.Author
	return "" //quote + author
}

/*
func GetQuotesAuthor() string {

	quotesBody, quotesAuthor := GetQuotes()

	return quotesAuthor + quotesBody
}


func GetImage() {

	fileUrl := "https://en.wikipedia.org/wiki/Albert_Einstein"

	err := DownloadFile("avatar.jpg", fileUrl)
	if err != nil {
		panic(err)
	}

}*/
