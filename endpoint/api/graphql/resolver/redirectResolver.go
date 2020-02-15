package resolver

import (
	"context"
	"fmt"
	"strconv"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	mw "clean_arch/endpoint/api/middleware"
	"clean_arch/infra/util"
	"clean_arch/registry"
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
func (r *redirectResolver) CreatedBy(ctx context.Context, obj *out.Redirect) (*out.Profile, error) {
	res := &out.Profile{
		ID:    obj.CreatedBy.ID,
		Name:  obj.CreatedBy.Name,
		Email: obj.CreatedBy.Email,
	}
	return res, nil
}

// mutationResolver
func (r *mutationResolver) CreateRedirect(ctx context.Context, input in.NewRedirect) (*out.Redirect, error) {
	_, claims, _ := mw.FromJWTContext(ctx)
	fmt.Println("createRedirect claims:", claims)
	if claims.ID != 0 {
		input.CID = util.Int642String(claims.ID)
	}
	fmt.Println("createRedirect input:", input)
	// for test hack
	if registry.Cfg.Mode == "test" && input.CID == "" {
		// for the first user id
		input.CID = "1"
	}
	fmt.Println("createRedirect input:", input)
	redirectID, err := r.RedirectUC.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	redirect, _ := r.RedirectUC.FindByID(ctx, &in.FetchRedirect{ID: string(redirectID)})
	return redirect, nil
}

// queryResolver -
func (r *queryResolver) Redirects(ctx context.Context) ([]*out.Redirect, error) {
	return r.RedirectUC.FindAll(ctx, &in.FetchAllOptions{})
}
func (r *queryResolver) FetchRedirectByCode(ctx context.Context, input in.FetchRedirectByCode) (*out.Redirect, error) {
	return r.RedirectUC.FindByCode(ctx, &input)
}
