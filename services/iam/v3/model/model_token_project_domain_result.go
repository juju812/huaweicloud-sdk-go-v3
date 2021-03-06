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

//
type TokenProjectDomainResult struct {
	// 账号名。
	Name string `json:"name"`
	// 账号ID。
	Id string `json:"id"`
}

func (o TokenProjectDomainResult) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TokenProjectDomainResult", string(data)}, " ")
}
