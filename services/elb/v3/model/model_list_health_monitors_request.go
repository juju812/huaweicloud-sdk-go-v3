/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListHealthMonitorsRequest struct {
	AdminStateUp        *bool     `json:"admin_state_up,omitempty"`
	Delay               *[]int32  `json:"delay,omitempty"`
	DomainName          *[]string `json:"domain_name,omitempty"`
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	ExpectedCodes       *[]string `json:"expected_codes,omitempty"`
	HttpMethod          *[]string `json:"http_method,omitempty"`
	Id                  *[]string `json:"id,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	Marker              *string   `json:"marker,omitempty"`
	MaxRetries          *[]int32  `json:"max_retries,omitempty"`
	MaxRetriesDown      *[]int32  `json:"max_retries_down,omitempty"`
	MonitorPort         *[]int32  `json:"monitor_port,omitempty"`
	Name                *[]string `json:"name,omitempty"`
	PageReverse         *bool     `json:"page_reverse,omitempty"`
	Timeout             *int32    `json:"timeout,omitempty"`
	Type                *[]string `json:"type,omitempty"`
	UrlPath             *[]string `json:"url_path,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
