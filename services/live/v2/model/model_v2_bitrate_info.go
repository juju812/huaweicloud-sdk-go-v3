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

// 实时码率
type V2BitrateInfo struct {
	// 域名。
	PublishDomain *string `json:"publish_domain,omitempty"`
	// 应用名称。
	App *string `json:"app,omitempty"`
	// 流名。
	Stream *string `json:"stream,omitempty"`
	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`
	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty"`
	// 数据
	DataList *[]int64 `json:"data_list,omitempty"`
}

func (o V2BitrateInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"V2BitrateInfo", string(data)}, " ")
}
