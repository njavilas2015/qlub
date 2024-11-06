# qlub

`qlub` es una herramienta diseñada para gestionar subdominios y configuraciones de Nginx. Proporciona una interfaz fácil de usar para configurar y administrar subdominios, ayudando a simplificar la gestión de servidores web.

## Índice

1. [Características](#características)
2. [Instalación](#instalación)
3. [Uso](#uso)
   - [Ejemplos](#ejemplos)
4. [Configuración](#configuración)
5. [Contribución](#contribución)
6. [Licencia](#licencia)
7. [Contacto](#contacto)

## Características

- **Gestión de Subdominios**: Permite agregar, eliminar y modificar subdominios fácilmente.
- **Configuración de Nginx**: Genera configuraciones de Nginx automáticamente basadas en los subdominios definidos.
- **Compatibilidad con SSL**: Soporta la configuración de certificados SSL para una navegación segura.

## Instalación

Para instalar `qlub`, puedes descargar el binario desde el siguiente enlace:

- [Descargar qlub](https://github.com/njavilas2015/qlub/releases/download/v1.0.1/qlub)

```bash
wget https://github.com/njavilas2015/qlub/releases/download/v1.0.1/qlub
```

Después de descargar, asegúrate de que el binario sea ejecutable y mueve el archivo a un directorio en tu `PATH`:

```bash
chmod +x qlub
sudo mv qlub /usr/local/bin/
```

## Uso
Para usar `qlub`, puedes ejecutar el siguiente comando en la terminal:

```bash
qlub --config <ruta_al_archivo_json> --watch #para detectar cambios y actualizar config

```

## Configuración
El archivo de configuración JSON debe contener una lista de subdominios. Cada subdominio puede tener los siguientes campos:

- name: Nombre del subdominio.
- location: Lista de instancias del servicio.
- ssl: (opcional) Si el subdominio debe usar HTTPS.
- ssl_cert: (opcional) Ruta al certificado SSL.
- ssl_cert_key: (opcional) Ruta a la clave del certificado SSL.


```json
[
   {
        "name": "domain.com",
        "location": [
            {
                "path": "/",
                "port": "443",
                "instances": [
                    "site_docker_1",
                    "site_server_2"
                ]
            },
            {
                "path": "/api",
                "port": "8000",
                "instances": [
                    "api_docker_1",
                    "api_server_2"
                ]
            }
        ],
        "ssl": true,
        "ssl_cert": "/etc/letsencrypt/fullchain.pem",
        "ssl_cert_key": "/etc/letsencrypt/privkey.pem"
    }
]
```

## Docker Compose 
Puedes descargar la imagen lista para trabajar `docker pull njavilas/qlub:latest`

```yml
services:
  qlub:
    image: njavilas/qlub:latest
    volumes:
      - ./subdomains.json:/app/subdomains.json
      - ./nginx.conf:/app/nginx.conf
```

## Build

go build -o qlub

## Contacto
Si tienes alguna pregunta o necesitas soporte, no dudes en contactarme:

Nombre: Javier Avila
Email: [njavilas2015@gmail.com]
GitHub: njavilas2015