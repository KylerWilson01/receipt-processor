FROM golang:1.23 as build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /server

FROM build-stage AS test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /server /server
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]
