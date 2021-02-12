package utils

import (
	"context"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

var (
	IsDevMode = os.Getenv("MODE") == "dev"
)

func GetGRPCMSG(pbErr error) error {
	if pbErr == nil {
		return nil
	}
	s, _ := status.FromError(pbErr)
	return errors.New(s.Message())
}

func SetValGinCtx(ctx *gin.Context, name string, val interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), name, val)
	ctx.Request = ctx.Request.WithContext(newCtx)
}
