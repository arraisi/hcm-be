package memory

import (
	"errors"
	"sync"
	"time"

	"hcm-be/internal/domain"
	"github.com/google/uuid"
)

type UserRepo struct {
	mu    sync.RWMutex
	items map[string]domain.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{items: make(map[string]domain.User)}
}

func (r *UserRepo) FindAll() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.User, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return out, nil
}

func (r *UserRepo) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &v, nil
}

func (r *UserRepo) Create(u domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	r.items[u.ID] = u
	return nil
}
