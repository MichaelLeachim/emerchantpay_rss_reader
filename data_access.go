// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 21:55 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"io/ioutil"
	"net/http"
)

type feedGetter interface {
	Get(string) (string, error)
}

type httpFeedGetter struct{}

func newHttpFeedGetter() feedGetter {
	return httpFeedGetter{}
}

func (w httpFeedGetter) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
