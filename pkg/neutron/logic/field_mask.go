package logic

import "github.com/Juniper/asf/pkg/models"

func buildDataResourcePath(fields ...string) string {
	fieldPath := models.JoinPath(fields...)
	return models.JoinPath("data.resource", fieldPath)
}

func buildContextPath(fields ...string) string {
	fieldPath := models.JoinPath(fields...)
	return models.JoinPath("context", fieldPath)
}
