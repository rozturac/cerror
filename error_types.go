package cerror

// ErrorType indicate to error level
type ErrorType string

const (
	DomainError      ErrorType = "DomainError"      // DomainError contains domain errors
	BusinessError    ErrorType = "BusinessError"    // DomainError business domain errors
	ApplicationError ErrorType = "ApplicationError" // DomainError contains application errors
)
