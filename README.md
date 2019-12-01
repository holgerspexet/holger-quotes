
# Build and Run
Run the project locally
```
go run main.go
```

Build the project
```
go generate ./...               // this bundles static files and templates
go build -tags release main.go
```


# Configuration

| Variable                   | default       | comments                                                         |
|----------------------------|---------------|------------------------------------------------------------------|
| HOLGER_QUOTES_PORT         | 9000          | The port of the http server                                      |
| HOLGER_QUOTES_STORAGE_TYPE | sqlite        | Type of storage (memory, sqlite)                                 |
| HOLGER_QUOTES_SQLIGHT_PATH | ./sqlite3.sql | Path to the sqlite3 storage file                                 |
| HOLGER_QUOTES_HOSTING      | /             | The base URI at which the site is hosted                         |


# Dependencies
This package uses go-sqlite3 which need gcc in order to complie.
Se below how to install gcc on different platforms

## Windows
1) Download gcc from http://tdm-gcc.tdragon.net/download
2) Install the downloaded file