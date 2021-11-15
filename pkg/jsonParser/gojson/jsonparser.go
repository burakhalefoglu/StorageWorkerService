package gojson

import (
	"errors"
	"github.com/goccy/go-json"
)

type goJson struct {
}

func GoJsonConstructor() *goJson {
	return &goJson{}
}

func (g *goJson) EncodeJson(v interface{}) (*[]byte, error) {
	value, err := json.Marshal(&v)
	if err != nil {
		return nil, errors.New("Can not marshal Value")
	}
	return &value, nil
}

func (g *goJson) DecodeJson(message *[]byte, v interface{}) error {
	err := json.Unmarshal(*message, &v)
	if err != nil {
		return errors.New("Can not unmarshal JSON")
	}
	return nil
}