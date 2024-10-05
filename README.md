
<p align="center">
  <img src="./static/beschrapi-logo.png" alt="Beschrapi Logo" width="200">
</p>

<h1 align="center">Beschrapi</h1>

<p align="center">
  <strong>
  Beschrapi is a tool designed for building APIs using YAML files. 
  </strong>
</p>


Beschrapi (bəˈʃraːpi), derived from the German word "beschreiben" (meaning "to describe") and the English acronym "API" (Application Programming Interface), is a tool designed for building APIs using YAML files. It offers a simple and efficient way to create APIs with minimal code.

> [!NOTE]  
> This is only the example repository for the Beschrapi tool. Currently Beschrapi is under development and not yet released.
>
> But you can still use Beschrapi to build your APIs. Just clone this repository to get started.

# Get started

## Requirements

- Docker is all you need :)

## Installation

1. Clone this repository
    ```bash
    git clone https://github.com/beschrapi/example.git
    ```

2. Start the Beschrapi server
    ```bash
    docker-compose up
    ```

3. Open your browser and explore the following routes:
   - [http://localhost:8080/version](http://localhost:8080/version)
   - [http://localhost:8080/products](http://localhost:8080/products)
   - [http://localhost:8080/products/1](http://localhost:8080/products/1)
   - [http://localhost:8080/file](http://localhost:8080/file)


4. You can also try out POST requests using Postman or any other API testing tool.
   - [http://localhost:8080/products/create](http://localhost:8080/products/create) (POST)
        - Body: `{"name": "Product 1", "price": 100}`
   - [http://localhost:8080/code](http://localhost:8080/code) (POST)
        - Body: `{"a": 1, "b": 2}`

# IntelliSense

Beschrapi offers IntelliSense support for VSCode.

> [!WARNING]
> IntelliSense is not fully working correctly yet. But you can still use it to get some help. (It does a pretty good job at the moment)

Just include the following line at the beginning of your YAML file:
```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/beschrapi/schema/refs/heads/master/schema.json config.bscrp.yaml
```

# Configuration

Beschrapi uses a simple YAML configuration file to define the API. 

The configuration file consists of the following sections:

- `server`: Server configuration like port and host.
- `database`: Database configuration 
- `paths`: API endpoints and their configuration. 
- `auth`: Defines the authentication to protect routes. (Not yet implemented)

For example look at the [config.bscrp.yml](./config.bscrp.yml) file.

# Paths

Beschrapi uses the `paths` section in the configuration file to define the API endpoints. Currently, Beschrapi supports the following HTTP methods:

- `GET`
- `POST`

More methods will be added in the future.

## Define a path

To define a path, you need to specify the path and the method. 

```yaml
paths:
  /users: # Your path URL
    get:  # Method to use
        json: | # The action for response
            [
                {
                    "id": 1,
                    "name": "John Doe"
                },
                {
                    "id": 2,
                    "name": "Jane Doe"
                }
            ]
```

## Route parameters

Beschrapi supports route parameters. You can define a route parameter by using curly braces `{variable}`

```yaml
paths:
  /users/{id}:
    get:
        json: |
            {
                "id": {id},
                "name": "John Doe"
            }
```

You can also define multiple route parameters.

```yaml
paths:
  /users/{id}/posts/{post_id}:
    get:
        json: |
            {
                "id": {id},
                "post_id": {post_id},
                "title": "Hello World",
                "content": "This is a post"
            }
```

## SQL

Beschrapi can execute SQL queries to fetch data from the database. You can use the `sql` section to define the SQL query.

```yaml
paths:
  /users:
    get:
        sql: SELECT * FROM users
```

## Code

Beschrapi offers the possibility to execute code in the response. You can use the `code` section to define the code to execute.

```yaml
paths:
  /users:
    get:
        code: code_file_name.go
```

The code file should contain the following:

```go
package main

import (
	"encoding/json"
	"net/http"
	"github.com/beschrapi/pkg"
)

func Run(w http.ResponseWriter, r *http.Request, app *pkg.App) {
}
```

## File

Beschrapi can also return a file. You can use the `file` section to define the file to return.

```yaml
paths:
  /users:
    get:
        file: file_name.txt
```


# Issues

If you have any issues or questions, please feel free to open an issue on this or the pkg repository.




