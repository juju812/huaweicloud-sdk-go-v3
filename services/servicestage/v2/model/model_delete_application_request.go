/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationRequest struct {
	ApplicationId string `json:"application_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
