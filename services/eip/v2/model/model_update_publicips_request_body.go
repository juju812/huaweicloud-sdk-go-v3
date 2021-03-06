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

// 更新弹性公网IP的请求体
type UpdatePublicipsRequestBody struct {
	Publicip *UpdatePublicipOption `json:"publicip"`
}

func (o UpdatePublicipsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipsRequestBody", string(data)}, " ")
}
