package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"jasanfarah/chat-service/internal/api/dto"
	"jasanfarah/chat-service/internal/api/validation"
	"time"
)

// ErrorHandler checks for domain-specific errors and sets the appropriate HTTP status code and response body
func ErrorHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil {
			// Check if the error is a DomainError
			var domainErr *validation.DomainError
			if errors.As(err, &domainErr) {
				// Prepare an error response DTO based on the properties of a domain error
				errorResponse := dto.DomainErrorResponse{
					Message:         domainErr.Message,
					Timestamp:       time.Now(),
					DomainErrorCode: domainErr.Code,
				}

				return ctx.Status(domainErr.StatusCode).JSON(errorResponse)
			}

			var e *fiber.Error
			if errors.As(err, &e) {
				// You can optionally log the error here
				// Return the Fiber error directly to use its status code and message
				return ctx.Status(e.Code).JSON(fiber.Map{
					"error":     e.Message,
					"timestamp": time.Now().Format(time.RFC3339),
				})
			}

			// For non-domain errors, internal server error response
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":     "Internal Server Error",
				"message":   "Something went wrong. Please try again later.",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		}
		return nil
	}
}