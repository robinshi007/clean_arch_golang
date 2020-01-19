package resolver

import (
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	"context"
	"strconv"
)

// Account -
func (r *Resolver) Account() gen.AccountResolver {
	return &accountResolver{r}
}

// accountResolver
type accountResolver struct{ *Resolver }

func (r *accountResolver) ID(ctx context.Context, obj *out.Account) (string, error) {
	res := strconv.FormatInt(obj.ID, 10)
	return res, nil
}

// mutationResolver
func (r *mutationResolver) CreateAccount(ctx context.Context, input in.NewAccount) (*out.Account, error) {
	accountID, err := r.AccountUC.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	account, _ := r.AccountUC.FindByID(ctx, &in.FetchAccount{ID: string(accountID)})
	return account, nil
}
func (r *mutationResolver) UpdateAccount(ctx context.Context, input in.EditAccount) (*out.Account, error) {
	return r.AccountUC.Update(ctx, &input)
}
func (r *mutationResolver) UpdateAccountPassword(ctx context.Context, input in.EditAccountPassword) (*out.Account, error) {
	return r.AccountUC.UpdatePassword(ctx, &input)
}
func (r *mutationResolver) DeleteAccount(ctx context.Context, input in.FetchAccount) (*out.Account, error) {
	account, err := r.AccountUC.FindByID(ctx, &input)
	if err != nil {
		return nil, err
	}
	err = r.AccountUC.Delete(ctx, &input)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// queryResolver
func (r *queryResolver) Accounts(ctx context.Context) ([]*out.Account, error) {
	return r.AccountUC.FindAll(ctx, &in.FetchAllOptions{})
}
func (r *queryResolver) FetchAccount(ctx context.Context, input in.FetchAccount) (*out.Account, error) {
	return r.AccountUC.FindByID(ctx, &input)
}
