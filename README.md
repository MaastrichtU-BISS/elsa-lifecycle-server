# ELSA lifecycle tool server

## How to run?

Install Go on your system https://go.dev/doc/install

### Windows?
Also install mingw-w64 GCC through https://www.msys2.org/. GCC is used to run C code.

Enable C code to run using Go, since sqlite relies on it (see [go.mod](./go.mod))

```bash
go env -w CGO_ENABLED=1
```


To run the project locally in development mode:

```bash
go run .
```

### Seeding the database

By default, the database is not seeded. To seed the database (clear and insert demo data), set the environment variable `SEED=TRUE` when starting the server:

```bash
SEED=TRUE go run .
```

You can also use this with a custom database path:

```bash
SEED=TRUE DB_PATH="/tmp/my-elsa.db" go run .
```

**Warning:** Seeding will clear and repopulate the relevant tables. Only use this in development or when you want to reset the database.

## Class diagram
First draft of the class / model diagram can be found below. This is made using drawio where the xml metadata is encapsulated in the .png metadata. This can be edited in VS Code using the [Draw.io Integration plugin](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio).

![class diagram](./uml_model.drawio.png "UML model")

## Configuration: CORS_ALLOW_ORIGINS

You can override allowed CORS origins with the environment variable `CORS_ALLOW_ORIGINS`.
Provide a comma-separated list of origins. The server supports simple glob-style wildcards:

- `*` matches any sequence of characters
- `?` matches a single character

Examples:

- Allow only a specific origin:

	CORS_ALLOW_ORIGINS="http://example.com"

- Allow localhost with any port and a subdomain pattern:

	CORS_ALLOW_ORIGINS="http://localhost:*,https://*.example.org"

Notes:

- Patterns are converted to regular expressions internally. For safety, malformed patterns are ignored.
- If the variable is empty or unset, the server defaults to `http://localhost` and `http://localhost:*`.

## Configuration: DB_PATH

The server uses a SQLite database file. By default the project uses the file at `database/db/elsa.db`.

You can override the path to the SQLite database file with the environment variable `DB_PATH`.

Examples:

- Use a custom database file:

	DB_PATH="/tmp/my-elsa.db"

- Use the default (unset the variable):

	unset DB_PATH

Notes:

- If `DB_PATH` is not set or is empty, the server falls back to `database/db/elsa.db`.
- The path can be absolute or relative to the project working directory.

When running inside the provided Docker image the application runs from `/app`, so the container-friendly default path is:

	/app/database/db/elsa.db

Use an absolute `DB_PATH` when running the container to avoid ambiguity.

## Run as Docker image


### Seeding the database in Docker

To seed the database when running in Docker, set the `SEED` environment variable:

```bash
docker run -p 8080:8080 -e SEED=TRUE -v $(pwd)/db:/app/database/db ghcr.io/maastrichtu-biss/elsa-lifecycle-server
```

You can combine this with a custom database path:

```bash
docker run -p 8080:8080 -e SEED=TRUE -e DB_PATH="/db/elsa.sqlite" -v $(pwd)/db:/db ghcr.io/maastrichtu-biss/elsa-lifecycle-server
```

**Warning:** Seeding will clear and repopulate the relevant tables. Only use this in development or when you want to reset the database.

