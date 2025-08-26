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
