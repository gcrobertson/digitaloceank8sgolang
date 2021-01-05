# golang:1.14           = 820 mb
# golang:1.14-alpine    = 381 mb
FROM golang:1.14-alpine

WORKDIR /go/src/app

# 
COPY ./vendor .
#
#
#
COPY . .
#
#   @TODO: Improve cache!
#

# Download and install a dependency
RUN go get -d -v ./...
RUN go install -v ./...

# Tell the image what to do when it starts as a container
CMD ["go", "run", "main.go"]



