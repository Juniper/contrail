package cassandra

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/services"
)

// EventProcessor writes events to cassandra and implements service.EventProcessor interface
type EventProcessor struct {
	config Config
}

// NewEventProcessor returns new cassandra.EventProcessor
func NewEventProcessor() *EventProcessor {
	cfg := GetConfig()
	return &EventProcessor{
		config: cfg,
	}
}

// Process is a method needed to implement service.EventProcessor interface
func (p *EventProcessor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	log.Debugf("Processing event %+v for cassandra", event)
	// connect to the cluster
	cluster := getCluster(p.config)
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	if err = handleEvent(session, event); err != nil {
		return nil, err
	}

	return event, nil
}

const (
	selectResource        = "SELECT key, column1, value FROM obj_uuid_table WHERE key = ?"
	insertQuery           = "INSERT INTO obj_uuid_table (key, column1, value) VALUES (?, ?, ?)"
	deleteRowByColumn1    = "DELETE FROM obj_uuid_table WHERE key=? and column1=?"
	deleteResource        = "DELETE FROM obj_uuid_table WHERE key=?"
	insertIntoFQNameTable = "INSERT INTO obj_fq_name_table (key, column1, value) VALUES (?, ?, ?)"
	deleteFromFQNameTable = "DELETE FROM obj_fq_name_table WHERE key=? and column1=?"
	cassandraTimeFormat   = "2006-01-02T15:04:05.999999"
)

func handleEvent(session *gocql.Session, event *services.Event) error { // nolint: interfacer
	rsrc := event.GetResource()
	switch event.Operation() {
	case services.OperationCreate, services.OperationUpdate:
		updateIdPermsOnCreate(rsrc)

		cassandraMap, err := resourceToCassandraMap(rsrc)
		if err != nil {
			return err
		}

		// select the object's children and backrefs from cassandra and update our object map
		iter := session.Query(selectResource, rsrc.GetUUID()).Iter()
		var uuid, column1, value string
		for iter.Scan(&uuid, &column1, &value) {
			if _, ok := cassandraMap[column1]; isRefRow(column1) && !ok {
				// This is a symmetric ref's backref, leave it intact
				cassandraMap[column1] = value
			}

			if isChildrenRow(column1) || isBackrefRow(column1) {
				cassandraMap[column1] = value
			}
			fmt.Println(uuid, column1, value)
		}

		batch := gocql.NewBatch(gocql.LoggedBatch)

		// delete the old object from cassandra
		withTimestamp(batch, deleteResource, rsrc.GetUUID())

		// insert the new version
		for column1, value := range cassandraMap {
			withTimestamp(batch, insertQuery, rsrc.GetUUID(), column1, value)
		}

		// if the object has a parent, update the parent's children
		if parentUUID := rsrc.GetParentUUID(); parentUUID != "" {
			updateChildRow(batch, rsrc)
		}

		// update backrefs in the referred objects
		for column1, value := range cassandraMap {
			if !isRefRow(column1) {
				continue
			}

			data := strings.Split(column1, ":")
			referredType, referredColumn1 := data[1], data[2]
			updateBackref(batch, rsrc, referredType, referredColumn1, value)
		}

		updateFQNameToUUIDRow(batch, rsrc)
		if err = session.ExecuteBatch(batch); err != nil {
			return err
		}
	case services.OperationDelete:
		batch := gocql.NewBatch(gocql.LoggedBatch)
		deleteChildRow(batch, rsrc)
		deleteFQNameToUUIDRow(batch, rsrc)
		withTimestamp(batch, deleteResource, rsrc.GetUUID())
		if err := session.ExecuteBatch(batch); err != nil {
			return err
		}
	}
	return nil
}

func withTimestamp(b *gocql.Batch, stmt string, args ...interface{}) {
	strs := strings.Split(stmt, " WHERE")
	stmt = strings.Join([]string{strs[0], " USING TIMESTAMP ", fmt.Sprint(time.Now().UnixNano())}, "")
	if len(strs) > 1 {
		stmt += " WHERE" + strs[1]
	}
	log.Warn(stmt)
	b.Query(stmt, args...)
}

func childColumn1(r services.Resource) string {
	childType := strings.Replace(r.Kind(), "_", "-", -1)
	return fmt.Sprintf("children:%s:%s", childType, r.GetUUID())
}

func getFQNameToUUIDColumn1(r services.Resource) string {
	return strings.Join(append(r.GetFQName(), r.GetUUID()), ":")
}

func updateIdPermsOnCreate(r services.Resource) {
	p := r.GetIDPerms()
	t := time.Now().Format(cassandraTimeFormat)
	p.Created = t
	p.LastModified = t
}

func deleteFQNameToUUIDRow(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, deleteFromFQNameTable, r.Kind(), getFQNameToUUIDColumn1(r))
}

func createFQNameToUUIDRow(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, insertIntoFQNameTable, r.Kind(), getFQNameToUUIDColumn1(r), "null")
}

func updateFQNameToUUIDRow(b *gocql.Batch, r services.Resource) {
	deleteFQNameToUUIDRow(b, r)
	createFQNameToUUIDRow(b, r)
}

func deleteChildRow(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, deleteRowByColumn1, r.GetParentUUID(), childColumn1(r))
}

func createChildRow(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, insertQuery, r.GetParentUUID(), childColumn1(r), "null")
}

func updateChildRow(b *gocql.Batch, r services.Resource) {
	deleteChildRow(b, r)
	createChildRow(b, r)
}

func updateBackref(b *gocql.Batch, r services.Resource, referredType, referredUUID, value string) {
	var referredColumn1 string
	if r.Kind() == referredType { // TODO: is this correct? _ vs -
		// This is a symmetric ref, we add a "ref:" to the referred object
		referredColumn1 = fmt.Sprintf("ref:%s:%s", r.Kind(), r.GetUUID())
	} else {
		// We add a "backref:" to the referred object
		referredColumn1 = fmt.Sprintf("backref:%s:%s", r.Kind(), r.GetUUID())
	}

	withTimestamp(b, deleteRowByColumn1, referredUUID, referredColumn1)
	withTimestamp(b, insertQuery, referredUUID, referredColumn1, value)
}

func isRefRow(r string) bool {
	return strings.HasPrefix(r, "ref:")
}

func isBackrefRow(r string) bool {
	return strings.HasPrefix(r, "backref:")
}

func isChildrenRow(r string) bool {
	return strings.HasPrefix(r, "children:")
}
