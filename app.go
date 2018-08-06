package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
)

// InputJSON 入力データ
type InputJSON struct {
	URL string `json:"url`
}

// Item 各商品ブロックのデータ
type Item struct {
	URL   string `json:"url"`
	Image string `json:"image"`
	Title string `json:"title"`
}

// OutputJSON 出力データ
type OutputJSON []Item

func getYahooNewItems(target InputJSON) (OutputJSON, error) {
	var url, image, title string
	var outputs OutputJSON

	html, _ := goquery.NewDocument(target.URL)
	html.Find(".elWrap").Each(func(_ int, block *goquery.Selection) {
		block.Find(".elImage").Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ = s.Attr("href")
		})
		block.Find(".elImage").Find("img").Each(func(_ int, s *goquery.Selection) {
			image, _ = s.Attr("src")
		})
		block.Find(".elImage").Find("img").Each(func(_ int, s *goquery.Selection) {
			title, _ = s.Attr("alt")
		})
		item := Item{
			URL:   url,
			Image: image,
			Title: title,
		}
		outputs = append(outputs, item)
	})

	return outputs, nil
}

func main() {
	lambda.Start(getYahooNewItems)
}
