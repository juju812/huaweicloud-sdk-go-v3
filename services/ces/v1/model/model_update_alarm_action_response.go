/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAlarmActionResponse struct {
}

func (o UpdateAlarmActionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAlarmActionResponse", string(data)}, " ")
}