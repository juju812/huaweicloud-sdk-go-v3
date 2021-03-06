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

// Request Object
type NeutronUpdateFloatingIpRequest struct {
	FloatingipId string                              `json:"floatingip_id"`
	Body         *NeutronUpdateFloatingIpRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFloatingIpRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFloatingIpRequest", string(data)}, " ")
}
