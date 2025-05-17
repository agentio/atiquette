package lexutil

import (
	"encoding/json"
)

type typeExtractor struct {
	Type string `json:"$type" cborgen:"$type"`
}

func TypeExtract(b []byte) (string, error) {
	var te typeExtractor
	if err := json.Unmarshal(b, &te); err != nil {
		return "", err
	}

	return te.Type, nil
}
