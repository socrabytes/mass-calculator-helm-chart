{{/*
Define a template that returns the full name of the release.
*/}}
{{- define "mass-calculator.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" | replace "." "-" -}}
{{- end -}}

{{/*
Define a template that returns the labels for the release.
*/}}
{{- define "mass-calculator.labels" -}}
app.kubernetes.io/name: {{ include "mass-calculator.name" . }}
helm.sh/chart: {{ include "mass-calculator.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Define a template that returns the selector labels for the release.
*/}}
{{- define "mass-calculator.selectorLabels" -}}
app.kubernetes.io/name: {{ include "mass-calculator.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Define a template that returns the chart name.
*/}}
{{- define "mass-calculator.name" -}}
{{ .Chart.Name | replace "." "-" }}
{{- end -}}

{{/*
Define a template that returns the chart name and version.
*/}}
{{- define "mass-calculator.chart" -}}
{{ .Chart.Name }}-{{ .Chart.Version | replace "." "-" }}
{{- end -}}
