/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowMappingRequest struct {
	Id string `json:"id"`
}

func (o KeystoneShowMappingRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowMappingRequest", string(data)}, " ")
}
