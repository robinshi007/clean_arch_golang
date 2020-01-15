package resolver

import (
	"context"
	"strconv"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	"clean_arch/infra/util"
)

// Redirect -
func (r *Resolver) Redirect() gen.RedirectResolver {
	return &redirectResolver{r}
}

type redirectResolver struct{ *Resolver }

func (r *redirectResolver) ID(ctx context.Context, obj *out.Redirect) (string, error) {
	res := strconv.FormatInt(obj.ID, 10)
	return res, nil
}

func (r *redirectResolver) CreatedAt(ctx context.Context, obj *out.Redirect) (string, error) {
	res := obj.CreatedAt.Format(util.TimeFormatStr)
	return res, nil
}

// mutationResolver
func (r *mutationResolver) CreateRedirect(ctx context.Context, input in.NewRedirect) (*out.Redirect, error) {
	redirectID, err := r.RedirectUC.Save(ctx, &input)
	if err != nil {
		return nil, err
	}
	redirect, _ := r.RedirectUC.FindByID(ctx, &in.FetchRedirect{ID: string(redirectID)})
	return redirect, nil
}

// queryResolver -
func (r *queryResolver) Redirects(ctx context.Context) ([]*out.Redirect, error) {
	return r.RedirectUC.FindAll(ctx, &in.FetchRedirects{})
}
func (r *queryResolver) FetchRedirectByCode(ctx context.Context, input in.FetchRedirectByCode) (*out.Redirect, error) {
	return r.RedirectUC.FindByCode(ctx, &input)
}
