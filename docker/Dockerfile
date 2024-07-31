# FROM golang:1.22.5

# WORKDIR /app

# ENV config=docker

# # Cài đặt Delve
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# COPY . /app

# RUN go mod download

# EXPOSE 9092
# EXPOSE 40000

# # Sử dụng Delve để chạy ứng dụng trong chế độ debug
# CMD ["dlv", "debug", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "--continue", "cmd/api/main.go"]
FROM golang:1.22.5

WORKDIR /app

ENV config=docker

COPY .. /app

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 9092

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main
