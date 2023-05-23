[![Go Reference](https://pkg.go.dev/badge/github.com/abu-lang/goabu.svg)](https://pkg.go.dev/github.com/bancodobrasil/featws-resolver-adapter-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/bancodobrasil/featws-resolver-adapter-go/blob/feat/swagger-ptbr/LICENSE)

# Featws Reslver Bridge [![About_en](https://github.com/yammadev/flag-icons/blob/master/png/US.png?raw=true)](https://github.com/bancodobrasil/featws-resolver-adapter-go/blob/develop/README-PTBR.md)

## How to run

In order to run this project, you need to have certain prerequisites set up on your machine. These prerequisites include:

 - [Golang](https://go.dev/doc/install)
 - [Swaggo](https://github.com/swaggo/swag/blob/master/README_pt.md#come%C3%A7ando)

To run the project, follow these steps:

- Open the terminal in the project directory and run the command `go mod tidy` to ensure that all required dependencies are installed.

- Then, run the command `swag init` to initialize Swagger and generate the necessary API documentation.

- Finally, run the command `make run` to start the project.

The project will run on `localhost:9000/`. To access the Swagger documentation [click here](http://localhost:9000/swagger/index.html#/).

By following these steps, the project will be up and running, and you will be able to access the API documentation through Swagger.