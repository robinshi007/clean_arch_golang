package postgres

import (
	"context"

	"clean_arch/domain/repository"
	"clean_arch/registry"
)

// NewCasbinRuleRepo -
func NewCasbinRuleRepo() repository.CasbinRuleRepository {
	return &casbinRuleRepo{}
}

type casbinRuleRepo struct {
}

func (u *casbinRuleRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	row := registry.Db.QueryRowContext(ctx, "SELECT Count(*) from casbin_rules")
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
