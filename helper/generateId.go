package helper

import (
	"github.com/matoous/go-nanoid/v2"
)

func GenerateId(prefix string) (string, error) {
	id, err := gonanoid.New(10)

	return prefix + "_" + id, err
}
