package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

var (
	client         *http.Client
	lawsonInterval = 5
	lawsonHost     = "http://www.lawson.com.cn/promotions"
)

func init() {
	client = &http.Client{}
}

func CheckErr(err error, funcName string) {
	if err != nil {
		log.Fatal(funcName, err.Error())
	}
}

func GetLawsonUrls() {
	request, err := http.NewRequest("GET", lawsonHost, nil)
	CheckErr(err, "request")
	response, err := client.Do(request)
	CheckErr(err, "response")

	buf, err := ioutil.ReadAll(response.Body)
	CheckErr(err, "ReadAll")
	d, err := goquery.NewDocumentFromReader(strings.NewReader(string(buf)))
	//fmt.Println(string(buf))
	CheckErr(err, "NewDocumentFromReader")

	pages := make(map[string]int) //网页的url和出现的频率
	d.Find(".pagination").Find("li").Each(func(i int, s *goquery.Selection) {
		dd, _ := s.Find("a").Attr("href")
		pages[dd] += 1
		//fmt.Println(dd)
	})
	fmt.Println(pages)

	d.Find(".col-xs-4").Each(func(i int, s *goquery.Selection) {
		dd, _ := s.Find("a").Attr("href")
		fmt.Println(dd)
	})

	//fmt.Println(string(buf))
}

// main for test
func main() {
	GetLawsonUrls()
}

// func Start() {

// }

// func GetCrawlerResult(url string) ([]byte, error) {
// 	request, err := http.NewRequest("GET", url, nil)
// 	CheckErr(err, "request")
// 	response, err := client.Do(request)
// 	CheckErr(err, "response")
// 	fmt.Println(response.Header)
// 	buf, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf, nil
// }
