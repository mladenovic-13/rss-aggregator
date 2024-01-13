# RSS Aggregation REST API

## Description

This project is a REST API for aggregating RSS feeds. It utilizes PostgreSQL as the database, sqlc as the ORM, and the chi framework in the GoLang ecosystem. Authorization is implemented using bearer tokens.

## Installation

1. Install Go and set up your workspace.
2. Clone the repository:
   ```bash
   git clone https://github.com/mladenovic-13/rss-aggregator.git
   cd rss-aggregation-api
   ```

## Usage

1. Create your .env file following the example provided in .env.example.

2. Build and run the application:

```bash
make migrate-up
make generate
make run
```

3. The API will be accessible at http://localhost:8000/v1/. Use tools like curl or Postman for interacting with the API.

## Authorization

The API uses bearer tokens for authorization. Include the token in the request header:

```makefile
  Authorization: Bearer YOUR_ACCESS_TOKEN
```

Authorization: Bearer YOUR_ACCESS_TOKEN

## API Routes

- `GET /healthz`: Checking if the server is live.
- `POST /users`: Create user.
- `GET /users`: Get user.
- `POST /feeds`: Create feed (private).
- `GET /feeds`: Get feeds.
- `POST /feeds`: Create feed following (private).
- `GET /feeds`: Get feed follows (private).
- `DELETE /feeds/{feedFollowId}`: Delete feed follows (private).
