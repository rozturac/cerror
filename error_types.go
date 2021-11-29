package cerror

// ErrorType indicate to error level
type ErrorType string

const (
	// DomainError contains domain errors
	DomainError ErrorType = "DomainError"
	// BusinessError business domain errors
	BusinessError ErrorType = "BusinessError"
	// ApplicationError contains application errors
	ApplicationError ErrorType = "ApplicationError"
)
