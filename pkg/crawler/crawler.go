package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
)

const (
	tplShowBook = `
	书名： 《%s》-- %s		评分：%.1f%s
	`
	searchBookURL = `https://www.douban.com/j/search?q=%s&start=%d&cat=1001`
)

// Crawler represents the data crawler.
type Crawler struct {
	Options
}

// Options contains the crawler supported options.
type Options struct {
	Verbose   bool
	Keyword   string
	Start     int
	Page      int
	LocalMode bool
	LocalFile string
	Interval  int
}

// Book represents the book information.
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

// Response represents the HTTP response data struct from Douban.
type Response struct {
	Total int
	Items []string
	Limit int
	More  bool
}

// New creates crawler instance.
func New() (*Crawler, error) {
	return &Crawler{}, nil
}

// SearchList lists books.
func (c *Crawler) SearchList() error {
	if c.Verbose {
		log.SetLevel(log.DebugLevel)
	}
	resp, err := c.fetchData()
	if err != nil {
		log.Fatalf("fetch data failed: %v", err)
		return err
	}
	bookList := []Book{}
	for c.Start < resp.Limit*c.Page {
		log.Debugf("total: %d, current: %d~%d",
			resp.Total, c.Start, c.Start+resp.Limit)
		bookSubList := parseBookInfo(resp)
		bookList = append(bookList, bookSubList...)
		c.Start = c.Start + resp.Limit
		resp, err = c.fetchData()
		if err != nil {
			return err
		}
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

	return nil

}

func (c *Crawler) fetchData() (*Response, error) {
	data, err := c.getDataBytes()
	resp := &Response{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		log.Fatal("parse response data failed:", err)
		return nil, err
	}
	return resp, nil
}

func (c *Crawler) getDataBytes() ([]byte, error) {
	if c.LocalMode {
		return ioutil.ReadFile(c.LocalFile)
	}
	fetchURL := fmt.Sprintf(searchBookURL, url.QueryEscape(c.Keyword), c.Start)
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
