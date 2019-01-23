package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

func GetQuoteV2(n int) string {
	var auteur string
	var citation string
	//var aQuote string
	var message string
	response, err := http.Get("http://quotes.rest/qod.json?category=students")
	//response, err := http.Get("https://andruxnet-random-famous-quotes.p.rapidapi.com")
	checkError(err)

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	checkError(err)

	quote := string(contents)

	valueOfQuote := gjson.Get(quote, "contents.quotes.0.quote")
	valueOfAuthor := gjson.Get(quote, "contents.quotes.0.author")
	valueOfBackground := gjson.Get(quote, "contents.quotes.0.background")
	valueOfTiltle := gjson.Get(quote, "contents.quotes.0.title")
	valueOfMsg := gjson.Get(quote, "message")
	if n == 1 {
		auteur = valueOfAuthor.String()
		return auteur
	}
	if n == 2 {
		citation = valueOfQuote.String()
		return citation
	}
	if n == 3 {
		imageLink := valueOfBackground.String()
		return imageLink
	}
	if n == 4 {
		titre := valueOfTiltle.String()
		return titre
	}
	if n == 5 {
		message := valueOfMsg.String()
		return message
	}

	println(valueOfAuthor.String())
	println(valueOfBackground.String())
	println(valueOfMsg.String())
	url := "https://en.wikipedia.org/wiki/" + valueOfAuthor.String()
	//aQuote = citation + auteur // + url
	getImage(url)
	message = valueOfMsg.String()
	return message

	//return aQuote
}

func checkError(err error) {
	if err != nil {
		//fmt.Println(err)
		os.Exit(1)
	}
}

func getImage(url string) string {

	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(url)
	checkError(err)
	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	defer res.Body.Close()

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// You can now save it to disk or whatever...
	ioutil.WriteFile("/home/avni/gocode/src/API-IN-GOLANG/static/img-avni", data, 0666)

	log.Println("I saved your image buddy!")
	return url
}
