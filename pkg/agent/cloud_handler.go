package agent

import (
	"fmt"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func handleCloud(e *services.Event, a *Agent) error {
	context := pongo2.Context{
		"resource": e.GetResource().ToMap(),
		"action":   e.Operation(),
		"input":    a.config,
	}
	context["src_dir"] = "/etc/contrail/contrail-cloud-config.tmpl"
	context["dir"] = fmt.Sprintf("/var/tmp/%s/config/%s", basemodels.KindToSchemaID(e.Kind()), e.GetUUID())

	switch e.Operation() {
	case OperationCreate:
		return cloudCreate(context)
	case OperationUpdate:
		return cloudUpdate(context)
	case OperationDelete:
		return cloudDelete(context)
	}

	return nil
}

func cloudCreate(context pongo2.Context) error {
	logrus.Debug(fmt.Sprintf("AGENT creating %s", context["dir"]))
	output, err := commandHandler(fmt.Sprintf("mkdir -p %s", context["dir"]))
	if err != nil {
		return errors.Wrapf(err, "AGENT create %s failed", context["dir"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT created %s", context["dir"]))

	err = templateHandler(
		context["src_dir"].(string),
		fmt.Sprintf("%s/contrail-cloud-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(
		fmt.Sprintf(
			"AGENT creating cloud (contrailgo cloud -c %s/contrail-cloud-config.yml)",
			context["dir"],
		),
	)
	output, err = commandHandler(
		fmt.Sprintf("contrailgo cloud -c %s/contrail-cloud-config.yml", context["dir"]),
	)
	if err != nil {
		return errors.Wrap(err, "AGENT cloud create failed")
	}
	logrus.Debug(output)
	logrus.Debug("AGENT cloud create complete")

	return nil
}

func cloudUpdate(context pongo2.Context) error {
	err := templateHandler(
		context["src_dir"].(string),
		fmt.Sprintf("%s/contrail-cloud-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(
		fmt.Sprintf(
			"AGENT updating cloud (contrailgo cloud -c %s/contrail-cloud-config.yml)",
			context["dir"],
		),
	)
	output, err := commandHandler(
		fmt.Sprintf("contrailgo cloud -c %s/contrail-cloud-config.yml", context["dir"]),
	)
	if err != nil {
		return errors.Wrap(err, "AGENT cloud update failed")
	}
	logrus.Debug(output)
	logrus.Debug("AGENT cloud update complete")

	return nil
}

func cloudDelete(context pongo2.Context) erro {
	err := templateHandler(
		context["src_dir"].(string),
		fmt.Sprintf("%s/contrail-cloud-config.yml", context["dir"]),
		context,
		a,
	)
	if err != nil {
		return errors.Wrap(err, "config file generation from template failed")
	}

	logrus.Debug(fmt.Sprintf("AGENT deleting %s", context["dir"]))
	output, err := commandHandler(
		fmt.Sprintf("rm -rf %s", context["dir"]),
	)
	if err != nil {
		return errors.Wrapf(err, "AGENT delete %s FAILED", context["dir"])
	}
	logrus.Debug(output)
	logrus.Debug(fmt.Sprintf("AGENT deleted %s", context["dir"]))

	return nil
}
