package request

type PostDriver struct {
	Name     string `json:"name" binding:"required"`
	RaceName string `json:"race_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type PutDriver struct {
	Name     string `json:"name" binding:"required"`
	RaceName string `json:"race_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
