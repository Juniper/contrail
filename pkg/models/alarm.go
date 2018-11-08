package models

import (
	"encoding/json"

	"github.com/Juniper/contrail/pkg/errutil"
)

func (alarm *Alarm) CheckAlarmRules() error {
	rules := alarm.GetAlarmRules()
	if rules == nil {
		return errutil.ErrorBadRequestf("AlarmRules cannot be nil")
	}

	orList := rules.OrList
	if orList == nil {
		return errutil.ErrorBadRequestf("OrList cannot be nil")
	}

	for _, andList := range orList {
		if andList == nil {
			return errutil.ErrorBadRequestf("AndList cannot be nil")
		}
		for _, andCond := range andList.AndList {
			if andCond.Operand2 == nil {
				return errutil.ErrorBadRequestf("Operand2 cannot be nil")
			}
		}
	}

	return nil
}

func (alarm *Alarm) CheckAlarmExpressions() error {
	for _, andList := range alarm.GetAlarmRules().OrList {
		for _, andCond := range andList.AndList {
			op := andCond.Operand2
			if op.JSONValue != "" {
				if op.UveAttribute != "" {
					return errutil.ErrorBadRequestf("Operand2 should have JSONValue or UveAttribute filled, not both")
				}

				if andCond.Operation == "range" {
					var parsed []int
					err := json.Unmarshal([]byte(op.JSONValue), &parsed)
					if err != nil {
						return errutil.ErrorBadRequestf("Couldn't parse JSONValue\n"+
							"Value should be 2 element integer array [x,y] where x<y%v", err)
					}

					if len(parsed) != 2 || parsed[0] >= parsed[1] {
						return errutil.ErrorBadRequestf("JSONValue should be 2 element integer array [x,y] where x<y")
					}

				} else {
					var parsed interface{}
					err := json.Unmarshal([]byte(op.JSONValue), &parsed)
					if err != nil {
						return errutil.ErrorBadRequestf("Couldn't parse JSONValue\n%v", err)
					}
				}
			} else if op.UveAttribute == "" {
				return errutil.ErrorBadRequestf("Operand2 should have JSONValue or UveAttribute filled")
			}
		}
	}

	return nil
}
