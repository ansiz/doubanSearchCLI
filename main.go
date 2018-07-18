package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	Name            string
	Author          string
	Publisher       string
	RatingNums      float64
	TotalComment    string
	TotalCommentNum int
	Cover           string
	Desc            string
}

type Response struct {
	Totol int
	Items []string
	Limit int
	More  bool
}

const (
	tplShowBook = `
	书名： 《%s》-- %s		评分：%.1f%s
	`
	searchURL = `https://www.douban.com/j/search?q=%s&start=%d&cat=1001`
)

func getDataBytes(searchVal string, start int, local bool) ([]byte, error) {
	if local {
		return ioutil.ReadFile("./test/test.json")
	}
	fetchURL := fmt.Sprintf(searchURL, url.QueryEscape(searchVal), start)
	res, err := http.Get(fetchURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return ioutil.ReadAll(res.Body)
}

func fetchData(searchVal string, start int, local bool) (*Response, error) {
	data, err := getDataBytes(searchVal, start, local)
	resp := &Response{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		log.Fatal("parse response data failed:", err)
		return nil, err
	}
	return resp, nil
}

func parseBookInfo(resp *Response) []Book {
	bookList := []Book{}
	for _, item := range resp.Items {
		var book Book
		r := bytes.NewReader([]byte(item))
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".rating-info").Each(func(i int, s *goquery.Selection) {
			selRateNums := s.Children().First().Next()
			rateNum, err := strconv.ParseFloat(selRateNums.Text(), 32)
			if err != nil {
				book.TotalComment = "（评价人数不足）"
				book.Publisher = selRateNums.Next().Text()
			} else {
				book.RatingNums = rateNum
				book.TotalComment = selRateNums.Next().Text()
				count := strings.Split(strings.Trim(book.TotalComment, "("), "人")
				if len(count) > 1 {
					countNum, _ := strconv.Atoi(count[0])
					book.TotalCommentNum = countNum
				}
				book.Publisher = selRateNums.Next().Next().Text()
			}

		})
		doc.Find(".nbg").Each(func(i int, s *goquery.Selection) {
			book.Name, _ = s.Attr("title")
		})
		bookList = append(bookList, book)
	}
	return bookList
}

func ExampleScrape() {
	var current int
	resp, err := fetchData("架构", current, false)
	if err != nil {
		log.Fatalf("fetch data failed: %v", err)
	}
	bookList := []Book{}
	for current < resp.Limit*5 {
		log.Printf("Total: %d, Current: %d~%d", resp.Totol, current, current+resp.Limit)
		bookSubList := parseBookInfo(resp)
		bookList = append(bookList, bookSubList...)
		current = current + resp.Limit
		resp, err = fetchData("架构", current, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current Book list length: %d", len(bookList))
	}
	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].RatingNums > bookList[j].RatingNums {
			return true
		}
		if bookList[i].RatingNums < bookList[j].RatingNums {
			return false
		}
		return bookList[i].TotalCommentNum > bookList[j].TotalCommentNum
	})

	for _, book := range bookList {
		fmt.Printf(tplShowBook, book.Name, book.Publisher, book.RatingNums, book.TotalComment)
	}

}

func main() {
	ExampleScrape()
}
