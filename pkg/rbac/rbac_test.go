package rbac

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
)

const apiAccessListUUID = "default-api-access-list8_uuid"
const adminUser = "admin"

func TestCheckCommonPermissions(t *testing.T) {
	type args struct {
		ctx     context.Context
		aaaMode string
		kind    string
		op      Action
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RBAC virtual-network create with no RBAC",
			args: args{
				ctx:     auth.NoAuth(ctx),
				aaaMode: "",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC virtual-network create with RBAC enabled as admin",
			args: args{
				ctx:     auth.NoAuth(ctx),
				aaaMode: "rbac",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC virtual-network create with RBAC enabled as Member  with RBAC disabled",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				aaaMode: "",
				kind:    "virtual-network",
				op:      ActionCreate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckCommonPermissions(tt.args.ctx, tt.args.aaaMode, tt.args.kind, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChecktCommonPermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RBAC project create with  RBAC enabled as Member and global RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       globalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with  RBAC enabled as Member and domain RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       domainAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with  RBAC enabled as Member and project RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       projectAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create  and project RBAC rule and shared permissions",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       projectAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create  and project RBAC rule and shared permissions",
			args: args{
				ctx:     userAuth(ctx, "project_red_uuid"),
				l:       globalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with RBAC enabled as Member and global RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       globalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with RBAC enabled as Member and wildcard global RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       wildcardGlobalAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with RBAC enabled as Member and domain RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
				l:       domainAccessRuleList(),
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionCreate,
			},
		},
		{
			name: "RBAC project create with RBAC enabled as Member and project RBAC rule",
			args: args{
				ctx:     userAuth(ctx, "default-project"),
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
				t.Errorf("ChecktPermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckObjectPermissions(t *testing.T) {
	type args struct {
		ctx     context.Context
		p       *models.PermType2
		aaaMode string
		kind    string
		op      Action
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "resource  access of owner with full permission",
			args: args{
				ctx: userAuth(ctx, "project_blue_uuid"),
				p: &models.PermType2{
					Owner:        "project_blue_uuid",
					OwnerAccess:  7,
					GlobalAccess: 0,
					Share:        []*models.ShareType{{TenantAccess: 7, Tenant: "project:project_red_uuid"}},
				},
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionRead,
			},
		},
		{
			name: "resource read by owner without permission",
			args: args{
				ctx: userAuth(ctx, "project_blue_uuid"),
				p: &models.PermType2{
					Owner:        "project_blue_uuid",
					OwnerAccess:  2,
					GlobalAccess: 0,
					Share:        []*models.ShareType{{TenantAccess: 7, Tenant: "project:project_red_uuid"}},
				},
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionRead,
			},
			wantErr: true,
		},
		{
			name: "resource read by owner without permission",
			args: args{
				ctx: userAuth(ctx, "project_blue_uuid"),
				p: &models.PermType2{
					Owner:        "project_blue_uuid",
					OwnerAccess:  4,
					GlobalAccess: 0,
					Share:        []*models.ShareType{{TenantAccess: 7, Tenant: "project:project_red_uuid"}},
				},
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionRead,
			},
		},

		{
			name: "resource access by the shared project with permission",
			args: args{
				ctx: userAuth(ctx, "project_red_uuid"),
				p: &models.PermType2{
					Owner:        "project_blue_uuid",
					OwnerAccess:  7,
					GlobalAccess: 0,
					Share:        []*models.ShareType{{TenantAccess: 7, Tenant: "project:project_red_uuid"}},
				},
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionUpdate,
			},
		},
		{
			name: "Project access by no shared project",
			args: args{
				ctx: userAuth(ctx, "default-project"),
				p: &models.PermType2{
					Owner:        "project_blue_uuid",
					OwnerAccess:  7,
					GlobalAccess: 0,
					Share:        []*models.ShareType{{TenantAccess: 7, Tenant: "project:project_red_uuid"}},
				},
				aaaMode: "rbac",
				kind:    "project",
				op:      ActionDelete,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckObjectPermissions(tt.args.ctx, tt.args.p, tt.args.aaaMode, tt.args.kind, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func userAuth(ctx context.Context, tenant string) context.Context {
	id := keystone.NewAuthIdentity("default-domain", tenant, "bob", []string{"Member"})
	return auth.WithIdentity(ctx, id)
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

func rbacRuleEntryAdd(l *models.APIAccessList, kind string) {
	m := models.MakeRbacRuleType()
	m.RuleObject = kind
	p := models.MakeRbacPermType()
	p.RoleCrud = "CRUD"
	p.RoleName = "Member"
	m.RulePerms = append(m.RulePerms, p)
	l.APIAccessListEntries.RbacRule = append(l.APIAccessListEntries.RbacRule, m)
}
