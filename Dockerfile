FROM golang:1.26.1-alpine AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /app/src  \
    && mkdir -p /app/bin

WORKDIR /app/src

COPY go.mod go.sum ./
RUN go mod download &&  \
    go mod verify

COPY . ./

RUN go build -v \
    -o /app/bin/server \
    -ldflags="-s -w" \
    /app/src/cmd/server

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/bin/server /server
COPY --from=build /app/src/public /public

EXPOSE 8080

USER nonroot:nonroot

CMD [ "/server" ]
