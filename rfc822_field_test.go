// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-18 22:47 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRfc822parseIntoSelf(t *testing.T) {

	assert.Equal(t, (&rfc822{}).parseIntoSelf("Mon, 02 Jan 2006 15:04:05 -0700"), nil)
	assert.NotEqual(t, (&rfc822{}).parseIntoSelf("Mon, 02 Jan 2006 15:BB:05 -0700"), nil)

	valueItem := rfc822{}
	valueItem.parseIntoSelf("Mon, 02 Jan 2006 15:04:05 -0700")
	assert.Equal(t, valueItem.Time.String(), "2006-01-02 15:04:05 -0700 -0700")

}
