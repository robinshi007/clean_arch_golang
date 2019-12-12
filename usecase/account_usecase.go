package usecase

import (
	"context"
	"strconv"
	"time"

	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/repository"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/infra/util"
)

type accountUsecase struct {
	repo repository.AccountRepository
	pre  presenter.AccountPresenter
}

// NewAccountUseCase -
func NewAccountUseCase(
	repo repository.AccountRepository,
	pre presenter.AccountPresenter,
	timeout time.Duration,
) usecase.AccountUsecase {
	return &accountUsecase{
		repo: repo,
		pre:  pre,
	}
}

func (au *accountUsecase) GetAll(ctx context.Context, num int64) ([]*out.Account, error) {
	accounts, err := au.repo.GetAll(ctx, &repository.AccountListOptions{})
	time.Sleep(3 * time.Second)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccounts(ctx, accounts), nil
}

func (au *accountUsecase) GetByID(ctx context.Context, input *in.FetchAccount) (*out.Account, error) {

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := in.ToID(input.ID)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	account, err := au.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccount(ctx, account), nil
}

func (au *accountUsecase) GetByEmail(ctx context.Context, input *in.FetchAccountByEmail) (*out.Account, error) {

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}

	account, err := au.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccount(ctx, account), nil
}

func (au *accountUsecase) GetByName(ctx context.Context, input *in.FetchAccountByName) (*out.Account, error) {

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}

	account, err := au.repo.GetByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccount(ctx, account), nil
}

func (au *accountUsecase) Create(ctx context.Context, input *in.NewAccount) (out.ID, error) {
	if err := in.Validate(input); err != nil {
		return out.ID("-1"), model.ErrEntityBadInput
	}

	PasswordHash, err := util.HashPassword(input.Password)
	if err != nil {
		return out.ID("-1"), err
	}

	res, err := au.repo.Create(ctx, &model.UserAccount{
		Email:    input.Email,
		Password: PasswordHash,
	})
	if err != nil {
		return out.ID("-1"), err
	}
	id := strconv.FormatInt(res, 10)
	return out.ID(id), err
}

func (au *accountUsecase) UpdatePassword(ctx context.Context, input *in.EditAccount) (*out.Account, error) {

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	account, err := au.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, model.ErrEntityNotFound
	}

	// check is dirty
	if account.Password == input.Password {
		return nil, model.ErrEntityNotChanged
	}
	accountNew, err := au.repo.Update(ctx, &model.UserAccount{
		UID:      account.UID,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	accountNew.Email = account.Email
	return au.pre.ViewAccount(ctx, accountNew), nil
}

func (au *accountUsecase) Delete(ctx context.Context, input *in.FetchAccount) error {
	err := in.Validate(input)
	if err != nil {
		return model.ErrEntityBadInput
	}

	id, err := in.ToID(input.ID)
	if err != nil {
		return err
	}
	return au.repo.Delete(ctx, id)
}

func (au *accountUsecase) Login(ctx context.Context, input *in.LoginAccountByEmail) (bool, error) {

	if err := in.Validate(input); err != nil {
		return false, model.ErrEntityBadInput
	}

	account, err := au.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return false, err
	}
	result := util.ComparePassword(input.Password, account.Password)
	return result, nil
}
