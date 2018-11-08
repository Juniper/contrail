package types

import (
	"context"
	"github.com/Juniper/contrail/pkg/errutil"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

//CreateAlarm does pre check for alarm.
func (sv *ContrailTypeLogicService) CreateAlarm(
	ctx context.Context,
	request *services.CreateAlarmRequest) (response *services.CreateAlarmResponse, err error) {

	alarm := request.GetAlarm()

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			// check if alarm rules are valid
			err = alarm.Validate()
			if err != nil {
				return errutil.ErrorBadRequestf(err.Error())
			}

			response, err = sv.BaseService.CreateAlarm(ctx, request)
			if err != nil {
				return err
			}

			return nil
		})

	return response, err
}

//UpdateAlarm does pre check for alarm.
func (sv *ContrailTypeLogicService) UpdateAlarm(
	ctx context.Context,
	request *services.UpdateAlarmRequest) (response *services.UpdateAlarmResponse, err error) {

	alarm := request.GetAlarm()
	fm := request.GetFieldMask()

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			// check if alarm rules was present in JSON
			if basemodels.FieldMaskContains(&fm, models.AlarmFieldAlarmRules) {

				// check if alarm rules are valid
				err = alarm.Validate()
				if err != nil {
					return errutil.ErrorBadRequestf(err.Error())
				}
			}

			if response, err = sv.BaseService.UpdateAlarm(ctx, request); err != nil {
				return err
			}

			return nil
		})

	return response, err
}
