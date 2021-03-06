/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateProjectV4Request struct {
	ProjectId string                  `json:"project_id"`
	Body      *UpdateProjectRequestV4 `json:"body,omitempty"`
}

func (o UpdateProjectV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateProjectV4Request", string(data)}, " ")
}
