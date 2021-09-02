## Fail on purpose

```yaml
{{ if eq .Values.storageClassName "foobar1" }}
  # ...
{{ else if eq .Values.storageClassName "foobar2" }}
  # ...
{{ else }}
  {{ fail ".storageClassName is not recognized" }}
{{ end }}
```
