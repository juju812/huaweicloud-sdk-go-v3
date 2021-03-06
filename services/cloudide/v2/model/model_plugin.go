/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type Plugin struct {
	// 插件属性
	Attribute *string `json:"attribute,omitempty"`
	// 插件名
	Name *string `json:"name,omitempty"`
}

func (o Plugin) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Plugin", string(data)}, " ")
}
