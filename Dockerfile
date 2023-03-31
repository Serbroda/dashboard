# syntax=docker/dockerfile:1

FROM node:18-alpine AS build-frontend

WORKDIR /work

COPY frontend/package.json .

RUN npm install

COPY frontend/ .

RUN npm run build


FROM golang:1.20-alpine AS build-go

WORKDIR /work

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY --from=build-frontend /work/dist ./frontend/dist

RUN GOOS=linux GOARCH=amd64 go build -o dashboard main.go


FROM alpine:3.9 AS run

WORKDIR /app

COPY --from=build-go /work/dashboard ./dashboard

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/dashboard"]
