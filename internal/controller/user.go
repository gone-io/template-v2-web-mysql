package controller

import (
	"template_module/internal/interface/entity"
	"template_module/internal/interface/service"
	"template_module/internal/pkg/utils"

	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
)

type userCtr struct {
	gone.Flag
	a gone.RouteGroup `gone:"router-auth"`
	p gone.RouteGroup `gone:"router-pub"`

	iUser       service.IUser      `gone:"*"`
	iUserLogin  service.IUserLogin `gone:"*"`
	gone.Logger `gone:"*"`
}

func (c *userCtr) Mount() gone.GinMountError {
	c.Infof("mount user controller")
	c.p.
		POST("/users/login", func(in struct {
			req *entity.LoginParam `gone:"http,body"`
		}) (*entity.LoginResult, error) {
			return c.iUserLogin.Login(in.req)
		}).
		POST("/users/register", func(in struct {
			req *entity.RegisterParam `gone:"http,body"`
		}) (*entity.LoginResult, error) {
			return c.iUserLogin.Register(in.req)
		})

	c.a.
		GET("/users/me", func(ctx *gin.Context) (any, error) {
			userId := utils.GetUserId(ctx)
			return c.iUser.GetUserById(userId)
		}).
		POST("/users/logout", func(in struct {
			authorization string `gone:"http,header"`
		}) error {
			token, err := utils.GetBearerToken(in.authorization)
			if err != nil {
				return gone.ToError(err)
			}
			return c.iUserLogin.Logout(token)
		})
	return nil
}
