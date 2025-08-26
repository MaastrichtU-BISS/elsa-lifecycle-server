# ELSA lifecycle tool server

## How to run?

To run the project locally in development mode, you can run the following command:
```bash
go run .
```

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

Run the image with the default DB path (the image contains no DB file by default — mount a volume or set `DB_PATH`):

	# mount a host directory that contains (or will contain) the DB
	docker run -p 8080:8080 -v $(pwd)/db:/app/database/db ghcr.io/maastrichtu-biss/elsa-lifecycle-server

Or explicitly point to a host path inside the container with `DB_PATH`:

	docker run -p 8080:8080 -e DB_PATH="/db/elsa.sqlite" -v $(pwd)/db:/db ghcr.io/maastrichtu-biss/elsa-lifecycle-server

