package userdefined

import (
	"lisp_lambda-go/internal/userdefined/internal"
	"strconv"
)

type invalidNameDeterminer struct {
	keywords map[string]bool
}

func NewInvalidNameDeterminer(
	keywords map[string]bool,
) *invalidNameDeterminer {
	return &invalidNameDeterminer{keywords}
}

func (i invalidNameDeterminer) IsInvalidName(value string) bool {
	_, isKeyword := i.keywords[value]

	_, err := strconv.Atoi(value)
	isNumeric := err == nil

	return isKeyword || isNumeric
}

var _ internal.InvalidNameDeterminer = &invalidNameDeterminer{}