package repository

import (
	"context"
)

// CasbinRuleRepository -
type CasbinRuleRepository interface {
	Count(ctx context.Context) (int64, error)
}
