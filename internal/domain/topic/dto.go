package topic

type (
	CreateReq struct {
		Name       string `json:"name" binding:"required,gte=4,lte=50"`
		LanguageId int    `json:"languageId" binding:"required,number"`
	}
)
