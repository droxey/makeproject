# Changelog

## `v0.0.0`

- Ability to parse and render any file in  subfolders of the `templates` directory and inject `{{.ProjectType}}` and `{{.Name}}` variables.
- Strip `.tmpl` extension from rendered templates.
- Save parsed templates as original filename and save the project and new configuration to the `output` directory.
