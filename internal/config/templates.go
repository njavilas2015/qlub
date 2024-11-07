package config

const NginxTemplate = `events {

}

http {
    {{- range $subdomain := . }}

    {{- range $location := $subdomain.Location }}

    upstream {{ $location.Alias }}_upstream {
        {{- range $instances := $location.Instances }}
        server {{ $instances }}:{{ $location.Port }};
        {{- end }}
    }
    {{- end }}

    server {
        {{- if $subdomain.Ssl }}
        listen 443 ssl;
        ssl_certificate {{ $subdomain.SslCert }};
        ssl_certificate_key {{ $subdomain.SslCertKey }};
        {{- else }}
        listen 80;
        {{- end }}

        server_name {{ $subdomain.Domain }};

        {{- range $location := $subdomain.Location }}

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
    {{- end }}
}
`
