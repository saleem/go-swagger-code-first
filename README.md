# Swagger generation example

This is based largely on [Big K's blog](http://bigk.me/swagger/golang/go/2017/07/30/golang-and-openapi-how-to-create-documentation-for-your-rest-api-services/).

## Prerequisites

Install [go](https://golang.org/doc/install) and [go-swagger](https://github.com/go-swagger/go-swagger).

## Steps

Run this command to generate a `swagger.json` file:

`swagger generate spec -mo ./swagger.json`

This generated `swagger.json` file contains the documentation for the golang source files.

You can serve this file as a browsable web app by running this command:

`swagger serve swagger.json`


