package main

import (
	"log"

	"fmt"

	"github.com/sclevine/agouti"
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
