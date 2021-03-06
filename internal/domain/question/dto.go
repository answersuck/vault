package question

import "time"

type (
	CreateReq struct {
		Text       string `json:"text" binding:"required,gte=1,lte=200"`
		AnswerId   int    `json:"answerId" binding:"required,number"`
		MediaId    string `json:"mediaId" binding:"omitempty,uuid4"`
		LanguageId int    `json:"languageId" binding:"required,number"`
	}
)

type (
	CreateDTO struct {
		Text       string
		AnswerId   int
		MediaId    *string
		AccountId  string
		LanguageId int
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
