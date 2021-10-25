package gojson

import (
	"errors"
	"github.com/goccy/go-json"
	"log"
)

type GoJson struct {
}


func (g *GoJson) EncodeJson(v interface{}) (*[]byte, error) {
	value, err := json.Marshal(&v)
	if err != nil {
		log.Printf("Can not marshal Value")
		return nil, errors.New("Can not marshal Value")
	}
	return &value, nil
}

func (g *GoJson) DecodeJson(message *[]byte, v interface{}) error {
	err := json.Unmarshal(*message, &v)
	if err != nil {
		log.Printf("Can not unmarshal JSON")
		return errors.New("Can not unmarshal JSON")
	}
	return nil
}