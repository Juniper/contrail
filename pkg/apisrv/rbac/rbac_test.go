package rbac

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
)

const apiAccessListUUID = "default-api-access-list8_uuid"
const adminUser = "admin"

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
		},
		{
			name: "rbac virtual-network create with rbac enabled as admin",
			args: args{
				ctx:     noAuthCtx,
				l:       nil,
				aaaMode: "rbac",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
		},
		{
			name: "rbac virtual-network create with rbac enabled as Member  with rbac disabled",
			args: args{
				ctx:     userAuthCtx,
				l:       nil,
				aaaMode: "",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
		},
		{
			name: "rbac project create with rbac enabled and no RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       nil,
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
			wantErr: true,
		},
		{
			name: "rbac project create with rbac enabled as Member and global RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       globalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "rbac project create with rbac enabled as Member and wildcard global RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       wildcardGlobalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "rbac project create with rbac enabled as Member and domain RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       domainAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "rbac project create with rbac enabled as Member and project RBAC rule",
			args: args{
				ctx:     userAuthCtx,
				l:       projectAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
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

func rbacRuleEntryAdd(l *models.APIAccessList, kind string) {
	m := models.MakeRbacRuleType()
	m.RuleObject = kind
	p := models.MakeRbacPermType()
	p.RoleCrud = "CRUD"
	p.RoleName = "Member"
	m.RulePerms = append(m.RulePerms, p)
	l.APIAccessListEntries.RbacRule = append(l.APIAccessListEntries.RbacRule, m)
}

func globalAccessRuleList() []*models.APIAccessList {
	list := make([]*models.APIAccessList, 0)
	model := models.MakeAPIAccessList()
	model.UUID = apiAccessListUUID
	model.FQName = []string{"default-global-system-config", model.UUID}
	model.Perms2.Owner = adminUser
	model.ParentType = models.KindGlobalSystemConfig
	rbacRuleEntryAdd(model, "projects")
	return append(list, model)
}

func wildcardGlobalAccessRuleList() []*models.APIAccessList {
	list := make([]*models.APIAccessList, 0)
	model := models.MakeAPIAccessList()
	model.UUID = apiAccessListUUID
	model.FQName = []string{"default-global-system-config", model.UUID}
	model.Perms2.Owner = adminUser
	model.ParentType = models.KindGlobalSystemConfig
	rbacRuleEntryAdd(model, "*")
	return append(list, model)
}

func domainAccessRuleList() []*models.APIAccessList {
	list := make([]*models.APIAccessList, 0)
	model := models.MakeAPIAccessList()
	model.UUID = apiAccessListUUID
	model.FQName = []string{"default-domain", model.UUID}
	model.Perms2.Owner = adminUser
	model.ParentType = models.KindDomain
	rbacRuleEntryAdd(model, "projects")
	return append(list, model)
}

func projectAccessRuleList() []*models.APIAccessList {
	list := make([]*models.APIAccessList, 0)
	model := models.MakeAPIAccessList()
	model.UUID = apiAccessListUUID
	model.FQName = []string{"default-domain", "default-project", model.UUID}
	model.Perms2.Owner = adminUser
	model.ParentType = models.KindProject
	rbacRuleEntryAdd(model, "projects")
	return append(list, model)
}
