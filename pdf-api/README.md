# PDF Manipulation API

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
govendor sync
```

## Run
Build then launch server on port 8080 :
```
go build
./pdf-api server 8080
```

Then make a `POST` request on `/documents/:fileName/editeur` with actions to perform.

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

A new PDF named with current timestamp gets generated on disk.

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
 - Add the ability to upload a document
 - Add the ability to download a document
 - Dockerize
 - CircleCI ?
