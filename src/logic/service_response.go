package logic

type OperationStatus int

const (
	SUCCESS OperationStatus = iota
	FAILED
)

type ServiceResponse struct {
	status  OperationStatus
	message string
	body    []interface{}
}
