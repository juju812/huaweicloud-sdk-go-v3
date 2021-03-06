/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateRequestThrottlingPolicyV2Response struct {
	// API与流控策略的绑定关系列表
	ThrottleApplys *[]ThrottleBindingResp `json:"throttle_applys,omitempty"`
}

func (o AssociateRequestThrottlingPolicyV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Response", string(data)}, " ")
}
