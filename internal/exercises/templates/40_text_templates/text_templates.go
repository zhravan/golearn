package texttemplates

import (
	"bytes"
	"text/template"
)

// RenderTemplate takes a template string and data map,returns the rendered string or error.
func RenderTemplate(tmplStr string, data map[string]interface{}) (string, error) {
	tmpl, err := template.New("exercise").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	
	return buf.String(), nil
}

// FormatUserGreeting returns a greeting using text/template with users name
func FormatUserGreeting(name string) (string, error) {
	const greetingTmpl = "Hello, {{.Name}}!"
	data := map[string]interface{}{
		"Name": name,
	}
	return RenderTemplate(greetingTmpl, data)
}
