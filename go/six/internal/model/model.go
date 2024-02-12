package model

type PostRequest struct {
	Name string `json:"name"`
}

type StoredData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
