package render

import "text/template"

const commitsTemplate = `## {{ .RepoName }}
{{- range .Commits }}
- {{ . }}
{{- end }}`

var render = template.Must(template.New("commits").Parse(commitsTemplate))
