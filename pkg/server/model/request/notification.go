package request

type PostNotification struct {
	Message       string `json:"message"`
	FirstDriver   string `json:"first_driver"`
	SecondDriver  string `json:"second_driver"`
	ThirdDriver   string `json:"third_driver"`
	LicensePoints int    `json:"license_points"`
}

type PutNotification struct {
	Message       string `json:"message"`
	FirstDriver   string `json:"first_driver"`
	SecondDriver  string `json:"second_driver"`
	ThirdDriver   string `json:"third_driver"`
	LicensePoints int    `json:"license_points"`
}
