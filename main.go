package main

import (
	"log"

	"fmt"

	"time"

	"github.com/sclevine/agouti"
)

func main() {
	//goqueryexample()
	//loginvc()
	linkshare()
}
func linkshare() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal(err.Error())
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := page.Navigate("https://www.linkshare.ne.jp/"); err != nil {
		log.Fatal(err.Error())
	}

	loginUrl, err := page.URL()
	if err != nil {
		log.Fatal(err.Error(), loginUrl)
	}

	//トップページからログインするまで
	//page.Find("iframe#frame1").SwitchToFrame()
	fmt.Println(page.HTML())

	//ログインフォーム
	t, err := page.Find("#loginboxset > tbody > tr > td:nth-child(1) > div > form").Attribute("target")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(t)

	page.Find("#Lid").Fill("opencube")
	page.Find("#Lpas").Fill("YDaa9NyHV4tr")

	if err := page.Find("#loginboxset > tbody > tr > td:nth-child(1) > div > form > table > tbody > tr:nth-child(2) > td.Tdbgc.lslogintabletd3 > div > input[type=\"submit\"]").Click(); err != nil {
		log.Fatal(err.Error())
	}

	//リンクシェアは_blankでやっちゃう。
	page.NextWindow()

	time.Sleep(3 * time.Second)

	//ログイン処理完了後のトップページのスクリーンショット
	page.Screenshot("test2")
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

	//トップページからログインするまで
	page.Find("iframe#frame1").SwitchToFrame()
	fmt.Println(page.HTML())

	page.Find("input#login_form_emailAddress").Fill("osaifu@ceres-inc.jp")
	page.Find("input#login_form_encryptedPasswd").Fill("r8nzfrnz")

	if err := page.Find("input.btn_green").Click(); err != nil {
		log.Fatal(err.Error())
	}

	//ログイン処理完了後のトップページのスクリーンショット
	page.Screenshot("test")
}
