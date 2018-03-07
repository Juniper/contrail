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

func TestForwardingClass(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeForwardingClass()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var QosQueueCreateRef []*models.ForwardingClassQosQueueRef
	var QosQueueRefModel *models.QosQueue

	QosQueueRefUUID := uuid.NewV4().String()
	QosQueueRefUUID1 := uuid.NewV4().String()
	QosQueueRefUUID2 := uuid.NewV4().String()

	QosQueueRefModel = models.MakeQosQueue()
	QosQueueRefModel.UUID = QosQueueRefUUID
	QosQueueRefModel.FQName = []string{"test", QosQueueRefUUID}
	_, err = db.CreateQosQueue(ctx, &models.CreateQosQueueRequest{
		QosQueue: QosQueueRefModel,
	})
	QosQueueRefModel.UUID = QosQueueRefUUID1
	QosQueueRefModel.FQName = []string{"test", QosQueueRefUUID1}
	_, err = db.CreateQosQueue(ctx, &models.CreateQosQueueRequest{
		QosQueue: QosQueueRefModel,
	})
	QosQueueRefModel.UUID = QosQueueRefUUID2
	QosQueueRefModel.FQName = []string{"test", QosQueueRefUUID2}
	_, err = db.CreateQosQueue(ctx, &models.CreateQosQueueRequest{
		QosQueue: QosQueueRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	QosQueueCreateRef = append(QosQueueCreateRef,
		&models.ForwardingClassQosQueueRef{UUID: QosQueueRefUUID, To: []string{"test", QosQueueRefUUID}})
	QosQueueCreateRef = append(QosQueueCreateRef,
		&models.ForwardingClassQosQueueRef{UUID: QosQueueRefUUID2, To: []string{"test", QosQueueRefUUID2}})
	model.QosQueueRefs = QosQueueCreateRef

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

	_, err = db.CreateForwardingClass(ctx,
		&models.CreateForwardingClassRequest{
			ForwardingClass: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListForwardingClass(ctx, &models.ListForwardingClassRequest{
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
	if len(response.ForwardingClasss) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteForwardingClass(ctxDemo,
		&models.DeleteForwardingClassRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateForwardingClass(ctx,
		&models.CreateForwardingClassRequest{
			ForwardingClass: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteForwardingClass(ctx,
		&models.DeleteForwardingClassRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetForwardingClass(ctx, &models.GetForwardingClassRequest{
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
