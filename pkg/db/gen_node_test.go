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

func TestNode(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeNode()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var KeypairCreateRef []*models.NodeKeypairRef
	var KeypairRefModel *models.Keypair

	KeypairRefUUID := uuid.NewV4().String()
	KeypairRefUUID1 := uuid.NewV4().String()
	KeypairRefUUID2 := uuid.NewV4().String()

	KeypairRefModel = models.MakeKeypair()
	KeypairRefModel.UUID = KeypairRefUUID
	KeypairRefModel.FQName = []string{"test", KeypairRefUUID}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairRefModel,
	})
	KeypairRefModel.UUID = KeypairRefUUID1
	KeypairRefModel.FQName = []string{"test", KeypairRefUUID1}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairRefModel,
	})
	KeypairRefModel.UUID = KeypairRefUUID2
	KeypairRefModel.FQName = []string{"test", KeypairRefUUID2}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	KeypairCreateRef = append(KeypairCreateRef,
		&models.NodeKeypairRef{UUID: KeypairRefUUID, To: []string{"test", KeypairRefUUID}})
	KeypairCreateRef = append(KeypairCreateRef,
		&models.NodeKeypairRef{UUID: KeypairRefUUID2, To: []string{"test", KeypairRefUUID2}})
	model.KeypairRefs = KeypairCreateRef

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

	_, err = db.CreateNode(ctx,
		&models.CreateNodeRequest{
			Node: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListNode(ctx, &models.ListNodeRequest{
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
	if len(response.Nodes) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteNode(ctxDemo,
		&models.DeleteNodeRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateNode(ctx,
		&models.CreateNodeRequest{
			Node: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteNode(ctx,
		&models.DeleteNodeRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetNode(ctx, &models.GetNodeRequest{
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
