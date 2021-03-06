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
type AllowAddressNetmasksResult struct {
	// IP地址或网段，例如：192.168.0.1/24。
	AddressNetmask string `json:"address_netmask"`
	// 描述信息。
	Description string `json:"description"`
}

func (o AllowAddressNetmasksResult) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AllowAddressNetmasksResult", string(data)}, " ")
}
