package user

type Contract interface {
	Create(email, name, password string) (int64, error)
	GenerateToken(email, password string) (string, error)
}
