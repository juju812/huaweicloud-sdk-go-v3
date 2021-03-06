/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQueryHttpCodeRequest struct {
	PlayDomains []string  `json:"play_domains"`
	Code        *[]string `json:"code,omitempty"`
	Region      *[]string `json:"region,omitempty"`
	Isp         *[]string `json:"isp,omitempty"`
	StartTime   *string   `json:"start_time,omitempty"`
	EndTime     *string   `json:"end_time,omitempty"`
}

func (o ListQueryHttpCodeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQueryHttpCodeRequest", string(data)}, " ")
}
