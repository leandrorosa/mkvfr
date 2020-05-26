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
	parsedHTML := parseHTML(html)
	fmt.Println(translate(parsedHTML))
}

func callKvr(kvrNumber string) string {
	form := url.Values{}
	form.Add("zapnummer", kvrNumber)

	req, _ := http.NewRequest("POST", "https://www17.muenchen.de/EATWebSearch/Auskunft", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	hc := http.Client{}
	resp, _ := hc.Do(req)

	bytes, _ := ioutil.ReadAll(resp.Body)
	err := resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func parseHTML(plainHTML string) string {
	return html2text.HTML2Text(plainHTML)
}

func translate(text string) (string, error) {
	return gt.Translate(text, "de", "en")
}
