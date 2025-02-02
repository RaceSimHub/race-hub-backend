//go:generate mockgen -destination=./mock/user_mock.go -package=mock . Contract
package user

type Contract interface {
	Create(email, name, password string) (int64, error)
	GenerateToken(email, password string) (string, error)
}
