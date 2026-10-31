package service

import "errors"

type User struct {
	ID    int64
	Email string
}

var ErrNotFound = errors.New("not found")

type UserRepo interface{ ByEmail(string) (User, error) }
type Service struct{ repo UserRepo }

func New(r UserRepo) *Service { return &Service{r} }
func (s *Service) FindIDByEmail(email string) (int64, error) {
	u, e := s.repo.ByEmail(email)
	if e != nil {
		return 0, e
	}
	return u.ID, nil
}
