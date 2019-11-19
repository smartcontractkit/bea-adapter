# BEA Adapter

## Run in Docker

- make docker and run
```
make docker
docker run -p 8080:8080 -e API_KEY=<yourapikey> bea-adapter
```

## Run in Lambda

- Install packages

```
go install
```

- Build the binary

```
GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o cl-bea
```

- Zip the binary

```
zip cl-bea.zip cl-bea
```

- Upload to Lambda, use `cl-bea` as your handler
- Add the value for the `API_KEY` environment variable
- Enable Lambda support by setting the `LAMBDA` environment variable to `1`

## Configuration

| Key | Description |
|-----|-------------|
| `API_KEY` | Your BEA API key |
| `LAMBDA` | Set to 1 |

## Methods

This adapter will by default get the average of the 3 latest monthly DPCERG values.
No input parameters are required.
