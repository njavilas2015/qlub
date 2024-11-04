package config

import (
	"fmt"
	"os"
	"text/template"
)

const nginxTemplate = `http {
    upstream {{ .Name }} {
        {{- range .Instances }}
        server {{ . }}:{{ $.Port }};
        {{- end }}
    }

    server {
        {{- if .HTTPS }}
        listen 443 ssl;
        ssl_certificate {{ .SSLCert }};
        ssl_certificate_key {{ .SSLCertKey }};
        {{- else }}
        listen 80;
        {{- end }}

        server_name {{ .Name }};

        location / {
            proxy_pass http://{{ .Name }};
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
}
`

func GenerateNginxConfig(subdomain Subdomain) error {

	file, err := os.Create(fmt.Sprintf("%s_nginx.conf", subdomain.Name))

	if err != nil {
		return err
	}

	defer file.Close()

	raw_template, err := template.New("nginx").Parse(nginxTemplate)

	if err != nil {
		return err
	}

	return raw_template.Execute(file, subdomain)
}
