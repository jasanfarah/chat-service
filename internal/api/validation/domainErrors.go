package validation

import "github.com/gofiber/fiber/v2"

type DomainError struct {
	Code       int    // Domain-specific error code
	Message    string // Human-readable message
	StatusCode int    // HTTP status code
}

func NewDomainError(code int, message string, httpStatus int) *DomainError {
	return &DomainError{
		Code:       code,
		Message:    message,
		StatusCode: httpStatus,
	}
}

// Error implements the error interface for DomainError.
func (e *DomainError) Error() string {
	return e.Message
}

// Definitions for custom domain errors
var ConversationNotFound = NewDomainError(0, "Conversation not found", fiber.StatusNotFound)
var MessageNotFound = NewDomainError(3, "Message not found", fiber.StatusNotFound)

var UUIDIsNotValid = NewDomainError(1, "UUID is not valid", fiber.StatusBadRequest)
var ValidationJSONError = NewDomainError(2, "Failed to parse request body", fiber.StatusUnprocessableEntity)

var MessageNotCreated = NewDomainError(4, "Failed to create message", fiber.StatusInternalServerError)
var ConversationNotCreated = NewDomainError(5, "Failed to create conversation", fiber.StatusInternalServerError)

