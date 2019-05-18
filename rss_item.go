// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-18 22:36 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"time"
)

type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	Description string
	PubishDate  time.Time
}

type rssItemParsed struct {
	Title       string `xml:"title"`
	Source      string `xml:"source"`
	SourceURL   string `xml:"sourceURL"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubishDate  rfc822 `xml:"pubDate"`
}

func rssItemParsedToRssItemExported(input rssItemParsed) RssItem {
	return RssItem{
		Title:       input.Title,
		Source:      input.Source,
		SourceURL:   input.SourceURL,
		Link:        input.Link,
		Description: input.Description,
		PubishDate:  time.Time(input.PubishDate.Time),
	}
}
