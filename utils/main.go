package utils

import (
	"errors"
	"os"

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
