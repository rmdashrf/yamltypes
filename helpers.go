package yamltypes

import "time"

func unmarshalStringAndValidate(final *string, u func(interface{}) error, validation func(s string) error) error {
	var s string
	err := u(&s)
	if err != nil {
		return err
	}

	if err := validation(s); err != nil {
		return err
	} else {
		*final = s
		return nil
	}
}
func unmarshalTimeAndValidate(dst *time.Time, u func(interface{}) error,
	layout string) error {
	var s string
	err := u(&s)
	if err != nil {
		return err
	}

	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*dst = t
	return nil
}
