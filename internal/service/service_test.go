package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type stubRepo struct {
	users map[string]User
	err   error
	calls []string
}

func (r *stubRepo) ByEmail(email string) (User, error) {
	r.calls = append(r.calls, email)
	if r.err != nil {
		return User{}, r.err
	}
	u, ok := r.users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
func TestFindIDByEmail(t *testing.T) {
	r := &stubRepo{users: map[string]User{"test@example.com": {7, "test@example.com"}}}
	s := New(r)
	t.Run("found", func(t *testing.T) {
		id, e := s.FindIDByEmail("test@example.com")
		require.NoError(t, e)
		assert.Equal(t, int64(7), id)
	})
	t.Run("missing", func(t *testing.T) {
		id, e := s.FindIDByEmail("missing@example.com")
		assert.ErrorIs(t, e, ErrNotFound)
		assert.Zero(t, id)
	})
	t.Run("repository error", func(t *testing.T) {
		failure := errors.New("database unavailable")
		r.err = failure
		_, e := s.FindIDByEmail("test@example.com")
		assert.ErrorIs(t, e, failure)
	})
	assert.Equal(t, []string{"test@example.com", "missing@example.com", "test@example.com"}, r.calls)
}
