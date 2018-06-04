/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * Contrail Plugin Implementation
 *  The dependencies processor
 *  - object dependencies are evaluated when objects are
 *    created/modified or deleted
 *
 * TODO: Use this dependency tracker code in contrail plugin
 *       implementation. Plug it in the service pipeline
 */

package dependencies

import (
	"encoding/json"
	"reflect"
	"sync"

	"gopkg.in/oleiade/reflections.v1"

	log "github.com/sirupsen/logrus"
)

// NewDependencyProcessor - creates new instance
func NewDependencyProcessor(objCache map[string]map[string]interface{}) *DependencyProcessor {
	d := &DependencyProcessor{cache: objCache}
	d.Init()
	return d
}

// DependencyProcessor stores resources dependency
type DependencyProcessor struct {
	resources *sync.Map
	cache map[string]map[string]interface{}
}

// Init - initializes the dependency processor
func (d *DependencyProcessor) Init() {
	d.resources = &sync.Map{}
}

// Add - Add object to dependency processor list
func (d *DependencyProcessor) Add(key string, obj interface{}) {
	if _, ok := d.resources.Load(key); !ok {
		d.resources.Store(key, &sync.Map{})
	}
	if objMap, ok := d.resources.Load(key); ok {
		objMap.(*sync.Map).Store(d.getUUID(obj), obj)
		log.Infof("Adding %s(%s)", key, d.getUUID(obj))
	}
}

// GetResources - Get Resources from the dependency processor list
func (d *DependencyProcessor) GetResources() *sync.Map {
	return d.resources
}

// GetResourcesPretty - Get Resources from the dependency processor list
func (d *DependencyProcessor) GetResourcesPretty() string {
	tmpMap := make(map[string]map[string]interface{})
	d.resources.Range(func(k, v interface{}) bool {
		key := k.(string)
		value := v.(*sync.Map)
		tmpMap[key] = make(map[string]interface{})
		value.Range(func(k1, v1 interface{}) bool {
			tmpMap[key][k1.(string)] = v1
			return true
		})
		return true
	})

	b, err := json.MarshalIndent(tmpMap, "", "  ")
	if err != nil {
		log.Errorf("error marshalling: %v", err)
	}
	return string(b)
}

// getUUID gets the uuid string from the interface{} object
func (d *DependencyProcessor) getUUID(obj interface{}) string {
	uuid, err := reflections.GetField(obj, "UUID")
	if err == nil {
		return uuid.(string)
	}
	return ""
}

// canAdd checks if object can be been added to the dependency list
func (d *DependencyProcessor) canAdd(key string, obj interface{}) bool {
	if objMap, ok := d.resources.Load(key); ok {
		_, there := objMap.(*sync.Map).Load(d.getUUID(obj))
		if there == true {
			log.Infof("%s exists, not adding", d.getUUID(obj))
		}
		return !there
	}
	return true
}

// Evaluate - Evaluates object dependency based on the ReactionMap
func (d *DependencyProcessor) Evaluate(obj interface{}, objTypeStr, fromTypeStr string) {
	if _, ok := ReactionMap[objTypeStr]; !ok {
		return
	}

	log.Infof("Evaluating: Object: %s(%s) From: %s", objTypeStr, d.getUUID(obj), fromTypeStr)
	if !d.canAdd(objTypeStr, obj) {
		return
	}
	d.Add(objTypeStr, obj)

	for _, refObjTypeStr := range(ReactionMap[objTypeStr][fromTypeStr]) {
		fieldsToExtract := []string{refObjTypeStr+"Refs"}
		for _, fieldName := range fieldsToExtract {
			refObjTypeValues, err := reflections.GetField(obj, fieldName)
			if err != nil {
				// Refs dont exit, ignore
				continue
			}
			objValues := reflect.ValueOf(refObjTypeValues)
			for i := 0; i < objValues.Len(); i++ {
				interfaceObj := objValues.Index(i).Elem().Interface()
				uuid, _ := reflections.GetField(interfaceObj, "UUID")
				refObj := d.cache[refObjTypeStr][uuid.(string)]
				log.Infof("Evaluating: Object: %s %s(%s) From: %s", fieldName, refObjTypeStr, uuid.(string), objTypeStr)
				d.Evaluate(refObj, refObjTypeStr, objTypeStr)
			}
		}
		fieldsToExtract = []string{refObjTypeStr+"BackRefs"}
		for _, fieldName := range fieldsToExtract {
			refObjTypeValues, err := reflections.GetField(obj, fieldName)
			if err != nil {
				// BackRefs dont exit, ignore
				continue
			}
			objValues := reflect.ValueOf(refObjTypeValues)
			for i := 0; i < objValues.Len(); i++ {
				refObj := objValues.Index(i).Elem().Interface()
				log.Infof("Evaluating: Object: %s %s(%s) From: %s", fieldName, refObjTypeStr, d.getUUID(refObj), objTypeStr)
				d.Evaluate(refObj, refObjTypeStr, objTypeStr)
			}
		}
	}
}

