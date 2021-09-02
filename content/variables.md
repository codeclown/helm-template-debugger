## Wrap variable value in quotes

```yaml
# template.yaml
name: {{ .Values.storageClassName | quote }}
```

## Require default value

```yaml
# template.yaml
storageClassName: {{ .Values.storageClassName | required ".storageClassName must be set" }}
```
