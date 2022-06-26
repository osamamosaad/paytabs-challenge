## How to run the app
- Run `go mod tidy` to install dependencies
- Run  `go run .` from the root folder to start the application
- If you faced any problim for starting the HTTP server, maybe the problem related to the port or local URL. You can change the `SERVER_URL` & `SERVER_PORT` from config file [config/config.go](config/config.go)

## Test
- My focus was to write tests for `transaction-library` as sample and also because the most of business logic located there
- Run `go test -v  pkgs/transaction-library/*.go` to run

---
## API endpoints
### List All accounts
- URL: localhost:1323/accounts
- Method: GET
- CURL code: `curl --request GET \
  --url http://localhost:1323/accounts`

### Get account by ID
- URL: localhost:1323/accounts/{ID}
- Method: GET
- CURL code: 
    ```
    curl --request GET \
    --url http://localhost:1323/accounts/17f904c1-806f-4252-9103-74e7a5d3e340
    ```

### Make a transaction
- URL: localhost:1323/transactions
- Method: POST
- Body:
    ```
    {
        "to": "17f904c1-806f-4252-9103-74e7a5d3e340",
        "from": "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
        "amount": 408.11
    }
    ```

- CURL code: 
    ```
    curl --request POST \
    --url http://localhost:1323/transactions \
    --header 'Accept: application/json' \
    --header 'Content-Type: application/json' \
    --data '{
        "to": "17f904c1-806f-4252-9103-74e7a5d3e340",
        "from": "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
        "amount": 408.11
    }'
    ```
### _**NOTE:** please change the url and the port if you change it from config file._
