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
	"reflect"
	"sync"

	"gopkg.in/oleiade/reflections.v1"

	log "github.com/sirupsen/logrus"
)

// NewDependencyProcessor - creates new instance
func NewDependencyProcessor(objCache *sync.Map) *DependencyProcessor {
	d := &DependencyProcessor{cache: objCache}
	d.Init()
	return d
}

// DependencyProcessor stores resources dependency
type DependencyProcessor struct {
	// A list of resources which the Dependency Tracker keeps
	// when it is invoked with a resource CRUD
	resources *sync.Map

	// Local Cache to maintained by Intent compiler.
	// Dependency Tracker  uses this cache to find dependent resources.
	cache *sync.Map
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
		if there {
			log.Infof("%s exists, not adding", d.getUUID(obj))
		}
		return !there
	}
	return true
}

// getCachedObject gets the cached object
func (d *DependencyProcessor) getCachedObject(objTypeStr, uuid string) interface{} {
	if objMap, ok := d.cache.Load(objTypeStr); ok {
		if obj, present := objMap.(*sync.Map).Load(uuid); present {
			return obj
		}
		return nil
	}
	return nil
}

// Evaluate - Evaluates object dependency based on the ReactionMap
func (d *DependencyProcessor) Evaluate(obj interface{}, objTypeStr, fromTypeStr string) { // nolint: gocyclo
	if _, ok := ReactionMap[objTypeStr]; !ok {
		return
	}

	log.Infof("Evaluating: Object: %s(%s) From: %s", objTypeStr, d.getUUID(obj), fromTypeStr)
	if !d.canAdd(objTypeStr, obj) {
		return
	}
	d.Add(objTypeStr, obj)

	for _, refObjTypeStr := range ReactionMap[objTypeStr][fromTypeStr] {
		fieldsToExtract := []string{refObjTypeStr + "Refs"}
		for _, fieldName := range fieldsToExtract {
			refObjTypeValues, err := reflections.GetField(obj, fieldName)
			if err != nil {
				// Refs dont exist, ignore
				continue
			}
			objValues := reflect.ValueOf(refObjTypeValues)
			for i := 0; i < objValues.Len(); i++ {
				interfaceObj := objValues.Index(i).Elem().Interface()
				uuid, _ := reflections.GetField(interfaceObj, "UUID")
				refObj := d.getCachedObject(refObjTypeStr, uuid.(string))
				if refObj == nil {
					continue
				}
				log.Infof("Evaluating: Object: %s %s(%s) From: %s", fieldName, refObjTypeStr, uuid.(string), objTypeStr)
				d.Evaluate(refObj, refObjTypeStr, objTypeStr)
			}
		}
		fieldsToExtract = []string{refObjTypeStr + "BackRefs"}
		for _, fieldName := range fieldsToExtract {
			refObjTypeValues, err := reflections.GetField(obj, fieldName)
			if err != nil {
				// BackRefs dont exist, ignore
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
