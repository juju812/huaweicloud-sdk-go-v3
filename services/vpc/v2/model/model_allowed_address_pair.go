/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type AllowedAddressPair struct {
	// 功能说明：IP地址 取值范围：可以是IP地址或CIDR 约束：不支持0.0.0.0/0如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。
	IpAddress *string `json:"ip_address,omitempty"`
	// mac地址
	MacAddress *string `json:"mac_address,omitempty"`
}

func (o AllowedAddressPair) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AllowedAddressPair", string(data)}, " ")
}
