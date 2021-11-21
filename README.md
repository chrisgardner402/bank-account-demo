# bank-account-demo

## Architecture
- Layered Architecture in DDD
  - Application layer
  - Domain model layer
  - Infrastructure layer
- Project Layout
  ```
  .
  ├─ application
  │  ├─ request
  │  ├─ response
  │  ├─ rest
  │  └─ service
  ├─ domain
  │  ├─ aggregate
  │  ├─ entity
  │  ├─ factory
  │  └─ valueobject
  └─ infra
     └─ persistence
        ├─ mapper
        ├─ repository
        └─ sqlite
  ```
## API Endpoint
- **GET** `/health`
- **POST** `/account/deposit`
- **POST** `/account/withdraw`
- **POST** `/mass/deposit`

## Json Data
- request body `./application/request`
- response body `./application/response`

## Web Server
- Start API server
  ```bash
  go run main.go
  ```

## Database
- schema path `./infra/persistence/sqlite`
- connect database
  ```
  sqlite3 ./infra/persistence/sqlite/bank.db
  ```

## Version
- Language : `Go` _v1.16_
- Web Framework : `Echo` _v4.6.1_
- Database : `SQLite` _v3.32.3_
- Database Driver : `go-sqlite3` _v1.14.9_
