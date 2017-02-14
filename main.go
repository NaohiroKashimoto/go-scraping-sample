package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"golang.org/x/text/encoding/japanese"
)

func main() {
	//goqueryexample()
	loginvc()
}

func loginvc() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal(err.Error())
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := page.Navigate("https://www.valuecommerce.ne.jp/"); err != nil {
		log.Fatal(err.Error())
	}

	loginUrl, err := page.URL()
	if err != nil {
		log.Fatal(err.Error(), loginUrl)
	}

	page.Find("iframe#frame1").SwitchToFrame()
	fmt.Println(page.HTML())

	page.Find("input#login_form_emailAddress").Fill("osaifu@ceres-inc.jp")
	page.Find("input#login_form_encryptedPasswd").Fill("r8nzfrnz")

	if err := page.Find("input.btn_green").Click(); err != nil {
		log.Fatal(err.Error())
	}

	page.Screenshot("test")
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
