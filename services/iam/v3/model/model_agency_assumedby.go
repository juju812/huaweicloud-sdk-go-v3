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
type AgencyAssumedby struct {
	User *AgencyAssumedbyUser `json:"user"`
}

func (o AgencyAssumedby) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AgencyAssumedby", string(data)}, " ")
}
