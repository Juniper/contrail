package rbac

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
)

// NoAuth is used to create new no auth context
func userAuth(ctx context.Context) context.Context {
	Context := auth.NewContext(
		"default-domain", "default-project", "bob", []string{"Member"})
	var authKey interface{} = "auth"
	return context.WithValue(ctx, authKey, Context)
}

func TestCheckPermissions(t *testing.T) {
	type args struct {
		ctx     context.Context
		l       []*models.APIAccessList
		aaaMode string
		kind    string
		op      Action
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	noAuthCtx := auth.NoAuth(ctx)

	userAuthCtx := userAuth(ctx)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "rbac virtual-network create with no rbac",
			args: args{
				ctx:     noAuthCtx,
				l:       nil,
				aaaMode: "",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
			wantErr: false,
		},
		{
			name: "rbac virtual-network create with  rbac enabled as admin ",
			args: args{
				ctx:     noAuthCtx,
				l:       nil,
				aaaMode: "rbac",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
			wantErr: false,
		},
		{
			name: "rbac virtual-network create with  rbac enabled as Member  with rbac disabled",
			args: args{
				ctx:     userAuthCtx,
				l:       nil,
				aaaMode: "",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
			wantErr: false,
		},

		{
			name: "rbac virtual-network create with  rbac enabled as Member with rbac enabled",
			args: args{
				ctx:     userAuthCtx,
				l:       nil,
				aaaMode: "rbac",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPermissions(tt.args.ctx, tt.args.l, tt.args.aaaMode, tt.args.kind, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
