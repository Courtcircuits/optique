<h1 align="center"><b>(⌐■_■)</b><h1/>

<p align="center"><b>optique</b> : Go tooling for microservices (opinionated)</p>

## Why using optique ?

**optique** is the ShadCN for microservices. It is a tool that helps you to create a new microservice with a clean architecture (hexagonal) and an api-first approach.

The power of **optique** is from the `module` system. It allows you to create a boilerplate of a port (in the hexagonal meaning) that you would add to your repository folder.
E.g., optique doesn't come with a `database` module, by default, but you can create one by calling 

```bash
optique generate pg-database
```

This will create a new folder in your `infrastructure` folder with the name of the module you chose.And the cli will suggests how you can plug the database to your project.

Finally, projects created with optique **are not vendor-locked**. This means that you don't need optique to maintain your project.

## Project structure

```bash
.
├── config
│   └── config.go
├── config.json
├── core
├── docker-compose.yml
├── docs
├── http
│   ├── handler.go
│   ├── hello.go
│   ├── server.go
│   └── validator.go
├── infrastructure
├── justfile
├── main.go
├── README.md

```

- Config folder : Contains the configuration of the project to respect the [12 factor principles](https://12factor.net)
- Core folder : Contains the core of the project, the business logic. It would be the service layer of an hexagonal architecture
- Http folder : Contains the http routes of the project. It would be the port layer of an hexagonal architecture.
- Infrastructure folder : Contains the infrastructure of the project. It would be the data layer of an hexagonal architecture/repositories.
- Docs folder : Contains the documentation of the project. You are advised to use [swaggo](https://github.com/swaggo/swag) to generate the documentation based on the comments in your code.
- Justfile : Contains the commands to run the project. We are not using `optique` cli to run the project, because I don't want your project to be vendor-locked with optique.
- `main.go` : Contains the main function of the project.
- `README.md` : Please document your project here. At least how to run it.

## Getting Started

```bash
go install github.com/Courtcircuits/optique/cli@latest
optique init my-project
```

This will create a new project in the `my-project` folder.

Then, to run the hello world project, you can run the following command:

```bash
air
```
