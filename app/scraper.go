package app

import (
	"fmt"
	"github.com/gocolly/colly"
	"jt-scraper-api/app/models"
	"unicode/utf8"
)

const (
	japanTimes = "https://www.japantimes.co.jp"
)

func makeCollector() (collector *colly.Collector) {
	c := colly.NewCollector(
		colly.AllowedDomains("japantimes.co.jp", "www.japantimes.co.jp"),
		colly.MaxDepth(1))
	return c
}

func visitContent(collector *colly.Collector, e *colly.HTMLElement) {
	link := e.Attr("href")
	link = e.Request.AbsoluteURL(link)
	collector.Visit(link)
}

func makeArticle(e *colly.HTMLElement) (article *models.Article) {
	var images []models.Image
	title := e.ChildText("h1")
	credit := e.ChildText("p.credit")
	writer := e.ChildText("h5.writer")
	date := e.ChildAttr("time", "datetime")
	e.ForEach("ul.slides > li > figure .fresco", func(_ int, el *colly.HTMLElement) {
		if el.Attr("data-fresco-group") == "attachment-group" {
			imageUrl := el.Attr("href")
			caption := el.Attr("data-fresco-caption")
			image := models.Image{
				caption,
				imageUrl,
			}
			images = append(images, image)
		}
	})

	url := fmt.Sprintf("%s%s", e.Request.URL.Host, e.Request.URL.Path)

	r, size := utf8.DecodeLastRuneInString(title)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	if title[len(title)-1:] == "." {
		title = title[:len(title)-size]
	}

	content := e.ChildText("#jtarticle > p")
	data := &models.Article{
		Title:   title,
		Content: content,
		Credit:  credit,
		Writer:  writer,
		Url:     url,
		Date:    date,
		Images:  images,
	}
	return data
}

func ScrapeUrl(url string) (a *models.Article, e error) {
	c := makeCollector()
	articleCollector := c.Clone()

	var article *models.Article
	articleCollector.OnHTML("div.main", func(e *colly.HTMLElement) {
		article = makeArticle(e)
	})
	if err := articleCollector.Visit(url); err != nil {
		return nil, err
	}
	return article, nil
}

func ScrapeTodaysMainArticles() (a []*models.Article, e error) {
	var articles []*models.Article
	var article *models.Article
	c := makeCollector()
	articleCollector := c.Clone()
	// Lead stories
	c.OnHTML("div.lead-stories > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleCollector, e)
	})

	// Top stories
	c.OnHTML("div.top-stories > a.wrapper-link.top-story", func(e *colly.HTMLElement) {
		visitContent(articleCollector, e)
	})

	// Editor picks
	c.OnHTML("div.editors-picks > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleCollector, e)
	})

	/**
	  Japantimes <div.main-content> is separated into different <section>
	  Each section consist of a feature article <div.featured>
	  A subsection list section of articles that relates to that section <ul
	*/
	c.OnHTML("div.featured > > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleCollector, e)

	})

	c.OnHTML("ul.module_articles > li.index-loop-article > a", func(e *colly.HTMLElement) {
		visitContent(articleCollector, e)
	})
	/**
	End of section collector.
	*/

	articleCollector.OnHTML("div.main", func(e *colly.HTMLElement) {
		article = makeArticle(e)
	})

	c.Visit(japanTimes)
	articles = append(articles, article)
	return articles, nil
}
