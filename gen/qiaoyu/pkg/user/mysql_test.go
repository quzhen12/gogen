package user

import (
	"context"
	"testing"

	"qiaoyu/pkg/test"
)

var (
	r   Repository
	ctx context.Context
)

func init() {
	test.Init()

	r = NewMysqlRepo()
	ctx = context.Background()
}

func TestFindUserByUsername(t *testing.T) {
	user, err := r.FindUserByUsername(ctx, "tom")
	if err != nil {
		t.Error(err)
	}
	t.Logf("result: %+v", user)
}

func TestDeleteUserById(t *testing.T) {
	ret, err := r.DeleteUserById(ctx, 12)
	if err != nil {
		t.Error(err)
	}
	t.Logf("result: %+v", ret)
}

func TestCreateUser(t *testing.T) {
	err := r.CreateUser(ctx, "test12", "pp")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	n, err := r.UpdateUser(ctx, &User{Email: "test12", Password: "122"})
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}
