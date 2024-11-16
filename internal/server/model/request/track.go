package request

type PostTrack struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
}

type PutTrack struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
}
