package userhttp

import (
	"time"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
)

type createUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func newCreateUserResponse(output *usecase.CreateUserOutput) createUserResponse {
	return createUserResponse{
		ID:        output.ID,
		Email:     output.Email,
		Name:      output.Name,
		CreatedAt: formatTime(output.CreatedAt),
		UpdatedAt: formatTime(output.UpdatedAt),
	}
}

func formatTime(value time.Time) string {
	return value.Format(time.RFC3339)
}
