## Loop over an array of values

```yaml
# values.yaml
volumeIds:
- volume-1
- volume-2

# template.yaml
volumes:
  {{ range $volumeId := .Values.volumeIds }}
  - volumeName: {{ $volumeId }}
  {{ end }}
```

## Loop over an array of values, with keys

```yaml
# values.yaml
configuration:
  POSTGRES_DB: foobar-1
  POSTGRES_USER: postgres

# template.yaml
volumes:
  {{ range $key, $value := .Values.configuration }}
  - {{ $key }}: {{ $value }}
  {{ end }}
```

## Loop without assigning to a variable

```yaml
# values.yaml
volumeIds:
- volume-1
- volume-2

# template.yaml
volumes:
  {{ range .Values.volumeIds }}
  - volumeName: {{ . }}
  {{ end }}
```
