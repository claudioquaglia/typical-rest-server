<!-- Autogenerate by Typical-Go. DO NOT EDIT.--> 

# Typical-RESTful-Server

Example of typical and scalable RESTful API Server for Go

## Getting Started

This is intruction to start working with the project:
1. Install go
2. Clone the project

## Usage

There is no specific requirement to run the application. 

## Development Tool

Use `./typicalw` to execute development task

## Make a release

Use `./typicalw release -github=[TOKEN]` to make the release.t r You can found the release in `release` folder or github release page

## Configurations

Application configuration

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|	
|APP_ADDRESS|String|:8089|true||	

Postgres configuration

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|	
|PG_DBNAME|String|typical-rest-server|true||	
|PG_USER|String|default_user|true||	
|PG_PASSWORD|String|default_password|true||	
|PG_HOST|String|localhost|||	
|PG_PORT|Integer|5432|||	
|PG_MIGRATIONSRC|String|scripts/migration|||	


