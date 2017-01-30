# PDF Manipulation API

## Run in docker
```
docker run -d -p 8080:8080 zenika/pdf-api server 8080
```

## Installation
Link the project in your `$GOPATH` :
```
mkdir -p $GOPATH/src/github.com/Zenika
ln -s $(pwd) $GOPATH/src/github.com/Zenika/pdf-api
```

As a prerequisite for managing dependencies, install `govendor` :
```
go get -u github.com/kardianos/govendor
```

Fetch go dependencies :
```
cd $GOPATH/src/github.com/Zenika/pdf-api
govendor sync
```

## Build
```
cd $GOPATH/src/github.com/Zenika/pdf-api
go build
```

## Run
Build then launch server on port 8080 :
```
./pdf-api server 8080
```

PDF files from `default-documents/` folder will be automatically copied to `documents/` folder at startup, and available to download and modify.

Upload new files with `PUT` requests on `/document/:file/content` (`file` without `.pdf` suffix) with content type `application/pdf`.

Download a file with a `GET` request on `/document/:file/content` (`file` without `.pdf` suffix).

Modify a file with a `POST` request on `/document/:file/editeur` (`file` without `.pdf` suffix) with actions to perform.

Example :
```
[
  {
    "action": "LEFT_ROTATE_PAGE",
    "page": 1
  },
  {
    "action": "RIGHT_ROTATE_PAGE",
    "page": 2
  },
  {
    "action": "LEFT_ROTATE_PAGE",
    "page": 3
  },
  {
    "action": "LEFT_ROTATE_PAGE",
    "page": 3
  },
  {
    "action": "DELETE_PAGE",
    "page": 4
  },
  {
    "action": "MOVE_PAGE",
    "page": 53,
    "target": 1
  }
]
```

WARNING ! The PDF file gets modified in place without backup !

## Run in CLI mode
Download a PDF to manipulate :
```
curl http://www.syntec.fr/fichiers/Annexes/20130719184036_Convention_Syntec_Annexe_06.pdf -o syntec.pdf
```

### Examples
Left rotate a page :
```
./pdf-api left-rotate-page syntec.pdf 1 test.pdf
```

Delete a page :
```
./pdf-api delete-page syntec.pdf 2 test.pdf
```

Move a page :
```
./pdf-api move-page syntec.pdf 54 1 test.pdf
```

## TODO
 - Add logs
 - CircleCI -> WIP
