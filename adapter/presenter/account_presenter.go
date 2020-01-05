package presenter

import (
	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/usecase/out"
	"context"
	"strconv"
)

// NewAccountPresenter -
func NewAccountPresenter() presenter.AccountPresenter {
	return &accountPresenter{}
}

type accountPresenter struct {
}

// ViewError -
func (u accountPresenter) ViewError(ctx context.Context, err error) *out.Error {
	return &out.Error{
		Code:    "500",
		Message: err.Error(),
	}
}

// ViewAccountID -
func (u accountPresenter) ViewAccountID(ctx context.Context, account *model.UserAccount) out.ID {
	return out.ID(strconv.FormatInt(account.UID, 10))
}

// ViewAccount -
func (u accountPresenter) ViewAccount(ctx context.Context, account *model.UserAccount) *out.Account {
	return &out.Account{
		ID:        account.UID,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

// ViewAccounts -
func (u accountPresenter) ViewAccounts(ctx context.Context, accounts []*model.UserAccount) []*out.Account {
	res := make([]*out.Account, len(accounts))
	for i, account := range accounts {
		res[i] = &out.Account{
			ID:        account.UID,
			Name:      account.Name,
			Email:     account.Email,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}
	}
	return res
}
