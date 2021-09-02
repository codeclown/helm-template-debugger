## Indenting

```yaml
# values.yaml
environmentVariables:
  POSTGRES_DB: foobar-1
  POSTGRES_USER: postgres

# template.yaml
env:
  {{ .Values.environmentVariables | toYaml | indent 2 }}
```

## Indenting with a prepended newline

```yaml
# values.yaml
environmentVariables:
  POSTGRES_DB: foobar-1
  POSTGRES_USER: postgres

# template.yaml
env: {{ .Values.environmentVariables | toYaml | nindent 2 }}
```
