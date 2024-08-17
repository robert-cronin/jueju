{{- define "namespace.validateLabels" -}}
{{- if .Values.istio.enabled -}}
    {{- $namespace := lookup "v1" "Namespace" "" .Release.Namespace }}
    {{- if not $namespace }}
        {{- fail (printf "Namespace %s does not exist" .Release.Namespace) }}
    {{- end }}
    {{- if not (hasKey $namespace.metadata.labels "istio-injection") }}
        {{- fail (printf "Namespace %s is missing required label: istio-injection" .Release.Namespace) }}
    {{- end }}
{{- end }}
{{- end }}
