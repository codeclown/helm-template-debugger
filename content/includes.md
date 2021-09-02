## Include a template

```yaml
# _helpers.tpl
{{ define "your-project.annotations" }}
your-project/foobar: value1
{{ end }}

# template.yaml
annotations:
  {{ include "your-project.annotations" | toYaml | indent 2 }}
```

## Passing a scope to a template

```yaml
# _helpers.tpl
{{ define "your-project.annotations" }}
your-project/foobar: {{ .Release.Name }}
{{ end }}

# template.yaml
annotations:
  {{ include "your-project.annotations" . | toYaml | indent 2 }}
```
