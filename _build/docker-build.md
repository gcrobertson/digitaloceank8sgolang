### Build the Dockerfile into a Docker Image locally:

Run commands in the project root [millicent/]:

> docker build -t gcrobertson/gregorysolutions:v1.0.1 .

### Run the Docker Image as a Docker Container locally:

> docker run -it -p 80:8080 gcrobertson/gregorysolutions:v1.0.1

### Prune all images
> docker image prune --all

### Build the Dockerfile for production:

> docker build -t gcrobertson/gregorysolutions:v1.0.1 .

### Push the Dockerfile to Dockerhub

> docker push gcrobertson/gregorysolutions:v1.0.1

### Access through URL:
> curl localhost

