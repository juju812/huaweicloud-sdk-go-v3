/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateSharedBandwidthsResponse struct {
	// 批创的带宽对象的列表
	Bandwidths *[]BatchBandwidthResp `json:"bandwidths,omitempty"`
}

func (o BatchCreateSharedBandwidthsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateSharedBandwidthsResponse", string(data)}, " ")
}
