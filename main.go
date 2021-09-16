package gohelper

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
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
func ValidateStruct(validate *validator.Validate, data interface{}) error {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}
	var errName []string
	for _, e := range err.(validator.ValidationErrors) {
		errName = append(errName, e.Field())
	}
	var message = ""
	var sizeError = len(errName)
	if sizeError == 0 {
		message = ""
	} else if sizeError == 1 {
		message = errName[0]
	} else if sizeError == 2 {
		message = errName[0] + " & " + errName[1]
	} else {
		message = strings.Join(errName[0:sizeError-1], ", ") + " & " + errName[sizeError-1]
	}
	return fmt.Errorf("%s. tidak valid", message)
}
