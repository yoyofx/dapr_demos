# Run myapp http server
```bash
cd ./server

dapr run --app-port 8223 --app-id myapp --app-protocol http --dapr-http-port 3501 go run .
```
# Run client of myapp
```bash
cd ./client

dapr run --app-id myapp-client --app-protocol http  --dapr-http-port 3500 go run .
```

# Run all
if you installed dapr-cli version v1.11.0 , you should execute the following command
```bash
dapr run -f .
```