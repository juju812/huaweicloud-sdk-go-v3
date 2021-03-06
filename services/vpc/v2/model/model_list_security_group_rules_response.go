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

// Response Object
type ListSecurityGroupRulesResponse struct {
	// 安全组规则对象列表
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty"`
}

func (o ListSecurityGroupRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSecurityGroupRulesResponse", string(data)}, " ")
}
