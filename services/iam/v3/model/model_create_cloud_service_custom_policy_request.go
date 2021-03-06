/*
 * IAM
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
type CreateCloudServiceCustomPolicyRequest struct {
	Body *CreateCloudServiceCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateCloudServiceCustomPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateCloudServiceCustomPolicyRequest", string(data)}, " ")
}
