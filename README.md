# ELSA lifecycle tool server

## How to run?

Install Go on your system https://go.dev/doc/install and Docker (for the local Postgres database).

Start Postgres (creates the `elsa` database, plus `elsa_test` for e2e tests):

```bash
docker compose up -d
```

Create a `.env` file from the example and set `JWT_SECRET` (the server won't start without it):

```bash
cp .env.example .env
```

To run the project locally in development mode:

```bash
go run .
```


# Seeding the database

The recommended way to seed the database (clear and insert demo data) is to use the dedicated seeder CLI:

```bash
go run cmd/seed/main.go
```

This will load all seed data from the `database/seeds/` directory (including JSON and JSON-LD files) and populate the database accordingly.

Flags:

- `-skip-users`: don't create the demo users from `users.json`. Use this in production, since the demo user has a publicly known password.
- `-force`: seed even if the database contains journals or registered users. Seeding refuses by default, because it deletes them.

You can also seed a different database by setting the `DATABASE_URL` environment variable:

```bash
DATABASE_URL="postgres://elsa:elsa@localhost:5432/elsa_test?sslmode=disable" go run cmd/seed/main.go
```

## Seed Files and Schemas

All seed data is located in the `database/seeds/` directory. The following files are used:

### users.json
- **Schema:**
	- `ID` (string, UUID)
	- `Email` (string)
	- `PasswordHash` (string)

### lifecycles.json
- **Schema:**
	- `Title` (string)
	- `Description` (string)
	- `General` (string, markdown)
	- `Introduction` (string, markdown)

### tools.json
- **Schema:**
	- `Title` (string)
	- `Description` (string)
	- `URL` (string)
	- `Cover` (string)
	- `Tags` (string)
	- `Type` (string)
	- `FormFile` (string, path to JSON-LD in `database/seeds/forms/`)

### reflections.json
The sections of a lifecycle's journal, shown in the order of this file.
- **Schema:**
	- `Title` (string, section label shown in the journal index)
	- `Context` (string, introductory paragraph)
	- `Description` (string, the question the user answers)
	- `Considerations` (string, JSON array of strings)
	- `FormFile` (string, path to JSON-LD in `database/seeds/forms/`)
	- `FurtherReflectionFormFile` (string, path to JSON-LD in `database/seeds/forms/`)
	- `LifecycleID` (integer, foreign key)

### reflection_answers.json
- **Schema:**
	- `Form` (string, JSON)
	- `BinaryEvaluation` (integer)
	- `ReflectionID` (integer, foreign key)
	- `UserID` (string, UUID)

### recommendations.json
- **Schema:**
	- `ReflectionID` (integer, foreign key: position of the section in `reflections.json`, starting at 1)
	- `ToolID` (integer, foreign key: position of the tool in `tools.json`, starting at 1)

### recommendation_answers.json
- **Schema:**
	- `Form` (string, JSON)
	- `File` (string, optional)
	- `RecommendationID` (integer, foreign key)
	- `UserID` (string, UUID)

### forms/ (directory)
- Contains referenced JSON-LD form files, e.g.:
	- `value_sensitive_design.json`
	- `reflections/generic_reflection_form.jsonld`
	- `reflections/generic_further_reflection_form.jsonld`

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

## Configuration: JWT_SECRET

Required. The key used to sign login tokens, at least 32 characters. The server refuses to start without it. Generate one with:

```bash
openssl rand -base64 48
```

Use a different secret per environment. Changing it logs out all users.

## Configuration: DATABASE_URL

The server uses a PostgreSQL database, configured with the `DATABASE_URL` environment variable (see [.env.example](./.env.example)).

Examples:

- URL format:

	DATABASE_URL="postgres://user:password@host:5432/elsa?sslmode=require"

- Key/value format:

	DATABASE_URL="host=localhost user=elsa password=elsa dbname=elsa port=5432 sslmode=disable"

Notes:

- If `DATABASE_URL` is not set or is empty, the server falls back to the `docker compose` database: `postgres://elsa:elsa@localhost:5432/elsa?sslmode=disable`.
- Tables are created and migrated automatically on startup.
- When running inside Docker, `localhost` refers to the container itself. Use the database host name instead (for example `host.docker.internal` or the compose service name).

## E2E tests

`POST /test/reset-db` drops and reseeds the database. It is only registered when `APP_ENV=test`, requires the `x-e2e-reset-secret` header to match `E2E_RESET_SECRET`, and refuses to run when `DATABASE_URL` points at the main `elsa` database. Use `elsa_test` instead:

```bash
APP_ENV=test E2E_RESET_SECRET=<secret> JWT_SECRET=<secret> DATABASE_URL="postgres://elsa:elsa@localhost:5432/elsa_test?sslmode=disable" go run .
```

# Run as Docker image


### Seeding the database in Docker

To seed the database when running in Docker, run the seeder CLI inside the container. For example:

```bash
docker run --rm -e DATABASE_URL="postgres://elsa:elsa@host.docker.internal:5432/elsa?sslmode=disable" ghcr.io/maastrichtu-biss/elsa-lifecycle-server go run cmd/seed/main.go
```

Adjust `DATABASE_URL` as needed for your setup.

**Warning:** Seeding will clear and repopulate the relevant tables. Only use this in development or when you want to reset the database.

