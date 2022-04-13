# Leaf Migration

This directory contains utilities support for [Leaf migration codegen](https://github.com/enricodg/leaf-migration-codegen).

## Basic Usage

New
```shell
$ go run main.go new --types <mysql|mongo|postgre> --name <migration-name>
```
Migrate
```shell
$ go run main.go migrate [--types <mysql,mongo,postgre>] [--version <VERSION>] [--verbose] [--specific]
```
Rollback
```shell
$ go run main.go rollback --version <VERSION> [--types <mysql,mongo,postgre>] [--verbose] [--specific]
```
Check
```shell
$ go run main.go check [--types <mysql,mongo,postgre>] [--version <VERSION>]
```

## Example

Project example can be found [here](https://github.com/enricodg/leaf-migration-example).