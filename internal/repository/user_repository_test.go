package repository

import (
	"context"
	"errors"
	"github.com/egorgasay/dockerdb/v3"
	"go-rest-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func Test_userRepository_CreateUser(t *testing.T) {
	invalidErr := errors.New("invalid")
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Valid user",

			args: args{
				user: model.User{
					Username: "dima",
					Password: "test",
					Session:  "ahjsuiwlght-12",
				},
			},
			wantErr: nil,
		},
		{
			name: "BAD",
			args: args{
				user: model.User{
					Username: "dima",
				},
			},
			wantErr: invalidErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			vdb, err := dockerdb.New(ctx, dockerdb.PostgresConfig("market").Build())
			if err != nil {
				log.Fatal(err)
			}
			defer vdb.Clear(ctx)

			gormDB, err := gorm.Open(postgres.Open(vdb.GetSQLConnStr()), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}

			if err := gormDB.Create(tt.args.user); err != nil {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
