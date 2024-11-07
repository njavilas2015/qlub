package config

const DefaultNginxTemplate = `worker_processes auto;
error_log /var/log/nginx/error.log warn;
pid /var/run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    keepalive_timeout  65;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;

    include /etc/nginx/conf.d/*.conf;
}
`

const NginxTemplate = `# ------------------------------------------------------------
# {{ .Domain }}
# ------------------------------------------------------------

{{- range $location := .Location }}

upstream {{ $location.Alias }}_upstream {
    {{- range $instances := $location.Instances }}
    server {{ $instances }}:{{ $location.Port }};
    {{- end }}
}
{{- end }}

server {
    {{- if .Ssl }}
    listen 443 ssl;
    ssl_certificate {{ .SslCert }};
    ssl_certificate_key {{ .SslCertKey }};
    {{- else }}
    listen 80;
    {{- end }}

    server_name {{ .Domain }};

    {{- range $location := .Location }}

    location {{ $location.Path }} {
        {{- if $location.Ssl }}
        proxy_pass https://{{ $location.Alias }}_upstream;
        {{- else }}
        proxy_pass http://{{ $location.Alias }}_upstream;
        {{- end }}
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection $http_connection;
    }
    {{- end }}
}
`
