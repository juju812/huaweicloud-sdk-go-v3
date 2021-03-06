/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSubNetworkInterfaceRequest struct {
	Body *CreateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o CreateSubNetworkInterfaceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSubNetworkInterfaceRequest", string(data)}, " ")
}
