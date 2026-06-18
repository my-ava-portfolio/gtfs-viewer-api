# https://www.bacancytechnology.com/blog/dockerize-golang-application

# Build Stage
# First pull Golang image
FROM golang:1.19-alpine as build-env
 
# Set environment variable
ENV APP_NAME app
ENV MAIN_FILE_PATH src/main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$MAIN_FILE_PATH
 
# Run Stage
FROM alpine:3.17.1
 
# Set environment variable
ENV APP_NAME app
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# if running localy
# COPY src/data data

RUN addgroup -S appgroup && adduser -S ava -G appgroup
USER ava

# Start app
CMD ./$APP_NAME
