{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified dependency name for the db.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "postgres.fullname" -}}
{{- printf "%s-%s" .Release.Name "postgresql" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified dependency name for the etcd.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "etcdclient.fullname" -}}
{{- printf "%s-%s" .Release.Name "etcdclient" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- define "etcd0.fullname" -}}
{{- printf "%s-%s" .Release.Name "etcd0" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- define "etcd1.fullname" -}}
{{- printf "%s-%s" .Release.Name "etcd1" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- define "etcd2.fullname" -}}
{{- printf "%s-%s" .Release.Name "etcd2" | trunc 63 | trimSuffix "-" -}}
{{- end -}}