package services

import (
	"net/http"
	"net/url"

	"github.com/Juniper/asf/pkg/models"
	"github.com/pkg/errors"
)

func ApplyHref(r models.Object, req *http.Request) error {
	if err := applyHrefOnResource(r, req); err != nil {
		return err
	}
	for _, ref := range r.GetReferences() {
		if err := applyHrefOnReference(r, &refToHrefAdapter{ref}, req); err != nil {
			return err
		}
	}
	for _, backRef := range r.GetBackReferences() {
		if err := applyHrefOnReference(r, backRef, req); err != nil {
			return err
		}
	}
	for _, child := range r.GetChildren() {
		if err := applyHrefOnReference(r, child, req); err != nil {
			return err
		}
	}
	return nil
}

type withHref interface {
	Kind() string
	GetUUID() string
	SetHref(string)
}

func applyHrefOnResource(o withHref, req *http.Request) error {
	objURL, err := url.Parse(RequestSchema(req) + req.Host + "/" + o.Kind() + "/")
	if err != nil {
		return errors.Wrap(err, "failed to parse connection url")
	}
	uuidURL, err := url.Parse(o.GetUUID())
	if err != nil {
		return errors.Wrap(err, "failed to parse uuid")
	}
	o.SetHref(objURL.ResolveReference(uuidURL).String())
	return nil
}

func applyHrefOnReference(from, to withHref, req *http.Request) error {
	refURL, err := url.Parse(RequestSchema(req) + req.Host + "/" + to.Kind() + "/")
	if err != nil {
		return errors.Wrapf(err, "failed to resolve '%s-%s' reference url", from.Kind(), to.Kind())
	}
	uuidURL, err := url.Parse(to.GetUUID())
	if err != nil {
		return errors.Wrapf(err, "failed to parse '%s-%s' reference url with uuid '%s'",
			from.Kind(), to.Kind(), to.GetUUID())
	}
	to.SetHref(refURL.ResolveReference(uuidURL).String())
	return nil
}

type refToHrefAdapter struct {
	models.Reference
}

func (r *refToHrefAdapter) Kind() string {
	return r.GetToKind()
}
