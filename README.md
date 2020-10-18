# Notes-App

This is a little pet-project aimed to try out new technologies.
Notes-App should allow users to create notes, view them, delete and update, as well as keep the notes private.

### About the stack:
* monolith backend on Go
* REST API based on Gin
* GORM as ORM
* SQLite database

### About dev plans:
* SPA Frontend [WIP]
* migrate from sqlite to postgres
* migrate from GORM
* Makefile
* tests
* add prometheus metrics
* dockerize

## Building

Make sure you have Go compiler version 1.14.
Clone the repository, enter the folder and run command:

	$ go build -o bin/server cmd/main.go
	$ ./bin/server

## Configuration

All configuration variables are set in .env file.
