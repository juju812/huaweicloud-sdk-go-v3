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
type AssociatePublicipsResponse struct {
	// 本次请求的编号
	RequestId *string           `json:"request_id,omitempty"`
	Publicip  *PublicipShowResp `json:"publicip,omitempty"`
}

func (o AssociatePublicipsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociatePublicipsResponse", string(data)}, " ")
}
