# ELSA lifecycle tool server

## How to run?

Install Go on your system https://go.dev/doc/install

### Windows?
Also install mingw-w64 GCC through https://www.msys2.org/. GCC is used to run C code.

Enable C code to run using Go, since sqlite relies on it (see [go.mod](./go.mod))

```bash
go env -w CGO_ENABLED=1
```

To run the project locally in development mode, you can run the following command:
```bash
go run .
```

## Class diagram
First draft of the class / model diagram can be found below. This is made using drawio where the xml metadata is encapsulated in the .png metadata. This can be edited in VS Code using the [Draw.io Integration plugin](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio).

![class diagram](./uml_model.drawio.png "UML model")
