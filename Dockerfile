FROM golang:1.20-alpine AS BUILDER
RUN apk add --no-cache gcc g++ git openssh-client
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download
RUN GO111MODULE=on go build -ldflags="-w -s" -o server

FROM golang:1.20-alpine
ENV WORK_DIR=/app
WORKDIR $WORK_DIR
COPY --from=builder /app/server $WORK_DIR
EXPOSE 8080
ENTRYPOINT ./server