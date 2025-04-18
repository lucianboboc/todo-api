package password

import "golang.org/x/crypto/bcrypt"

type Service interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) error
}

func NewService() Service {
	return service{}
}

type service struct {
}

func (s service) HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func (s service) ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
