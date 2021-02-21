package utils

import (
	"context"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func IsPong(response *redis.StatusCmd) bool {
	if response.Err() != nil {
		return false
	}
	return response.String() == "ping: PONG"
}

func UniqueIds(n []uint64) []uint64 {
	var (
		result []uint64
		m      = make(map[uint64]interface{})
	)

	{
		for _, i := range n {
			m[i] = struct{}{}
		}
	}
	{
		result = make([]uint64, 0, len(m))
		for val := range m {
			result = append(result, val)
		}
	}
	return result
}
