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

// 更新floatingip对象
type UpdateFloatingIpOption struct {
	// 端口id。
	PortId *string `json:"port_id,omitempty"`
}

func (o UpdateFloatingIpOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateFloatingIpOption", string(data)}, " ")
}
