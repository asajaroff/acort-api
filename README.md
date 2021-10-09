# URL Shortener 

## Description
This is a golang based API that receives one or more arguments and returns a shortened url.

## Design
* TODO: POST or GET?
* DB: Redis

## Data model

* Input
``` json
{
    "url": "https://example.com/pindongin?arg=v1,argv2",
    "name": "Pindongin",
    "api": "v1.beta"
}
```

* output

``` json
{
    "url": ""
}
```

## Setup

``` bash
GO11MODULE=on go get github.com/labstack/echo/v4
```
