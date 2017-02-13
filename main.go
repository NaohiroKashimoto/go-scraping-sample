package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
)

func main() {
	goqueryexample()
}

func goqueryexample() {

	response, err := http.Get("http://1.dev.osaifu2.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	//ShiftJISの処理（ShiftJIS->UTF-8）
	decoder := japanese.ShiftJIS.NewDecoder()
	text, err := decoder.String(document.Find("html body#top div#container div#content.clearfix div.main-column div.section.feature.item-block ul.clearfix li a.new span.detail span.name").Text())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(text)
}
