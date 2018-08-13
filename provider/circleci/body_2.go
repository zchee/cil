/*
 * CircleCI API
 *
 * The CircleCI API is a fully-featured JSON API that allows you to access all information and trigger all actions in CircleCI. 
 *
 * API version: v1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package circleci

type Body2 struct {

	BuildParameters *BuildParameters `json:"build_parameters,omitempty"`

	Parallel *Parallel `json:"parallel,omitempty"`

	Revision *Revision `json:"revision,omitempty"`
}
