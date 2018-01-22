package pocket

import "errors"

var (
	ErrInvalidArgs = errors.New("invalid args")
	ErrInvalidFunction = errors.New("invalid function")
	ErrEmptyAddr = errors.New("the addr is empty")
	ErrNoAccount = errors.New("no account found")
	ErrKeyNoData = errors.New("state key found, but no data")
	ErrInvalidAddr = errors.New("the addr is invalid(contains '_')")
)
