# Backend setup
FROM golang:1.22-alpine3.18
WORKDIR /app/backend
COPY . .
RUN mkdir -p /app/backend/bin 
RUN go build -o ./bin/main ./main.go
ENTRYPOINT [ "/app/backend/bin/main" ]
