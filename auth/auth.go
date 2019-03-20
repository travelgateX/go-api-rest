package auth

import (
	"context"
	"go-api-rest/utils"

	authorization "github.com/travelgateX/go-jwt-tools"
)

// Auth entities auth object
type Auth struct {
	authorization.Permissions
}

// NewAuth creates new authentication object
func NewAuth(ctx context.Context) Auth {
	user, _ := authorization.UserFromContext(ctx)
	return Auth{user.Permissions}
}

// GetPermissions obtains authorization permission object
func (a *Auth) GetPermissions() authorization.Permissions {
	return a.Permissions
}

// CheckAllPermission return cartesian product of all permissions
func (a *Auth) CheckAllPermission(product string, object string, per authorization.Permission, groups ...string) []string {
	var gr []string

	g1, _ := a.CheckPermission(product, object, per, groups...)
	g2, _ := a.CheckPermission(product, "all", per, groups...)
	g4, _ := a.CheckPermission("all", object, per, groups...)
	g3, _ := a.CheckPermission("all", "all", per, groups...)

	a1, _ := a.CheckPermission(product, object, "a", groups...)
	a2, _ := a.CheckPermission(product, "all", "a", groups...)
	a4, _ := a.CheckPermission("all", object, "a", groups...)
	a3, _ := a.CheckPermission("all", "all", "a", groups...)

	gr = utils.AppendUniqueSlices(gr, g1)
	gr = utils.AppendUniqueSlices(gr, g2)
	gr = utils.AppendUniqueSlices(gr, g3)
	gr = utils.AppendUniqueSlices(gr, g4)

	gr = utils.AppendUniqueSlices(gr, a1)
	gr = utils.AppendUniqueSlices(gr, a2)
	gr = utils.AppendUniqueSlices(gr, a3)
	gr = utils.AppendUniqueSlices(gr, a4)

	return gr
}
