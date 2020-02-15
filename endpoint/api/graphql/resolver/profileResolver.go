package resolver

import (
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	"context"
	"strconv"
)

// Profile -
func (r *Resolver) Profile() gen.ProfileResolver {
	return &profileResolver{r}
}

// profileResolver
type profileResolver struct{ *Resolver }

func (r *profileResolver) ID(ctx context.Context, obj *out.Profile) (string, error) {
	res := strconv.FormatInt(obj.ID, 10)
	return res, nil
}
