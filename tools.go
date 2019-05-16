// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-17 00:33 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	"encoding/xml"
	"time"
)

type rfc822 struct {
	time.Time
}

func newRfc822(input string) (rfc822, error) {

}
func (c *rfc822) parseTime(input string) (time.Time, error) {
	const RFC822 = "Mon, 02 Jan 2006 15:04:05 -0700"
	parse, err := time.Parse(RFC822, v)
	if err != nil {
		return err
	}

}

func (c *rfc822) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string
	d.DecodeElement(&v, &start)

	*c = rfc822{parse}
	return nil
}
