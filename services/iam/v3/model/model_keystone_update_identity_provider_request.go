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
type KeystoneUpdateIdentityProviderRequest struct {
	Id   string                                     `json:"id"`
	Body *KeystoneUpdateIdentityProviderRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateIdentityProviderRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateIdentityProviderRequest", string(data)}, " ")
}