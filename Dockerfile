FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]




# FROM golang:alpine

# # Set the Current Working Directory inside the container
# WORKDIR /go/src/k8s-envoy

# # Copy everything from the current directory to the PWD(Present Working Directory) inside the container
# COPY . .

# # Download all the dependencies
# # https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# RUN go get -d -v ./...

# # Install the package
# RUN go install -v ./...

# # Run the executable
# CMD ["k8s-envoy"]