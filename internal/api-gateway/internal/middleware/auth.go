package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			ctx.AbortWithStatusJSON(consts.StatusUnauthorized, utils.H{
				"error": "Unauthorized",
			})
			return
		}

		ctx.Next(c)
	}
}