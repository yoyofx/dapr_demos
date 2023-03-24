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

# Deploying your code in kubernetes
* Update Golang code as you see fit!
* Navigate to the directory of the app you want to build a new image for.
* Run docker build -t <YOUR_IMAGE_NAME> . . You can name your image whatever you like. If you're planning on hosting it on docker hub, then it should start with <YOUR_DOCKERHUB_USERNAME>/.
* Once your image has built you can see it on your machines by running docker images.
* To publish your docker image to docker hub (or another registry), first login: docker login. Then rundocker push <YOUR IMAGE NAME>.
* Update your .yaml file to reflect the new image name.
* Deploy your updated Dapr enabled app: kubectl apply -f <YOUR APP NAME>.yaml.

Build server and client Docker images, and then push them afterwards.

Example:
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

---
// state server
---
kubectl apply -f ./componsents/redis_state.yaml
kubectl apply -f stateserver.yaml
```

kubectl forward {stateserver service} 3000
```bash
curl -X POST -H "Content-Type: application/json" -d '{"value":"hello,world!"}' http://localhost:58522/state/value
```
