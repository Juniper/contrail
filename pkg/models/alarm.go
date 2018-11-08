package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (alarm *Alarm) Validate() error {

	err := alarm.validateAlarmRules()
	if err != nil {
		return err
	}

	err = alarm.validateAlarmExpressions()
	if err != nil {
		return err
	}

	return nil
}

func (alarm *Alarm) validateAlarmRules() error {
	rules := alarm.GetAlarmRules()
	if rules == nil {
		return errors.New("AlarmRules cannot be nil")
	}

	orList := rules.OrList
	if orList == nil {
		return errors.New("OrList cannot be nil")
	}

	for _, andList := range orList {
		if andList == nil {
			return errors.New("AndList cannot be nil")
		}
		for _, andCond := range andList.AndList {
			if andCond.Operand2 == nil {
				return errors.New("Operand2 cannot be nil")
			}
		}
	}

	return nil
}

func (alarm *Alarm) validateAlarmExpressions() error {
	for _, andList := range alarm.GetAlarmRules().OrList {
		for _, andCond := range andList.AndList {
			op := andCond.Operand2
			if op.JSONValue != "" {
				if op.UveAttribute != "" {
					return errors.New("Operand2 should have JSONValue or UveAttribute filled, not both")
				}

				if andCond.Operation == "range" {
					var parsed []int
					err := json.Unmarshal([]byte(op.JSONValue), &parsed)
					if err != nil {
						return errors.Errorf("Couldn't parse JSONValue\n%v", err)
					}

					if len(parsed) != 2 || parsed[0] >= parsed[1] {
						return errors.New("JSONValue should be 2 element integer array [x,y] where x<y")
					}

				} else {
					var parsed interface{}
					err := json.Unmarshal([]byte(op.JSONValue), &parsed)
					if err != nil {
						return errors.Errorf("Couldn't parse JSONValue\n%v", err)
					}
				}
			} else if op.UveAttribute == "" {
				return errors.New("Operand2 should have JSONValue or UveAttribute filled")
			}
		}
	}

	return nil
}
