package logic

type OperationStatus int

const (
	SUCCESS OperationStatus = iota
	FAILED
)

type ServiceResponse struct {
	Status  OperationStatus
	Message string
	Body    []interface{}
}
