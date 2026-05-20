package proxy

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/h-mall/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetUserByID(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.String(consts.StatusBadRequest, "invalid user id")
		return
	}

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.String(consts.StatusInternalServerError, "connect user service failed")
		return
	}

	defer conn.Close()

	client := api.NewUserServiceClient(conn)

	resp, err := client.GetUser(ctx, &api.GetUserRequest{
		Id: id,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, fmt.Sprintf("call user service failed: %v", err))
		return
	}

	c.JSON(consts.StatusOK, resp)
}