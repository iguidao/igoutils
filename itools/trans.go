package itools

import (
	"encoding/json"
)

func JsonToMap(jsonstr string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonstr), &m)
	if err != nil {
		return nil
	}
	return m
}
