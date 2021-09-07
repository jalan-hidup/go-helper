package gohelper

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	UserRolesUmat   string = "UMAT"
	UserRolesPastor string = "PASTOR"
	UserRolesAdmin  string = "ADMIN"
)

func contains(s *[]string, str *string) bool {
	for _, v := range *s {
		if v == *str {
			return true
		}
	}
	return false
}
func GenerateId(prefix string) string {
	now := time.Now().UnixMilli()
	return prefix + "-" + fmt.Sprint(rand.Intn(999-100+1)+100) + fmt.Sprint(now) + fmt.Sprint(rand.Intn(999-100+1)+100)
}
func ValidateUserRoles(ctx *context.Context, roles ...string) (*requestUser, error) {
	var user requestUser
	if len(roles) == 0 {
		return &user, status.Error(codes.InvalidArgument, "roles required")
	}
	md, ok := metadata.FromIncomingContext(*ctx)
	if !ok {
		return &user, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	listAuth := md.Get("x-apigateway-api-userinfo")
	if len(listAuth) == 0 {
		return &user, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	byteResult, err := base64.StdEncoding.DecodeString(listAuth[0])
	if err != nil {
		return &user, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	errValidateJsonUser := json.Unmarshal(byteResult, &user)
	if errValidateJsonUser != nil {
		return &user, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	permissionAccept := contains(&roles, &user.Roles)
	if !permissionAccept {
		return &user, status.Error(codes.PermissionDenied, "permission denied")
	}
	return &user, nil
}
