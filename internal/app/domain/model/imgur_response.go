package model

// ImgurResponse :
type ImgurResponse struct {
	Data    imgurResponseData `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

// imgurResponseData :
type imgurResponseData struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}
