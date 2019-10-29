# BEA Adapter

## Run in Lambda

- Install packages

```
go install
```

- Build the binary

```
go build -o cl-bea
```

- Zip the binary

```
zip cl-bea.zip cl-bea
```

- Upload to Lambda, use `cl-bea` as your handler
- Add the value for the `API_KEY` environment variable

## Configuration

| Key | Description |
|-----|-------------|
| `API_KEY` | Your BEA API key |

## Methods

This adapter will by default get the average of the 3 latest monthly DPCERG values.
No input is required.
