package main

import (
	"github.com/tidwall/gjson"
	"strings"
)

// getStringIfExists returns json value or default
func getStringIfExists(json gjson.Result, key string, def string) string {
	value := json.Get(key)
	if value.Exists() {
		return value.String()
	}
	return def
}

// getFloatIfExists returns json value or default
func getFloatIfExists(json gjson.Result, key string, def float64) float64 {
	value := json.Get(key)
	if value.Exists() {
		return value.Float()
	}
	return def
}

// getLongFlags convert flags map to string
func getLongFlags(json gjson.Result, flags []string) string {
	var result []string
	for _, flag := range flags {
		jFlag := json.Get(flag)
		if jFlag.Exists() && jFlag.Bool() {
			result = append(result, flag)
		}
	}
	return strings.Join(result, ",")
}
