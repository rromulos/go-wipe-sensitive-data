# Sensitive Data Filtering API

This repository is dedicated to improving the learning experience in the Go programming language.

The main objective of this repository is to develop a user-friendly API that effectively filters out sensitive information from JSON objects. By utilizing this API, users can confidently send JSON data, knowing that any sensitive data will be securely removed.

The filtering process will identify and exclude sensitive data fields, ensuring the privacy and confidentiality of the information.

This repository serves as a Proof of Concept (POC), showcasing the capabilities and potential of learning enhancements and sensitive data filtering in Go.
## How to use

#### Starting the application

```
go run main.go
```

The application will provide an API running on port 7777

## API Documentation

#### Endpoint to filter sensitive data

```http
  POST /filter
```

| Parâmetro   | Descrição                           | Mandatory |
| :---------- | :---------------------------------- | :-------- |
| `fields` | Defines the fields to be removed | `N` |
| `nodes` | Defines the nodes to be removed | `N` |
| `data` | Defines the data | `Y` |

### Payload - Removing name and email fields. Removing address node
```
{
  "fields": "age,name",
  "nodes": ["address"],
  "data": {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30,
    "address": {
      "street": "123 Main St",
      "city": "New York"
    }
  }
}

```

**Note:** There should be no spaces between the fields entered.

#### Result

```
{
    "email": "johndoe@example.com"
}
```

### Payload - Removing address node

```
{
  "fields": "",
  "nodes": ["address"],
  "data": {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30,
    "address": {
      "street": "123 Main St",
      "city": "New York"
    }
  }
}
```

#### Result

```
{
    "age": 30,
    "email": "johndoe@example.com",
    "name": "John Doe"
}
```

### Payload - Removing email only
```
{
  "fields": "email",
  "nodes": [""],
  "data": {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30,
    "address": {
      "street": "123 Main St",
      "city": "New York"
    }
  }
}
```

#### Result

```
{
    "address": {
        "city": "New York",
        "street": "123 Main St"
    },
    "age": 30,
    "name": "John Doe"
}
```

### Payload - Removing all data
```
{
  "fields": "age,name,email",
  "nodes": ["address"],
  "data": {
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30,
    "address": {
      "street": "123 Main St",
      "city": "New York"
    }
  }
}
```

#### Result

```
{}
```
## Running the tests

To run the tests, run the following command in the project root folder

```bash
  go test ./tests/...   
```

#### Result

```
ok  	github.com/rromulos/go-clean-sensitive-data/tests
```

## Screenshots - Postman

![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)
