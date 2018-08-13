/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI. 
 *
 * API version: v1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package circleci

import (
	"time"
)

type BuildWorkflow struct {

	Status *Status `json:"status,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Id string `json:"id,omitempty"`
}
