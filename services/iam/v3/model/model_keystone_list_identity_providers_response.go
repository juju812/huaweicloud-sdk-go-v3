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
type KeystoneListIdentityProvidersResponse struct {
	// 身份提供商信息列表。
	IdentityProviders *[]IdentityprovidersResult `json:"identity_providers,omitempty"`
	Links             *Links                     `json:"links,omitempty"`
}

func (o KeystoneListIdentityProvidersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListIdentityProvidersResponse", string(data)}, " ")
}
