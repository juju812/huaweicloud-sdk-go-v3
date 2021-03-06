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
type BatchDeletePublicipTagsRequest struct {
	PublicipId string                              `json:"publicip_id"`
	Body       *BatchDeletePublicipTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePublicipTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeletePublicipTagsRequest", string(data)}, " ")
}
