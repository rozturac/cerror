package cerror

type ErrorType string

const (
	DomainError      ErrorType = "DomainError"
	BusinessError    ErrorType = "BusinessError"
	ApplicationError ErrorType = "ApplicationError"
)
