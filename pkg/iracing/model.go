package iracing

type AuthResponse struct {
	AuthCode string `json:"authcode"`
}

type IRatingData struct {
	SportsCar  int `json:"sports_car"`
	FormulaCar int `json:"formula_car"`
	Oval       int `json:"oval"`
}

type Member struct {
	CustID   int            `json:"cust_id"`
	Licenses []LicenseEntry `json:"licenses"`
}

type LicenseEntry struct {
	Category string `json:"category"`
	IRating  int    `json:"irating"`
}

type MemberDataLink struct {
	Link string `json:"link"`
}

type MemberData struct {
	Members []Member `json:"members"`
}
