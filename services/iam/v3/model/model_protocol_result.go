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
type ProtocolResult struct {
	// 协议ID。
	Id string `json:"id"`
	// 映射ID。
	MappingId string         `json:"mapping_id"`
	Links     *ProtocolLinks `json:"links"`
}

func (o ProtocolResult) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ProtocolResult", string(data)}, " ")
}
