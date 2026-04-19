package iface

import (
	"go-auth/internal/model/domain"
)

type UserRepositoryInterface interface {
	FindByEmail(email string) (domain.User, error)
}
