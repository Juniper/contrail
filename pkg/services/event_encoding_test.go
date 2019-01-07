package services

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/gogo/protobuf/types"
	"github.com/magiconair/properties/assert"
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
