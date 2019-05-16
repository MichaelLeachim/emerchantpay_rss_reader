// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 21:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"encoding/xml"
	"sync"
)

type RssItem struct {
	Title       string `xml:"title"`
	Source      string `xml:"source"`
	SourceURL   string `xml:"sourceURL"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubishDate  rfc822 `xml:"pubDate"`
}

type Rss struct {
	Channel struct {
		Items []RssItem `xml:"item"`
	} `xml:"channel"`
} // `xml:"rss"`

func parseFeedByUrl(da FeedGetter, url string) ([]RssItem, error) {
	res, err := da.Get(url)
	if err != nil {
		return []RssItem{}, err
	}

	feed := Rss{}
	err = xml.Unmarshal([]byte(res), &feed)
	if err != nil {
		return []RssItem{}, err
	}
	return feed.Channel.Items, nil
}

func parseFeedByUrlsAsync(da FeedGetter, l Logger, urls []string) []RssItem {
	feedChan := make(chan RssItem)
	results := []RssItem{}
	go func() {
		for item := range feedChan {
			results = append(results, item)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	for _, link := range urls {
		go func() {
			defer wg.Done()
			items, err := parseFeedByUrl(da, link)
			if err != nil {
				l.Warn(err)
				return
			}
			for _, item := range items {
				feedChan <- item
			}
		}()
	}

	wg.Wait()
	return results
}

func Parse(urls []string) []RssItem {
	return parseFeedByUrlsAsync(newHttpFeedGetter(), newLogrusLogger(), urls)
}
