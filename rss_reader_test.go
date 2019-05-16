// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 22:37 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package emerchantpay_rss_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseFeedByUrl(t *testing.T) {
	mockfeed := newMockFeedGetter(0)

	feed, err := parseFeedByUrl(mockfeed, "https://news.ycombinator.com/rss")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(feed), 30)
	assert.Equal(t, feed[0],
		RssItem{Title: "Sony and Microsoft set rivalry aside for streaming alliance",
			Source:      "",
			SourceURL:   "",
			Link:        "https://asia.nikkei.com/Business/Business-deals/Sony-and-Microsoft-set-rivalry-aside-for-streaming-alliance",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19930955\">Comments</a>",
			PubishDate:  rfc822{Time: time.Time{wall: 0x0, ext: 63693624469, loc: (*time.Location)(0xc000062900)}}})

}

func genMockLinks(amount int) []string {
	data := []string{"https://www.theguardian.com/uk/rss", "https://news.ycombinator.com/rss", "not.existing.url"}
	result := []string{}
	for i := 0; i < amount/3; i++ {
		result = append(result, data...)
	}
	return result
}

func BenchmarkParseFeedByUrlsAsync(t *testing.B) {
	da := newMockFeedGetter(1000)
	logger := newMockLogger()
	mocklinks := genMockLinks(1000)
	for n := 0; n < t.N; n++ {
		parseFeedByUrlsAsync(da, logger, mocklinks)
	}
}

func TestParseFeedByUrlsAsync(t *testing.T) {

	da := newMockFeedGetter(100)
	logger := newMockLogger()
	mocklinks := genMockLinks(1000)
	feed := parseFeedByUrlsAsync(da, logger, mocklinks)
	assert.Equal(t, len(feed), 990)

}
