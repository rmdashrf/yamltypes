package yamltypes

import (
	"net"
	"net/url"
)

type HostPort string

type URL url.URL

func (h *HostPort) UnmarshalYAML(u func(interface{}) error) error {
	return unmarshalStringAndValidate((*string)(h), u, func(s string) error {
		_, _, err := net.SplitHostPort(s)
		return err
	})
}

func (f *URL) UnmarshalYAML(u func(interface{}) error) error {
	var s string
	if err := u(&s); err != nil {
		return err
	}

	// Go URL parse seems to accept everything.
	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	*f = (URL)(*parsed)
	return nil
}
