package core

type ResponseMetadata struct {
	RequestId string
}
type ResponseError struct {
	Type    string
	Code    string
	Message string
}
