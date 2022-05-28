package account

import (
	"time"
)

type (
	CreateRequest struct {
		Email    string `json:"email" binding:"required,email,lte=255"`
		Nickname string `json:"nickname" binding:"required,alphanum,gte=4,lte=25"`
		Password string `json:"password" binding:"required,gte=8,lte=71"`
	}

	ResetPasswordRequest struct {
		Login string `json:"login" binding:"required,email|alphanum"`
	}

	SetPasswordRequest struct {
		Password string `json:"password" binding:"required,gte=8,lte=64"`
	}
)

type (
	VerificationDTO struct {
		Email    string
		Code     string
		Verified bool
	}

	SetPasswordDTO struct {
		AccountId    string
		Token        string
		PasswordHash string
		UpdatedAt    time.Time
	}

	sendEmailDTO struct {
		To         string
		Subject    string
		Template   string
		Format     string
		FormatArgs []any
	}
)
