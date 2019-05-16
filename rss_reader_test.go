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
)

func TestParseFeedByUrl(t *testing.T) {
	mockfeed := newMockFeedGetter(0)

	feed, err := parseFeedByUrl(mockfeed, "https://news.ycombinator.com/rss")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(feed), 10)
	assert.Equal(t, feed[0], "")
}
