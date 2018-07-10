package compilation_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestIntentCompilationServiceProcessesBasicResourcesCreation(t *testing.T) {
	tests := []struct {
		dbDriver string
	}{
		{dbDriver: db.DriverMySQL},
		{dbDriver: db.DriverPostgreSQL},
	}

	// TODO: spawn IC

	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	for _, tt := range tests {
		t.Run(tt.dbDriver, func(t *testing.T)) {
			s := integration.NewRunningAPIServer(t, "../..", tt.dbDriver)
			defer s.Close(t)

			hc := integration.NewHTTPAPIClient(t, s.URL())

			// TODO: create project
			projectUUID := t.Name()+"-project"
			hc.CreateProject(t, project(projectUUID))

			// TODO: check project created in etcd

			// TODO: check acl 1 in etcd

			// TODO: check acl 2 in etcd

			// TODO: check application_policy_set in etcd

			// TODO: check security_group in etcd

			// TODO: create virtual network

			// TODO: check virtual_network in etcd

			// TODO: check route_target in etcd

			// TODO: check routing instance in etcd

			// TODO: create subnet

			// TODO: virtual network updated with network_ipam_refs in etcd

			// TODO: create virtual machine

			// TODO: check all resources in etcd after virtual machine create
		}
	}
}

//compilation:
//plugin_directory: 'etc/plugins/'
//    number_of_workers: 3
//max_job_queue_len: 3
//msg_queue_lock_time: 30
//msg_index_string: 'MsgIndex'
//    read_lock_string: "MsgReadLock"
//master_election: true
//
//# Plugin configuration
//plugin:
//handlers:
//create_handler: 'TestHandleCreate'
//            update_handler: 'TestHandleUpdate'
//            delete_handler: 'TestHandleDelete'

func project(uuid string) *models.Project {
	return &models.Project{
		UUID:       uuid,
		ParentType: integration.DomainType,
		ParentUUID: integration.DefaultDomainUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Quota:      &models.QuotaType{},
	}
}
