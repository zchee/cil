/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI. 
 *
 * API version: v1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package circleci

type Artifact struct {

	NodeIndex int32 `json:"node_index,omitempty"`

	Path string `json:"path,omitempty"`

	PrettyPath string `json:"pretty_path,omitempty"`

	Url string `json:"url,omitempty"`
}
