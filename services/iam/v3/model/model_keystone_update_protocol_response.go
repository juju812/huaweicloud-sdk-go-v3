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

// Response Object
type KeystoneUpdateProtocolResponse struct {
	Protocol *ProtocolResult `json:"protocol,omitempty"`
}

func (o KeystoneUpdateProtocolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateProtocolResponse", string(data)}, " ")
}
