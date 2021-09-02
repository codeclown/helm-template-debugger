## If, if-else

```yaml
{{ if .Values.enablePersistence }}
  # ...
{{ else if .Values.enableFilesystem }}
  # ...
{{ else }}
  # ...
{{ end }}
```

## Equality: eq, ne

```yaml
{{ if eq .Values.environment "production" }}
  # ...
{{ end }}

{{ if ne .Values.environment "production" }}
  # ...
{{ end }}
```

## Combining: and, or

```yaml
{{ if and (eq .Values.environment "production") (eq .Values.host "minikube") }}
  # ...
{{ end }}

{{ if or (eq .Values.environment "production") (eq .Values.host "minikube") }}
  # ...
{{ end }}
```

## Negation: not

```yaml
{{ if not (eq .Values.environment "production") }}
  # ...
{{ end }}

{{ if eq .Values.environment "production" | not }}
  # ...
{{ end }}
```

## Comparing numbers: gt, gte, lt, lte

```yaml
# if more than 3 items
{{ if gt (len .Values.items) 3 }}
  # ...
{{ end }}

# if more than or equal to 3 items
{{ if gte (len .Values.items) 3 }}
  # ...
{{ end }}

# if less than 3 items
{{ if lt (len .Values.items) 3 }}
  # ...
{{ end }}

# if less than or equal to 3 items
{{ if lte (len .Values.items) 3 }}
  # ...
{{ end }}
```