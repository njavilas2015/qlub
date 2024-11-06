package config

const NginxTemplate = `http {
    # Define common proxy headers
    map $http_upgrade upgrade;
    map $connection connection;

    {{- range $subdomain := . }}

    {{- range $location := $subdomain.Location }}

    upstream {{$subdomain.Name}}_{{ $location.Path }}_upstream {
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

        server_name {{ $subdomain.Name }};

        {{- range $location := $subdomain.Location }}

        location {{ $location.Path }} {
            {{- if $subdomain.Ssl }}
            proxy_pass https://{{$subdomain.Name}}_{{ $location.Path }}_upstream;
            {{- else }}
            proxy_pass http://{{$subdomain.Name}}_{{ $location.Path }}_upstream;
            {{- end }}
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
        {{- end }}
    }
    {{- end }}
}
`
