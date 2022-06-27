# gofiber-member-api

## TL;DR;

- Framework: Gofiber
- Database: SQLite
- ORM: gorm

## Configurations

| Name    | Description    | Remarks                                        |
| ------- | -------------- | ---------------------------------------------- |
| PROFILE | Profile        | - `dev`(default): Development </br> - `prd`: Production |
| PORT    | Listening Port | default: 8080                                  |

## API

- GET /api/v1/members
- GET /api/v1/members/:id
- POST /api/v1/members/:id
- PUT /api/v1/members/:id
- DELETE /api/v1/members/:id

## Tables

- members
- favorites

## Dockerize

```bash
  docker build -t member-api:1.0.0 .
  docker run -d --publish "8080:8080" --env PORT="8080" --name member-api member-api:1.0.0
```