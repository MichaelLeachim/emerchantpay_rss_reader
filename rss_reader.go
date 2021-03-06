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

type rss struct {
	Channel struct {
		Items []rssItemParsed `xml:"item"`
	} `xml:"channel"`
}

func parseFeedByUrl(da feedGetter, url string) ([]RssItem, error) {
	res, err := da.Get(url)
	if err != nil {
		return []RssItem{}, err
	}

	feed := rss{}
	err = xml.Unmarshal([]byte(res), &feed)
	if err != nil {
		return []RssItem{}, err
	}
	result := []RssItem{}
	for _, item := range feed.Channel.Items {
		result = append(result, rssItemParsedToRssItemExported(item))
	}
	return result, nil
}

func parseFeedByUrlsAsync(da feedGetter, l logger, urls []string) []RssItem {
	feedChan := make(chan RssItem)

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	for _, link := range urls {
		go func(url string) {
			defer wg.Done()
			items, err := parseFeedByUrl(da, url)
			if err != nil {
				l.Warn(err)
				return
			}
			for _, item := range items {
				feedChan <- item
			}
		}(link)
	}

	results := []RssItem{}
	go func() {
		for item := range feedChan {
			results = append(results, item)
		}
	}()

	wg.Wait()

	return results
}

func Parse(urls []string) []RssItem {
	return parseFeedByUrlsAsync(newHttpFeedGetter(), newLogrusLogger(), urls)
}
