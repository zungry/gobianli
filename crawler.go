package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fmt"

	"time"

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

func Lawson() {
	//Get all Urls need to be crawled
	GetLawsonUrls()
	//Get all items which is off
	// save to ES
}

func Quanjia() {

}

func Kuaike() {

}

func CrawlUrlByGet(url string) (*goquery.Document, error) {
	request, err := http.NewRequest("GET", url, nil)
	CheckErr(err, "request")
	response, err := client.Do(request)
	CheckErr(err, "response")

	buf, err := ioutil.ReadAll(response.Body)
	CheckErr(err, "ReadAll")
	return goquery.NewDocumentFromReader(strings.NewReader(string(buf)))
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
		if dd == "" {
			pages["/promotions?page=1"] = 1
		} else {
			pages[dd] += 1
		}

		//fmt.Println(dd)
	})
	for k, v := range pages {
		fmt.Println(k, v)
	}
	fmt.Println(pages)

	d.Find(".col-xs-4").Each(func(i int, s *goquery.Selection) {
		dd, _ := s.Find("a").Attr("href")
		fmt.Println(dd)
	})

	//fmt.Println(string(buf))
}

func start() {
	Lawson()
	now := time.Now()
	next := now.Add(time.Hour * 24) // Compute next zero clock
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	<-t.C
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
