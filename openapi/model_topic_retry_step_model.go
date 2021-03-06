/*
 * Algo.Run API 1.0-beta1
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0-beta1
 * Contact: support@algohub.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TopicRetryStepModel struct for TopicRetryStepModel
type TopicRetryStepModel struct {
	Index int32 `json:"index,omitempty"`
	BackoffDuration string `json:"backoffDuration,omitempty"`
	Repeat int32 `json:"repeat,omitempty"`
}
