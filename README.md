
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

# CI
Continues integration is run with github-actions.
Configurations can be found in .github/workflows/release.yml

## Releases
On each successfull push to the master branch a new release is automatically created

The version of this new release is one patch above the previously:
`2.5.7 => 2.5.8`

If, hoverver, the commit message contains `#major` or `#minor`, the corresponding version is bumped instead.
```
"Increase #major"       2.5.7 => 3.0.0
"Increase #minor"       2.5.7 => 2.6.0
```

## Atrifacts
The binary for each release can be found as an atrifact with the name holger-quotes. It is created with `GOOS=linux` `GOARCH=amd64`.

This link can be used to download the binary from the latest build
https://github.com/holgerspexet/holger-quotes/releases/latest/download/holger-quotes

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