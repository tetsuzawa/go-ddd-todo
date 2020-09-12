package playground

import "errors"

const notUsed = "not used"

func a() error {
	errors.New("playgrond")

	err := errors.New("error")
	if err != nil {
		return err
	}
	return nil
}
