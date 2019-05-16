// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 22:12 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"errors"
	"io/ioutil"
	"time"
)

type MockFeedGetter struct {
	delay int
}

func newMockFeedGetter(delay int) FeedGetter {
	return MockFeedGetter{delay: delay}
}

func (s MockFeedGetter) Get(url string) (string, error) {
	rsslink := ""

	switch url {
	case "https://news.ycombinator.com/rss":
		rsslink = "testdata/news.ycombinator.com.rss"
	case "https://www.theguardian.com/uk/rss":
		rsslink = "theguardian.rss"
	}
	time.Sleep(time.Duration(s.delay))

	if len(rsslink) == 0 {
		return "", errors.New("Cannot open URL: " + url)
	}

	bytes, err := ioutil.ReadFile(rsslink)
	if err != nil {
		panic("Directory ./testdata should exist in order for tests to run correctly")
	}
	return string(bytes), nil
}
