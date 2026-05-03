package iface

import (
	"context"
	
	"go-auth/internal/model/domain"
)

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}
