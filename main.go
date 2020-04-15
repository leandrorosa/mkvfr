package main

import (
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"github.com/k3a/html2text"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	println("Searching for "+os.Args[1])
	html := callKvr(os.Args[1])
	parsedHtml := parseHTML(html)
	fmt.Println(translate(parsedHtml))
}

func callKvr(kvrNumber string) string {
	form := url.Values{}
	form.Add("zapnummer", kvrNumber)
	form.Add("pbAbfragen", "information desk")

	req, _ := http.NewRequest("POST", "https://www17.muenchen.de/EATWebSearch/Auskunft", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "JSESSIONID=9C7440328C7D6097FCA92636B4531D08; _et_coid=c9c5697e8b9086bf8846dfc164c3c32e")

	hc := http.Client{}
	resp, _ := hc.Do(req)

	bytes, _ := ioutil.ReadAll(resp.Body)
	err := resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func parseHTML(plainHtml string) string {
	return html2text.HTML2Text(plainHtml)
}

func translate(text string) (string, error) {
	return gt.Translate(text, "de", "en")
}
