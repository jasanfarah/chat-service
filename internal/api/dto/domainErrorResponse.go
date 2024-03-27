package dto

import (
    "jasanfarah/chat-service/internal/api/validation"
	"time"
)

type DomainErrorResponse struct {
	Message         string    `json:"error"`
	Timestamp       time.Time `json:"timestamp"`
	DomainErrorCode int       `json:"domainErrorCode"`
}

func NewDomainErrorResponse(error *validation.DomainError) *DomainErrorResponse {
	return &DomainErrorResponse{
		Message:         error.Message,
		Timestamp:       time.Now(),
		DomainErrorCode: error.Code,
	}
}
