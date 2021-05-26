```bash
echo '

### VALUES ###

applicationType: deployment
applicationId: "foobar applicationId"
image: "foobar image"
memoryMb: 100
milliCpu: 100
environmentVariables:
- {}
volumeMounts:
- volumeId: foobar-volume-123
  mountPath: /var/data


### TEMPLATE ###

---
apiVersion: apps/v1
{{ if eq .Values.applicationType "deployment" -}}
kind: Deployment
{{ else if eq .Values.applicationType "statefulSet" -}}
kind: StatefulSet
{{ end -}}
metadata:
  name: {{ .Values.applicationId | quote }}
  labels:
    applicationId: {{ .Values.applicationId | quote }}
spec:
  replicas: 1
  selector:
    matchLabels:
      applicationId: {{ .Values.applicationId | quote }}
  template:
    metadata:
      labels:
        applicationId: {{ .Values.applicationId | quote }}
    spec:
      containers:
      - name: {{ .Values.applicationId | quote }}
        image: {{ .Values.image | quote }}
        resources:
          limits:
            memory: "{{ .Values.memoryMb }}Mi"
            cpu: "{{ .Values.milliCpu }}m"
          requests:
            memory: "{{ .Values.memoryMb }}Mi"
            cpu: "{{ .Values.milliCpu }}m"
{{- if not (empty .Values.environmentVariables) }}
        envFrom:
        - configMapRef:
            name: {{ .Values.applicationId | quote }}
{{- end }}
{{- if not (empty .Values.volumeMounts) }}
        volumeMounts:
{{- range $volumeMount := .Values.volumeMounts }}
        - mountPath: {{ $volumeMount.mountPath | quote }}
          name: {{ $volumeMount.volumeId | quote }}
{{- end }}
{{- end }}
{{- if not (empty .Values.volumeMounts) }}
      volumes:
{{- range $volumeMount := .Values.volumeMounts }}
      - name: {{ $volumeMount.volumeId | quote }}
        persistentVolumeClaim:
          claimName: {{ $volumeMount.volumeId | quote }}
{{- end }}
{{- end }}


' | curl -X POST --data-binary @- http://localhost:8080/generate
```
