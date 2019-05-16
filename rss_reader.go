// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 21:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"encoding/xml"
	"time"
)

type RssItem struct {
	Title       string    `xml:"title"`
	Source      string    `xml:"source"`
	SourceURL   string    `xml:"source"`
	Link        string    `xml:"link"`
	PubishDate  time.Time `xml:"pubDate"`
	Description string    `xml:"description"`
}
type RssFeed struct {
	Rss struct {
		Items []RssItem `xml:"item"`
	} `xml:"rss"`
}

func parseInput(input string) ([]RssItem, error) {
	feed := RssFeed{}
	err := xml.Unmarshal([]byte(input), &feed)
	if err != nil {
		return []RssItem{}, err
	}
	return feed.Rss.Items, nil
}

func Parse(urls []string) []RssItem {
	return []RssItem{}
}
