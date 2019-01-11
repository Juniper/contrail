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

func RBACRuleEntryAdd(l *models.APIAccessList) {

	m := models.MakeRbacRuleType()
	m.RuleObject = "projects"
	p := models.MakeRbacPermType()
	p.RoleCrud = "CRUD"
	p.RoleName = "Member"

	m.RulePerms = append(m.RulePerms, p)
	l.APIAccessListEntries.RbacRule = append(l.APIAccessListEntries.RbacRule, m)

}

func globalAccessRuleListGet() []*models.APIAccessList {

	list := make([]*models.APIAccessList, 0)

	model := models.MakeAPIAccessList()
	model.UUID = "default-api-access-list8_uuid"
	model.DisplayName = "default-api-access-list8_uuid"
	model.FQName = []string{"default-global-system-config", model.UUID}
	model.Perms2.Owner = "admin"
	model.ParentType = models.KindGlobalSystemConfig

	RBACRuleEntryAdd(model)
	list = append(list, model)

	return list
}

func domainAccessRuleListGet() []*models.APIAccessList {

	list := make([]*models.APIAccessList, 0)

	model := models.MakeAPIAccessList()
	model.UUID = "default-api-access-list8_uuid"
	model.DisplayName = "default-api-access-list8_uuid"
	model.FQName = []string{"default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	model.ParentType = models.KindDomain

	RBACRuleEntryAdd(model)
	list = append(list, model)

	return list
}

func projectAccessRuleListGet() []*models.APIAccessList {

	list := make([]*models.APIAccessList, 0)

	model := models.MakeAPIAccessList()
	model.UUID = "default-api-access-list8_uuid"
	model.DisplayName = "default-api-access-list8_uuid"
	model.FQName = []string{"default-domain", "default-project", model.UUID}
	model.Perms2.Owner = "admin"
	model.ParentType = models.KindProject

	RBACRuleEntryAdd(model)
	list = append(list, model)

	return list
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

	gl := globalAccessRuleListGet()
	dl := domainAccessRuleListGet()
	pl := projectAccessRuleListGet()

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
			name: "rbac project create with  rbac enabled as Member and global RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       gl,
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
			wantErr: false,
		},
		{
			name: "rbac project create with  rbac enabled as Member and domain RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       dl,
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
			wantErr: false,
		},
		{
			name: "rbac project create with  rbac enabled as Member and project RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       pl,
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
			wantErr: false,
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
