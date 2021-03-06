/*
 * ELB
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
type UpdateWhitelistResponse struct {
	Whitelist *WhitelistResp `json:"whitelist,omitempty"`
}

func (o UpdateWhitelistResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateWhitelistResponse", string(data)}, " ")
}
