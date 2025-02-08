# Running the application

## Go commands

```bash
go mod download
go build -o server .
go run ./server
```

## Makefile

Confirm the configuration is correct for your environment

```make
make run
```

## Dockerfile

```docker
docker build -t kyler-receipt-processor .
docker run --network="host" -t kyler-receipt-processor
```

# Endpoints

## Endpoint: Process Receipts

- Path: /receipts/process
- Method: POST
- Payload: Receipt JSON
- Response: JSON containing an id for the receipt.

Example Payload:

```json
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
```

Example Response:

```json
{
  "id": "8e870d34-858f-496e-8591-84c5664e4856"
}
```

## Endpoint: Get Points

- Path: /receipts/{id}/points
- Method: GET
- Response: A JSON object containing the number of points awarded.

Example Request:

`/receipts/6fcf0098-75f2-47f5-ad2d-0973729667fb/points`

Example Response:

```json
{
  "points": 15
}
```
