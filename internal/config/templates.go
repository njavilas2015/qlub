package config

const NginxTemplate = `http {
    # Define common proxy headers
    map $http_upgrade upgrade;
    map $connection connection;

    {{- range $subdomain := . }}

    upstream {{ $subdomain.Name }} {
        {{- range $subdomain.Instances }}
        server {{ . }}:{{ $subdomain.Port }};
        {{- end }}
    }

    server {
        {{- if $subdomain.HTTPS }}
        listen 443 ssl;
        ssl_certificate {{ $subdomain.SSLCert }};
        ssl_certificate_key {{ $subdomain.SSLCertKey }};
        {{- else }}
        listen 80;
        {{- end }}

        server_name {{ $subdomain.Name }};

        location / {
            proxy_pass http://{{ $subdomain.Name }};
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
    {{- end }}
}
`
