# bank-account-demo

## Architecture
- Layered architecture using DDD(Domain-Driven Design)
  - Application layer
  - Domain model layer
  - Infrastructure layer
- Project layout
  ```
  .
  ├─ application
  │  ├─ request
  │  ├─ response
  │  └─ service
  ├─ domain
  │  ├─ aggregate
  │  ├─ entity
  │  ├─ factory
  │  ├─ repository
  │  └─ valueobject
  └─ infrastructure
     ├─ persistence
     └─ sqlite
  ```

## API Endpoint
- **GET** `/health`
- **POST** `/account/deposit`
- **POST** `/account/withdraw`
- **POST** `/mass/deposit`

## Req/Res Data
- Request body `./application/request`
- Response body `./application/response`

## Web Server
- Start web server
  ```bash
  go run main.go
  ```

## Database
- Schema path `./infrastructure/sqlite`
- Connect database
  ```bash
  sqlite3 ./infrastructure/sqlite/bank.db
  ```

## Version
- Language : `Go` _v1.16_
- Web framework : `Echo` _v4.6.1_
- Database : `SQLite` _v3.32.3_
- Database driver : `go-sqlite3` _v1.14.9_
