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
type RemovePublicipsFromSharedBandwidthResponse struct {
}

func (o RemovePublicipsFromSharedBandwidthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemovePublicipsFromSharedBandwidthResponse", string(data)}, " ")
}
