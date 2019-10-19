
# Configuration

| Variable                   | default       | comments                                                         |
|----------------------------|---------------|------------------------------------------------------------------|
| HOLGER_QUOTES_PORT         | 9000          | The port of the http server                                      |
| HOLGER_QUOTES_TEMPLATE_DIR | ./templates   | A path to the directory containing the templates                 |
| HOLGER_QUOTES_STATIC_DIR   | ./static      | A path to the directory containing static assets                 |
| HOLGER_QUOTES_STORAGE_TYPE | sqlite        | Type of storage (memory, sqlite)                                 |
| HOLGER_QUOTES_SQLIGHT_PATH | ./sqlite3.sql | Path to the sqlite3 storage file                                 |
| HOLGER_QUOTES_HOSTING      | /             | The base URI at which the site is hosted (must end with a slash) |



# Dependencies
This package uses go-sqlite3 which need gcc in order to complie.
Se below how to install gcc on different platforms

## Windows
1) Download gcc from http://tdm-gcc.tdragon.net/download
2) Install the downloaded file