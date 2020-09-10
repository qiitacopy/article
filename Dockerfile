# Step1: build
FROM golang:1.15.0-alpine3.12 as build-step

RUN apk add --update --no-cache ca-certificates git

RUN mkdir /article
WORKDIR /article
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/article

# Step2: exec
FROM scratch

ENV DB_USER ""
ENV DB_PASS ""
ENV DB_PATH ""
ENV DB_NAME ""

COPY --from=build-step /go/bin/article /go/bin/article
ENTRYPOINT ["/go/bin/article"]