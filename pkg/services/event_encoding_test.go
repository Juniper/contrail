package services

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		option *EventOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name: "try to create event with empty option",
			args: args{
				option: &EventOption{},
			},
			wantErr: true,
		},
		{
			name: "create event with default (Create) operation",
			args: args{
				option: &EventOption{
					Kind: models.KindProject,
				},
			},
			want: &Event{
				Request: &Event_CreateProjectRequest{
					CreateProjectRequest: &CreateProjectRequest{
						Project: &models.Project{},
					},
				},
			},
		},
		{
			name: "create event with Create operation",
			args: args{
				option: &EventOption{
					Operation: OperationCreate,
					Kind:      models.KindProject,
					Data: map[string]interface{}{
						"name": "hoge",
					},
				},
			},
			want: &Event{
				Request: &Event_CreateProjectRequest{
					CreateProjectRequest: &CreateProjectRequest{
						Project: &models.Project{
							Name: "hoge",
						},
					},
				},
			},
		},
		{
			name: "create event with Update operation",
			args: args{
				option: &EventOption{
					Operation: OperationUpdate,
					Kind:      models.KindProject,
					UUID:      "hoge",
					FieldMask: &types.FieldMask{
						Paths: []string{"name"},
					},
					Data: map[string]interface{}{
						"name": "hoge",
					},
				},
			},
			want: &Event{
				Request: &Event_UpdateProjectRequest{
					UpdateProjectRequest: &UpdateProjectRequest{
						Project: &models.Project{
							UUID: "hoge",
							Name: "hoge",
						},
						FieldMask: types.FieldMask{
							Paths: []string{"name"},
						},
					},
				},
			},
		},
		{
			name: "create event with Delete operation",
			args: args{
				option: &EventOption{
					Operation: OperationDelete,
					Kind:      models.KindProject,
					UUID:      "hoge",
				},
			},
			want: &Event{
				Request: &Event_DeleteProjectRequest{
					DeleteProjectRequest: &DeleteProjectRequest{
						ID: "hoge",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEvent(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want, fmt.Sprintf("NewEvent() go:\n%v\nwant:\n%v", got, tt.want))
		})
	}
}

func TestEvent_ToMap(t *testing.T) {
	tests := []struct {
		name    string
		Request isEvent_Request
		want    map[string]interface{}
	}{
		{
			name: "empty event to map",
		},
		{
			name: "create event to map",
			Request: &Event_CreateProjectRequest{
				CreateProjectRequest: &CreateProjectRequest{
					Project: &models.Project{
						UUID: "hoge",
					},
				},
			},
			want: map[string]interface{}{
				"kind":      models.KindProject,
				"operation": OperationCreate,
				"data": &Project{
					UUID: "hoge",
				},
			},
		},
		{
			name: "update event to map",
			Request: &Event_UpdateProjectRequest{
				UpdateProjectRequest: &UpdateProjectRequest{
					Project: &models.Project{
						UUID: "hoge",
					},
				},
			},
			want: map[string]interface{}{
				"kind":      models.KindProject,
				"operation": OperationUpdate,
				"data": &Project{
					UUID: "hoge",
				},
			},
		},
		{
			name: "delete event to map",
			Request: &Event_DeleteProjectRequest{
				DeleteProjectRequest: &DeleteProjectRequest{
					ID: "hoge",
				},
			},
			want: map[string]interface{}{
				"kind":      models.KindProject,
				"operation": OperationDelete,
				"data": map[string]interface{}{
					"uuid": "hoge",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Request: tt.Request,
			}
			got := e.ToMap()
			assert.Equal(t, len(tt.want), len(got),
				fmt.Sprintf("Event.ToMap() returned invalid number of keys, \nexpected:\n%v\ngot\n%v", tt.want, got))
			for k, v := range tt.want {
				gotValue, ok := got[k]
				if assert.True(t, ok, fmt.Sprintf("missing key: %s with value: %v", k, v)) {
					assert.Equal(t, v, gotValue, "value under key: %s not equal to expected: %v", k, gotValue)
				}
			}
		})
	}
}
