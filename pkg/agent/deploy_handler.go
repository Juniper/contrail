package agent

import (
	"fmt"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Possible operations of events.
const (
	OperationCreate = "CREATE"
	OperationUpdate = "UPDATE"
	OperationDelete = "DELETE"
)

func handleDeploy(e *services.Event, a *Agent) error {
	// TODO(dji): need to double check the key values in context, e.g. does context contain resources.schema_id
	context := pongo2.Context{
		"resource": e.GetResource().ToMap(),
		"action":   e.Operation(),
		"input":    a.config,
	}
	context["res_type"] = basemodels.KindToSchemaID(e.Kind())
	context["dir"] = fmt.Sprintf("/var/tmp/%s/config/%s", context["res_type"], e.GetUUID())

	switch e.Operation() {
	case OperationCreate:
		return deployCreate(context)
	case OperationUpdate:
		return deployUpdate(context)
	case OperationDelete:
		return deployDelete(context)
	}

	return nil
}

func deployCreate(context pongo2.Context) error {
	logrus.Debug(fmt.Sprintf("AGENT creating %s", context["dir"]))
	output, err := commandHandler(fmt.Sprintf("mkdir -p %s", context["dir"]))
	if err != nil {
		return errors.Wrapf(err, "AGENT create %s failed", context["dir"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT created %s", context["dir"]))

	err = templateHandler(
		"/etc/contrail/contrail-deploy-config.tmpl",
		fmt.Sprintf("%s/contrail-deploy-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(
		fmt.Sprintf(
			"AGENT creating %s(contrailgo deploy -c %s/contrail-deploy-config.yml)",
			context["res_type"],
			context["dir"],
		),
	)

	output, err = commandHandler(
		fmt.Sprintf("contrailgo deploy -c %s/contrail-deploy-config.yml", context["dir"]),
	)
	if err != nil {
		return errors.Wrapf(err, "AGENT %s create failed", context["res_type"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT %s create complete", context["res_type"]))

	return nil
}

func deployUpdate(context pongo2.Context) error {
	err := templateHandler(
		"/etc/contrail/contrail-deploy-config.tmpl",
		fmt.Sprintf("%s/contrail-deploy-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(
		fmt.Sprintf(
			"AGENT updating %s(contrailgo deploy -c %s/contrail-deploy-config.yml)",
			context["res_type"],
			context["dir"],
		),
	)
	output, err := commandHandler(
		fmt.Sprintf("contrailgo deploy -c %s/contrail-deploy-config.yml", context["dir"]),
	)
	if err != nil {
		return errors.Wrapf(err, "AGENT %s update failed", context["res_type"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT %s update complete", context["res_type"]))

	return nil
}

func deployDelete(context pongo2.Context) error {
	err := templateHandler(
		"/etc/contrail/contrail-deploy-config.tmpl",
		fmt.Sprintf("%s/contrail-deploy-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(
		fmt.Sprintf(
			"AGENT deleting %s(contrailgo deploy -c %s/contrail-deploy-config.yml)",
			context["res_type"],
			context["dir"],
		),
	)
	output, err := commandHandler(
		fmt.Sprintf("contrailgo deploy -c %s/contrail-deploy-config.yml", context["dir"]),
	)
	if err != nil {
		return errors.Wrapf(err, "AGENT %s delete failed", context["res_type"])
	}
	logrus.Debug(output)

	logrus.Debug(fmt.Sprintf("AGENT deleting %s", context["dir"]))
	output, err = commandHandler(fmt.Sprintf("rm -rf  %s", context["dir"]))
	if err != nil {
		return errors.Wrapf(err, "AGENT delete %s failed", context["dir"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT deleted %s", context["dir"]))

	return nil
}
