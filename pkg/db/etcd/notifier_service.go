package etcd

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that performs writes to etcd.
type NotifierService struct {
	services.BaseService
}

// NewNotifierService creates a etcd Notifier Service.
func NewNotifierService(path string, codec models.Codec) (services.Service, error) {
	ec, err := NewClientByViper("etcd-notifier")
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to etcd server")
	}
	return &services.EventProducerService{
		Processor: &Processor{
			Path:   path,
			Client: ec,
			Codec:  codec,
			log:    logutil.NewLogger("etcd-notifier"),
		},
	}, nil
}

type Processor struct {
	Path   string
	Client *Client
	Codec  models.Codec
	log    *logrus.Entry
}

func (p *Processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	switch event.Operation() {
	case services.OperationCreate:
		return p.processCreate(ctx, event)
	case services.OperationUpdate:
		return p.processUpdate(ctx, event)
	case services.OperationDelete:
		return p.processDelete(ctx, event)
	}

	return nil, errors.Errorf("etcd notifier does not support event '%s'", event.Operation())
}

func (p *Processor) processCreate(ctx context.Context, event *services.Event) (*services.Event, error) {
	r := event.GetResource()
	if r == nil {
		return nil, errors.Errorf("got event with nil resource: %v", event)
	}
	key := models.ResourceKey(basemodels.KindToSchemaID(r.Kind()), r.GetUUID())

	jsonStr, err := p.Codec.Encode(r)
	if err != nil {
		return nil, errors.New("error encoding create data")
	}

	err = p.Client.Put(ctx, key, jsonStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create %s with UUID %q in etcd", r.Kind(), r.GetUUID())
	}

	p.log.WithField("uuid", r.GetUUID()).Debugf("Created %s in etcd", r.Kind())
	return event, nil
}

func (p *Processor) processUpdate(ctx context.Context, event *services.Event) (*services.Event, error) {
	r := event.GetResource()
	if r == nil {
		return nil, errors.Errorf("got event with nil resource: %v", event)
	}
	key := models.ResourceKey(basemodels.KindToSchemaID(r.Kind()), r.GetUUID())

	p.log.WithField("uuid", r.GetUUID()).Debug("Updating %s in etcd", r.Kind())
	return event, p.Client.InTransaction(ctx, func(ctx context.Context) error {
		txn := GetTxn(ctx)
		oldData := txn.Get(key)
		// TODO fieldmask
		newData, err := models.UpdateData(p.Codec, oldData, r, types.FieldMask{})
		if err != nil {
			return errors.Wrap(err, "error processing update data for etcd")
		}
		txn.Put(key, newData)
		return nil
	})
}

func (p *Processor) processDelete(ctx context.Context, event *services.Event) (*services.Event, error) {
	r := event.GetResource()
	if r == nil {
		return nil, errors.Errorf("got event with nil resource: %v", event)
	}
	err := p.Client.Delete(ctx, models.ResourceKey(basemodels.KindToSchemaID(r.Kind()), r.GetUUID()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to delete %s with UUID %q in etcd", r.Kind(), r.GetUUID())
	}

	p.log.WithField("uuid", r.GetUUID()).Debugf("Deleted %s in etcd", r.Kind())
	return event, nil
}

//func (p *Processor) processCreateRef(ctx context.Context, event *services.Event) (*services.Event, error) {
//	response, err := ns.BaseService.Create{{ refType }}(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//
//	ref := response.Get{{ refType }}()
//	fromKey := models.ResourceKey("{{ schema.ID }}", response.ID)
//	toKey := models.ResourceKey("{{ reference.LinkTo.ID }}", ref.UUID)
//
//	ns.log.WithField("from-key", fromKey).WithField("to-key", toKey).Debug("Creating {{ refType }} in etcd")
//	return response, ns.Client.InTransaction(ctx, func(ctx context.Context) error {
//		newFrom, newTo := &models.{{ schema.JSONSchema.GoName }}{}, &models.{{ reference.GoName }}{}
//		err := ns.handleRefWrapper(ctx, fromKey, newFrom, func() {
//			newFrom.Add{{ reference.GoName }}Ref(ref)
//		})
//		if err != nil {
//			return err
//		}
//
//		return ns.handleRefWrapper(ctx, toKey, newTo, func() {
//			newTo.Add{{ schema.JSONSchema.GoName }}Backref(&models.{{ schema.JSONSchema.GoName }}{UUID: response.ID})
//		})
//	})
//}
//
//func (p *Processor) processDeleteRef(ctx context.Context, event *services.Event) (*services.Event, error) {
//	response, err := ns.BaseService.Delete{{ refType }}(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//
//	ref := response.Get{{ refType }}()
//	fromKey := models.ResourceKey("{{ schema.ID }}", response.ID)
//	toKey := models.ResourceKey("{{ reference.LinkTo.ID }}", ref.UUID)
//
//	ns.log.WithField("from-key", fromKey).WithField("to-key", toKey).Debug("Deleting {{ refType }} in etcd")
//	return response, ns.Client.InTransaction(ctx, func(ctx context.Context) error {
//		newFrom, newTo := &models.{{ schema.JSONSchema.GoName }}{}, &models.{{ reference.GoName }}{}
//		err := ns.handleRefWrapper(ctx, fromKey, newFrom, func() {
//			newFrom.Remove{{ reference.GoName }}Ref(ref)
//		})
//		if err != nil {
//			return err
//		}
//
//		return ns.handleRefWrapper(ctx, toKey, newTo, func() {
//			newTo.Remove{{ schema.JSONSchema.GoName }}Backref(&models.{{ schema.JSONSchema.GoName }}{UUID: response.ID})
//		})
//	})
//}
//
//func (p *Processor) handleRefWrapper(
//	ctx context.Context, key string, obj proto.Message, handleRef func(),
//) error {
//	txn := GetTxn(ctx)
//	oldData := txn.Get(key)
//	if len(oldData) == 0 {
//		return nil
//	}
//
//	sObj := models.NewSerializedObject(oldData, obj, p.Codec)
//	if err := sObj.Map(handleRef); err != nil {
//		return err
//	}
//
//	txn.Put(key, sObj.GetData())
//	return nil
//}
