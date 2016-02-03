package yamltypes

import (
	"encoding/base64"
	"fmt"
)

type Base64 []byte

func (dst *Base64) UnmarshalYAML(u func(interface{}) error) error {
	var s string
	if err := u(&s); err != nil {
		return err
	}

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		trimmed := s
		if len(trimmed) > 50 {
			trimmed = trimmed[:50]
		}

		return fmt.Errorf("Error while decoding base64 data beginning with %s: %s", trimmed, err.Error())
		return err
	}

	*dst = Base64(b)
	return nil
}
