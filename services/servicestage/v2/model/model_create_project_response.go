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

// Response Object
type CreateProjectResponse struct {
	// 项目ID。
	Id *string `json:"id,omitempty"`
	// 项目名称。
	Name *string `json:"name,omitempty"`
	// 项目的clone url路径。
	CloneUrl *string `json:"clone_url,omitempty"`
}

func (o CreateProjectResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateProjectResponse", string(data)}, " ")
}
