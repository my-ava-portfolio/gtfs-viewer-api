# Build Stage
FROM golang:1.19-alpine AS build-env

ARG TARGETOS
ARG TARGETARCH

ENV APP_NAME=app
ENV MAIN_FILE_PATH=src/main.go
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$MAIN_FILE_PATH

# Run Stage
FROM alpine:3.17.1
ENV APP_NAME=app

COPY --from=build-env /$APP_NAME .

RUN addgroup -S appgroup && adduser -S ava -G appgroup
USER ava

CMD ["./app"]