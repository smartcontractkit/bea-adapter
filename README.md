# BEA Adapter

## Run with Docker

Build the image:

```bash
docker build . -t cl-bea-adapter
```

Run the Docker image:

```bash
docker run -d \
    -p 8080:8080 \
    -e API_KEY="Your_bea_api_key" \
    cl-bea-adapter
```

Send requests to: `http://docker-ip:8080/call`

## Configuration

| Key | Description |
|-----|-------------|
| `API_KEY` | Your BEA API key |

## Methods

This adapter will by default get the average of the 3 latest monthly DPCERG values.
No input is required.
