# Code generation

This directory contains files used to generate the [ccloud Go API Client](../../docs/README.md) with the [OpenAPI Generator](https://openapi-generator.tech).

## Overview
### OpenAPI specification
[openapi.json](./openapi.json) is the OAS 3.0 description for the Cockroach DB Cloud API. The most recent description can be downloaded by running `make fetch-latest-spec`.

### Templates
Our modified [mustache](http://mustache.github.io/mustache.5.html) template files customize the generated code for the [ccloud Go API Client](../../docs/README.md). 

The customized templates __do not support__ the following:
- Generating interfaces.
- [nullable](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.1.md#schemaNullable) fields.
- Operation grouping with tags. All API operation code is included in [pkg/client/service.go](../../pkg/client/service.go).
- More than one [server object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.1.md#server-object).
- [Server variables](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.1.md#server-variable-object).

To get the default templates:
`openapi-generator author template -g go -o <output directory>`

[OpenAPI Generator doc: Templating ](https://openapi-generator.tech/docs/templating/)

### Config file
[./config.yaml] sets generic and Go generator options.

## Use
Install `openapi-generator`. 
- `brew install openapi-generator`
- [OpenAPI Generator doc: Installation](https://openapi-generator.tech/docs/installation/)

Do not delete [.openapi-generator-ignore](../openapi-generator/.openapi-generator-ignore)! It prevents certain files from being generated. 

To regenerate the ccloud Go API Client code using the files in this directory, use `make generate-openapi-client`. 
The `openapi-generator` cli can also be used like:
`openapi-generator generate -g go -i <OpenAPI specification file> -o <output directory> -c <config file> -t <directory containing mustache templates>`

### Debugging
Use [debugging flags](https://openapi-generator.tech/docs/debugging/#templates) to inspect the template variables.

To see all the JSON models passed to model and operation (api), include `--global-property debugModels,debugOperations`.