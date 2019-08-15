# Go HTTP File Server
Share local file system by HTTP.

![Go HTTP File Server pages](doc/ghfs.gif)

## Compile
```bash
cd src
go build main.go
```

If default html template file under `src/tpl` changed, run
```bash
cd src
make mktpl
```
to re-embed templates into go files. Then compile the project again.

## Usage
```
-root <directory>
    Root directory of the server.
    Defaults to current working directory.

-listen <[ip]:port>
    Optional IP and port the server listens on. e.g. ":80".

-cert <file>
    Specify TLS certificate file.
    If both "cert" and "key" are specified, the server serves for HTTPS protocol.

-key <file>
    Specify key file of TLS certificate.
    If both "cert" and "key" are specified, the server serves for HTTPS protocol.

-template <file>
    Use a custom template file for rendering pages, instead of builtin template.
```
