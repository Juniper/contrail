package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestServiceTemplate(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "service_template")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeServiceTemplate()
	model.UUID = "service_template_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "service_template_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceApplianceSetcreateref []*models.ServiceTemplateServiceApplianceSetRef
	var ServiceApplianceSetrefModel *models.ServiceApplianceSet
	ServiceApplianceSetrefModel = models.MakeServiceApplianceSet()
	ServiceApplianceSetrefModel.UUID = "service_template_service_appliance_set_ref_uuid"
	ServiceApplianceSetrefModel.FQName = []string{"test", "service_template_service_appliance_set_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	ServiceApplianceSetrefModel.UUID = "service_template_service_appliance_set_ref_uuid1"
	ServiceApplianceSetrefModel.FQName = []string{"test", "service_template_service_appliance_set_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	ServiceApplianceSetrefModel.UUID = "service_template_service_appliance_set_ref_uuid2"
	ServiceApplianceSetrefModel.FQName = []string{"test", "service_template_service_appliance_set_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceApplianceSet(ctx, tx, &models.CreateServiceApplianceSetRequest{
			ServiceApplianceSet: ServiceApplianceSetrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.ServiceTemplateServiceApplianceSetRef{UUID: "service_template_service_appliance_set_ref_uuid", To: []string{"test", "service_template_service_appliance_set_ref_uuid"}})
	ServiceApplianceSetcreateref = append(ServiceApplianceSetcreateref, &models.ServiceTemplateServiceApplianceSetRef{UUID: "service_template_service_appliance_set_ref_uuid2", To: []string{"test", "service_template_service_appliance_set_ref_uuid2"}})
	model.ServiceApplianceSetRefs = ServiceApplianceSetcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "service_template_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(ctx, tx, &models.CreateProjectRequest{
			Project: projectModel,
		})
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.VrouterInstanceType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.Version", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.ServiceVirtualizationType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.ServiceType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.ServiceScaling", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.ServiceMode", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.OrderedInterfaces", ".", true)
	//
	//
	//
	//    if ".ServiceTemplateProperties.InterfaceType" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".ServiceTemplateProperties.InterfaceType", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".ServiceTemplateProperties.InterfaceType", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.InstanceData", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.ImageName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.Flavor", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceTemplateProperties.AvailabilityZoneEnable", ".", true)
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "service_template_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var ServiceApplianceSetref []interface{}
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"delete", "uuid":"service_template_service_appliance_set_ref_uuid", "to": []string{"test", "service_template_service_appliance_set_ref_uuid"}})
	//    ServiceApplianceSetref = append(ServiceApplianceSetref, map[string]interface{}{"operation":"add", "uuid":"service_template_service_appliance_set_ref_uuid1", "to": []string{"test", "service_template_service_appliance_set_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "ServiceApplianceSetRefs", ".", ServiceApplianceSetref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceTemplate(ctx, tx,
			&models.CreateServiceTemplateRequest{
				ServiceTemplate: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateServiceTemplate(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_service_template_service_appliance_set` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceApplianceSetRefs delete statement failed")
		}
		_, err = stmt.Exec("service_template_dummy_uuid", "service_template_service_appliance_set_ref_uuid")
		_, err = stmt.Exec("service_template_dummy_uuid", "service_template_service_appliance_set_ref_uuid1")
		_, err = stmt.Exec("service_template_dummy_uuid", "service_template_service_appliance_set_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(ctx, tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "service_template_service_appliance_set_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref service_template_service_appliance_set_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(ctx, tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "service_template_service_appliance_set_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref service_template_service_appliance_set_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceApplianceSet(
			ctx,
			tx,
			&models.DeleteServiceApplianceSetRequest{
				ID: "service_template_service_appliance_set_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref service_template_service_appliance_set_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
			ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListServiceTemplate(ctx, tx, &models.ListServiceTemplateRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ServiceTemplates) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceTemplate(ctxDemo, tx,
			&models.DeleteServiceTemplateRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceTemplate(ctx, tx,
			&models.DeleteServiceTemplateRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceTemplate(ctx, tx,
			&models.CreateServiceTemplateRequest{
				ServiceTemplate: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListServiceTemplate(ctx, tx, &models.ListServiceTemplateRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ServiceTemplates) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
