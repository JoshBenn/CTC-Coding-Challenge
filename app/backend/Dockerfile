# Backend setup
FROM golang:1.22-alpine3.18
WORKDIR /app/backend
COPY . .
RUN mkdir bin && go build -o ./bin/ main.go
ENTRYPOINT [ "./bin/main" ]
