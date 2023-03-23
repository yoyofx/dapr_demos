# Install Dapr (for mac)
CLI version: 1.9.1

Runtime version: 1.10.0
```bash
brew install dapr/tap/dapr-cli
```

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

# Run all in kubernetes
Build server and client Docker images, and then push them afterwards.
```bash
cd server
docker build docker build  -t {your_repository}/server:v0.1 . --platform=linux/amd64

cd client
docker build docker build  -t {your_repository}/client:v0.1 . --platform=linux/amd64
```
```bash
cd deploy

kubectl apply -f server.yaml

kubectl apply -f client.yaml
```