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
type UpdateAgencyRequestBody struct {
	Agency *UpdateAgencyOption `json:"agency"`
}

func (o UpdateAgencyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAgencyRequestBody", string(data)}, " ")
}
