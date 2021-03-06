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

type ApplicationConfigModify struct {
	// 环境ID。
	EnvironmentId string                                `json:"environment_id"`
	Configuration *ApplicationConfigModifyConfiguration `json:"configuration"`
}

func (o ApplicationConfigModify) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ApplicationConfigModify", string(data)}, " ")
}
