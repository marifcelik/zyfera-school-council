# Project Overview

This is a RESTful API project written in Go that manages student data, including their grades. The project has 2 main endpoints: `/create` and `/update/{stdNumber}`. 

The project uses standard `http` package for the HTTP server, a PostgreSQL database and GORM as the ORM.

## Running the Project

If you have Docker installed, you can simply run the project with the following command: 

```bash
docker-compose up -d
```

## Config

You can use environment variables to configure the application. The following environment variables are available:

| Field   | Type     | Description                                     | Default Value                                      |
| ------- | -------- | ----------------------------------------------- | -------------------------------------------------- |
| APP_ENV | `string` | The environment the application is running in   | "dev"                                              |
| DB_URL  | `string` | The URL of the PostgreSQL database              | "postgresql://postgres:pass@localhost:5432/school" |
| HOST    | `string` | The host address the application is running on  | "localhost"                                        |
| PORT    | `string` | The port number the application is listening on | "8080"                                             |

## Database Schema

The database schema is as follows:

`Grade`
```go
type Grade struct {
	ID        uint    `gorm:"primaryKey"`
	StudentID int     `gorm:"not null"`
	Code      string  `gorm:"not null"`
	Value     float64 `gorm:"not null"`
}
```

`Student`
```go
type Student struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    Surname   string `gorm:"not null"`
    StdNumber string `gorm:"unique;not null"`
    Grades   []Grade `gorm:"foreignKey:StudentID"`
}
```

## Endpoints

| Endpoint              | Method | Description                                  |
| --------------------- | ------ | -------------------------------------------- |
| `/health_check`       | GET    | Check the health of the application          |
| `/create`             | POST   | Create a new student with associated grades  |
| `/update/{stdNumber}` | PATCH  | Update an existing student's data and grades |

Example request for `/create`:
```jsonc
// Input
{
  "name": "arif",
  "surname": "celik",
  "stdNumber": "551",
  "grades": [
    {
      "code": "EV23",
      "value": 82
    }
  ]
}

// Output
{
  "data": {
    "name": "arif",
    "surname": "celik",
    "stdNumber": "551",
    "id": 18,
    "grades": [
      {
        "code": "EV23",
        "value": 82
      }
    ]
  },
  "status": "success"
}
```

Example request for `/update/{stdNumber}`:
```jsonc
// Input
{
  "name": "arif",
  "surname": "celik",
  "grades": [
    {
      "code": "EV23",
      "value": 82
    },
    {
      "code": "EV23",
      "value": 65
    },
    {
      "code": "MT122",
      "value": 56
    }
  ]
}

// Output
{
  "data": {
    "name": "arif",
    "surname": "celik",
    "stdNumber": "551",
    "id": 17,
    "grades": [
      {
        "code": "EV23",
        "value": 73.5
      },
      {
        "code": "MT122",
        "value": 56
      }
    ]
  },
  "status": "success"
}
```