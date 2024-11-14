package router

import (
	"github.com/gone-io/gone"
	"template_module/internal/interface/service"
	"template_module/internal/pkg/utils"
)

const IdAuthRouter = "router-auth"

type authRouter struct {
	gone.Flag
	gone.RouteGroup
	root  gone.RouteGroup    `gone:"*"`
	iUser service.IUserLogin `gone:"*"`
}

func (r *authRouter) GetGonerId() gone.GonerId {
	return IdAuthRouter
}

func (r *authRouter) AfterRevive() gone.AfterReviveError {
	r.RouteGroup = r.root.Group("/api", r.auth)
	return nil
}

func (r *authRouter) auth(ctx *gone.Context, in struct {
	authorization string `gone:"http,header"`
}) error {
	token, err := utils.GetBearerToken(in.authorization)
	if err != nil {
		return gone.ToError(err)
	}
	userId, err := r.iUser.GetUserIdFromToken(token)
	utils.SetUserId(ctx, userId)
	return err
}
