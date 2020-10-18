package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
)

func main532() {
    c := colly.NewCollector()
    c.OnHTML(".wikitable", func(e *colly.HTMLElement) {
        e.DOM.Children().Find("tr>td").Parent().Each(func(_ int, s *goquery.Selection) {
            symbols := s.Find("td")
            fmt.Println(symbols.Eq(0).Text(), "0x" + symbols.Eq(1).Text())
        })
    })
    c.Visit("https://en.wikipedia.org/wiki/Java_bytecode_instruction_listings")
}
