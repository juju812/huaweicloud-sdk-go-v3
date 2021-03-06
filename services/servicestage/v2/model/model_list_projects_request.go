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
type ListProjectsRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
	Namespace string `json:"namespace"`
}

func (o ListProjectsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}
