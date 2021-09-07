package gohelper_test

import (
	"context"
	"testing"

	gohelper "github.com/jalan-hidup/go-helper"
	"google.golang.org/grpc/metadata"
)

func BenchmarkGenerateId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gohelper.GenerateId("ID")
	}
}
func BenchmarkValidateUserRoles(b *testing.B) {
	header := metadata.New(map[string]string{"X-Apigateway-Api-Userinfo": "eyJpZCI6IjEyMyIsImNpdHkiOiJNQUtBU1NBUiIsImVtYWlsIjoibXVkaUBnbWFpbCIsInBob25lTnVtYmVyIjoiKzYyODEzNTQ3NTY5MTIiLCJyb2xlcyI6IkFETUlOIiwiYXV0aFRpbWUiOjE2MzEwNTI3OTA0ODksImV4cCI6MTYzMTA1Mjc5MDQ4OSwiaWF0IjoxNjMxMDUyNzkwNDg5LCJzdWIiOiJqYWxhbi1oaWR1cC0xMjMiLCJhdWgiOiIxMjMiLCJpc3MiOiJqYWxhbi1oaWR1cC1pc3MifQ=="})
	ctx := metadata.NewIncomingContext(context.Background(), header)
	for i := 0; i < b.N; i++ {
		gohelper.ValidateUserRoles(&ctx, "ADMIN")
	}
}
func TestValidateUserRoles(t *testing.T) {
	header := metadata.New(map[string]string{"X-Apigateway-Api-Userinfo": "eyJpZCI6IjEyMyIsImNpdHkiOiJNQUtBU1NBUiIsImVtYWlsIjoibXVkaUBnbWFpbCIsInBob25lTnVtYmVyIjoiKzYyODEzNTQ3NTY5MTIiLCJyb2xlcyI6IkFETUlOIiwiYXV0aFRpbWUiOjE2MzEwNTI3OTA0ODksImV4cCI6MTYzMTA1Mjc5MDQ4OSwiaWF0IjoxNjMxMDUyNzkwNDg5LCJzdWIiOiJqYWxhbi1oaWR1cC0xMjMiLCJhdWgiOiIxMjMiLCJpc3MiOiJqYWxhbi1oaWR1cC1pc3MifQ=="})
	ctx := metadata.NewIncomingContext(context.Background(), header)
	user, err := gohelper.ValidateUserRoles(&ctx, "ADMIN")
	t.Log(user.Email)
	if err != nil {
		t.Error(err)
	}
	// t.Log(user.Id)
}
