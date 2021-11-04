# bank-account-demo

## API
- **GET** `/health`
- **POST** `/account/deposit`
- **POST** `/account/withdraw`
- **POST** `/mass/deposit`

## Req / Res Body
- `./jsondata/jsondata.go`

## Layer
- __handler__ → __validate__ or __repository__ → __model__

## Database
- `./sqlite/bank.db`

## Version
- Language : `Go` _v1.16_
- Web Framework : `Echo` _v4.6.1_
- Database : `SQLite` _v3.32.3_
- Database Driver : `go-sqlite3` _v1.14.9_

## Usage
- Start API server
  ```bash
  go run main.go
  ```
