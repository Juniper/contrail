// nolint
package db

import (
	"context"
	"github.com/satori/go.uuid"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestProject(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeProject()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var AliasIPPoolCreateRef []*models.ProjectAliasIPPoolRef
	var AliasIPPoolRefModel *models.AliasIPPool

	AliasIPPoolRefUUID := uuid.NewV4().String()
	AliasIPPoolRefUUID1 := uuid.NewV4().String()
	AliasIPPoolRefUUID2 := uuid.NewV4().String()

	AliasIPPoolRefModel = models.MakeAliasIPPool()
	AliasIPPoolRefModel.UUID = AliasIPPoolRefUUID
	AliasIPPoolRefModel.FQName = []string{"test", AliasIPPoolRefUUID}
	_, err = db.CreateAliasIPPool(ctx, &models.CreateAliasIPPoolRequest{
		AliasIPPool: AliasIPPoolRefModel,
	})
	AliasIPPoolRefModel.UUID = AliasIPPoolRefUUID1
	AliasIPPoolRefModel.FQName = []string{"test", AliasIPPoolRefUUID1}
	_, err = db.CreateAliasIPPool(ctx, &models.CreateAliasIPPoolRequest{
		AliasIPPool: AliasIPPoolRefModel,
	})
	AliasIPPoolRefModel.UUID = AliasIPPoolRefUUID2
	AliasIPPoolRefModel.FQName = []string{"test", AliasIPPoolRefUUID2}
	_, err = db.CreateAliasIPPool(ctx, &models.CreateAliasIPPoolRequest{
		AliasIPPool: AliasIPPoolRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	AliasIPPoolCreateRef = append(AliasIPPoolCreateRef,
		&models.ProjectAliasIPPoolRef{UUID: AliasIPPoolRefUUID, To: []string{"test", AliasIPPoolRefUUID}})
	AliasIPPoolCreateRef = append(AliasIPPoolCreateRef,
		&models.ProjectAliasIPPoolRef{UUID: AliasIPPoolRefUUID2, To: []string{"test", AliasIPPoolRefUUID2}})
	model.AliasIPPoolRefs = AliasIPPoolCreateRef

	var NamespaceCreateRef []*models.ProjectNamespaceRef
	var NamespaceRefModel *models.Namespace

	NamespaceRefUUID := uuid.NewV4().String()
	NamespaceRefUUID1 := uuid.NewV4().String()
	NamespaceRefUUID2 := uuid.NewV4().String()

	NamespaceRefModel = models.MakeNamespace()
	NamespaceRefModel.UUID = NamespaceRefUUID
	NamespaceRefModel.FQName = []string{"test", NamespaceRefUUID}
	_, err = db.CreateNamespace(ctx, &models.CreateNamespaceRequest{
		Namespace: NamespaceRefModel,
	})
	NamespaceRefModel.UUID = NamespaceRefUUID1
	NamespaceRefModel.FQName = []string{"test", NamespaceRefUUID1}
	_, err = db.CreateNamespace(ctx, &models.CreateNamespaceRequest{
		Namespace: NamespaceRefModel,
	})
	NamespaceRefModel.UUID = NamespaceRefUUID2
	NamespaceRefModel.FQName = []string{"test", NamespaceRefUUID2}
	_, err = db.CreateNamespace(ctx, &models.CreateNamespaceRequest{
		Namespace: NamespaceRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NamespaceCreateRef = append(NamespaceCreateRef,
		&models.ProjectNamespaceRef{UUID: NamespaceRefUUID, To: []string{"test", NamespaceRefUUID}})
	NamespaceCreateRef = append(NamespaceCreateRef,
		&models.ProjectNamespaceRef{UUID: NamespaceRefUUID2, To: []string{"test", NamespaceRefUUID2}})
	model.NamespaceRefs = NamespaceCreateRef

	var ApplicationPolicySetCreateRef []*models.ProjectApplicationPolicySetRef
	var ApplicationPolicySetRefModel *models.ApplicationPolicySet

	ApplicationPolicySetRefUUID := uuid.NewV4().String()
	ApplicationPolicySetRefUUID1 := uuid.NewV4().String()
	ApplicationPolicySetRefUUID2 := uuid.NewV4().String()

	ApplicationPolicySetRefModel = models.MakeApplicationPolicySet()
	ApplicationPolicySetRefModel.UUID = ApplicationPolicySetRefUUID
	ApplicationPolicySetRefModel.FQName = []string{"test", ApplicationPolicySetRefUUID}
	_, err = db.CreateApplicationPolicySet(ctx, &models.CreateApplicationPolicySetRequest{
		ApplicationPolicySet: ApplicationPolicySetRefModel,
	})
	ApplicationPolicySetRefModel.UUID = ApplicationPolicySetRefUUID1
	ApplicationPolicySetRefModel.FQName = []string{"test", ApplicationPolicySetRefUUID1}
	_, err = db.CreateApplicationPolicySet(ctx, &models.CreateApplicationPolicySetRequest{
		ApplicationPolicySet: ApplicationPolicySetRefModel,
	})
	ApplicationPolicySetRefModel.UUID = ApplicationPolicySetRefUUID2
	ApplicationPolicySetRefModel.FQName = []string{"test", ApplicationPolicySetRefUUID2}
	_, err = db.CreateApplicationPolicySet(ctx, &models.CreateApplicationPolicySetRequest{
		ApplicationPolicySet: ApplicationPolicySetRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ApplicationPolicySetCreateRef = append(ApplicationPolicySetCreateRef,
		&models.ProjectApplicationPolicySetRef{UUID: ApplicationPolicySetRefUUID, To: []string{"test", ApplicationPolicySetRefUUID}})
	ApplicationPolicySetCreateRef = append(ApplicationPolicySetCreateRef,
		&models.ProjectApplicationPolicySetRef{UUID: ApplicationPolicySetRefUUID2, To: []string{"test", ApplicationPolicySetRefUUID2}})
	model.ApplicationPolicySetRefs = ApplicationPolicySetCreateRef

	var FloatingIPPoolCreateRef []*models.ProjectFloatingIPPoolRef
	var FloatingIPPoolRefModel *models.FloatingIPPool

	FloatingIPPoolRefUUID := uuid.NewV4().String()
	FloatingIPPoolRefUUID1 := uuid.NewV4().String()
	FloatingIPPoolRefUUID2 := uuid.NewV4().String()

	FloatingIPPoolRefModel = models.MakeFloatingIPPool()
	FloatingIPPoolRefModel.UUID = FloatingIPPoolRefUUID
	FloatingIPPoolRefModel.FQName = []string{"test", FloatingIPPoolRefUUID}
	_, err = db.CreateFloatingIPPool(ctx, &models.CreateFloatingIPPoolRequest{
		FloatingIPPool: FloatingIPPoolRefModel,
	})
	FloatingIPPoolRefModel.UUID = FloatingIPPoolRefUUID1
	FloatingIPPoolRefModel.FQName = []string{"test", FloatingIPPoolRefUUID1}
	_, err = db.CreateFloatingIPPool(ctx, &models.CreateFloatingIPPoolRequest{
		FloatingIPPool: FloatingIPPoolRefModel,
	})
	FloatingIPPoolRefModel.UUID = FloatingIPPoolRefUUID2
	FloatingIPPoolRefModel.FQName = []string{"test", FloatingIPPoolRefUUID2}
	_, err = db.CreateFloatingIPPool(ctx, &models.CreateFloatingIPPoolRequest{
		FloatingIPPool: FloatingIPPoolRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FloatingIPPoolCreateRef = append(FloatingIPPoolCreateRef,
		&models.ProjectFloatingIPPoolRef{UUID: FloatingIPPoolRefUUID, To: []string{"test", FloatingIPPoolRefUUID}})
	FloatingIPPoolCreateRef = append(FloatingIPPoolCreateRef,
		&models.ProjectFloatingIPPoolRef{UUID: FloatingIPPoolRefUUID2, To: []string{"test", FloatingIPPoolRefUUID2}})
	model.FloatingIPPoolRefs = FloatingIPPoolCreateRef

	//create project to which resource is shared
	projectModel := models.MakeProject()

	projectModel.UUID = uuid.NewV4().String()
	projectModel.FQName = []string{"default-domain-test", projectModel.UUID}
	projectModel.Perms2.Owner = "admin"

	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:" + projectModel.UUID, TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	_, err = db.CreateProject(ctx,
		&models.CreateProjectRequest{
			Project: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListProject(ctx, &models.ListProjectRequest{
		Spec: &models.ListSpec{Limit: 1,
			Filters: []*models.Filter{
				&models.Filter{
					Key:    "uuid",
					Values: []string{model.UUID},
				},
			},
		}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Projects) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteProject(ctxDemo,
		&models.DeleteProjectRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateProject(ctx,
		&models.CreateProjectRequest{
			Project: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteProject(ctx,
		&models.DeleteProjectRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetProject(ctx, &models.GetProjectRequest{
		ID: model.UUID})
	if err == nil {
		t.Fatal("expected not found error")
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}
	return
}
