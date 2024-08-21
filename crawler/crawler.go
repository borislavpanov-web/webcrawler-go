package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"io"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func ScrapePage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	utf8Reader, enc, err := determineEncoding(resp.Body)
	if err != nil {
		return "", err
	}

	if enc == nil {
		utf8Reader = transform.NewReader(resp.Body, encoding.Nop.NewDecoder())
	}

	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	doc.Find(".item").Each(func(index int, item *goquery.Selection) {
		title := item.Find(".title").Text()
		price := strings.TrimSpace(item.Find(".price").Text())
		params :=  strings.TrimSpace(item.Find(".params").Text())
		formattedItem := fmt.Sprintf("Title: %s\nPrice: %s\n Params: %s\n", title, price, params)
		fmt.Fprintln(&result, formattedItem)
	})

	return result.String(), nil
}

func ScrapeAllPages(startPage, endPage int) (string, error) {
	var allResults []string
	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make(chan error, endPage-startPage+1)

	for i := endPage; i >= startPage; i-- {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			url := fmt.Sprintf("https://www.mobile.bg/obiavi/avtomobili-dzhipove/mercedes-benz/p-%d", pageNum)
			data, err := ScrapePage(url)
			if err != nil {
				errors <- err
				return
			}
			mu.Lock()
			allResults = append([]string{data}, allResults...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		return "", fmt.Errorf("encountered errors during scraping")
	}

	var finalResult strings.Builder
	for _, result := range allResults {
		finalResult.WriteString(result)
	}

	return finalResult.String(), nil
}

func determineEncoding(r io.Reader) (io.Reader, encoding.Encoding, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	enc, name, _ := charset.DetermineEncoding(bytes, "")
	log.Printf("Detected charset: %s", name)
	return transform.NewReader(strings.NewReader(string(bytes)), enc.NewDecoder()), enc, nil
}
