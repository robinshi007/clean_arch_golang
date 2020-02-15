package usecase

import (
	"context"
	"strconv"

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
func NewAccountUsecase(
	repo repository.AccountRepository,
	pre presenter.AccountPresenter,
) usecase.AccountUsecase {
	return &accountUsecase{
		repo: repo,
		pre:  pre,
	}
}

func (au *accountUsecase) FindAll(ctx context.Context, input *in.FetchAllOptions) ([]*out.Account, error) {
	accounts, err := au.repo.FindAll(ctx, &repository.ListOptions{
		Query: "",
		LimitOffset: &repository.LimitOffset{
			Limit:  50,
			Offset: 0,
		},
	})
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccounts(ctx, accounts), nil
}

func (au *accountUsecase) FindByID(ctx context.Context, input *in.FetchAccount) (*out.Account, error) {
	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := in.ToID(input.ID)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	account, err := au.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccount(ctx, account), nil
}

func (au *accountUsecase) FindByEmail(ctx context.Context, input *in.FetchAccountByEmail) (*out.Account, error) {
	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}

	account, err := au.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	return au.pre.ViewAccount(ctx, account), nil
}

func (au *accountUsecase) FindByName(ctx context.Context, input *in.FetchAccountByName) (*out.Account, error) {
	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}

	account, err := au.repo.FindByName(ctx, input.Name)
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
		Name:     input.Name,
		Password: PasswordHash,
	})
	if err != nil {
		return out.ID("-1"), err
	}
	id := strconv.FormatInt(res, 10)
	return out.ID(id), err
}

func (au *accountUsecase) Update(ctx context.Context, input *in.EditAccount) (*out.Account, error) {
	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	account, err := au.repo.FindByID(context.Background(), id)
	if err != nil {
		return nil, model.ErrEntityNotFound
	}

	// check is dirty
	if account.Name == input.Name {
		return nil, model.ErrEntityNotChanged
	}
	accountNew, err := au.repo.Update(ctx, &model.UserAccount{
		UID:  account.UID,
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	accountNew.Email = account.Email
	return au.pre.ViewAccount(ctx, accountNew), nil
}
func (au *accountUsecase) UpdatePassword(ctx context.Context, input *in.EditAccountPassword) (*out.Account, error) {
	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	account, err := au.repo.FindByID(context.Background(), id)
	if err != nil {
		return nil, model.ErrEntityNotFound
	}

	isMatch := util.ComparePassword(input.PasswordCurrent, account.Password)
	if isMatch != true {
		return nil, model.ErrPasswordIncorrect
	}

	passwordHash, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// check is dirty
	if account.Password == passwordHash {
		return nil, model.ErrEntityNotChanged
	}
	accountNew, err := au.repo.UpdatePassword(ctx, &model.UserAccount{
		UID:      account.UID,
		Password: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	accountNew.Name = account.Name
	accountNew.Email = account.Email
	return au.pre.ViewAccount(ctx, accountNew), nil
}

func (au *accountUsecase) Delete(ctx context.Context, input *in.FetchAccount) error {
	if err := in.Validate(input); err != nil {
		return model.ErrEntityBadInput
	}
	id, err := in.ToID(input.ID)
	if err != nil {
		return err
	}
	// ignore super admin user
	if id == 1 {
		return model.ErrEntityNotFound
	}
	return au.repo.Delete(ctx, id)
}

func (au *accountUsecase) Login(ctx context.Context, input *in.LoginAccountByEmail) (bool, int64, string, error) {
	if err := in.Validate(input); err != nil {
		return false, -1, "", model.ErrEntityBadInput
	}

	account, err := au.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return false, -1, "", err
	}
	result := util.ComparePassword(input.Password, account.Password)
	return result, account.UID, account.Name, nil
}
