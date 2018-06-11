package compilation

import (
	"testing"
)

func TestIntentCompilationService_HandleMessage(t *testing.T) {
	ics := &IntentCompilationService{
		Etcd:    tt.fields.Etcd,
		Cfg:     tt.fields.Cfg,
		Service: tt.fields.Service,
		locker:  tt.fields.locker,
	}
	ics.HandleMessage(tt.args.ctx, tt.args.index, tt.args.oper, tt.args.key, tt.args.newValue)
}
