FROM golang:1.22-alpine3.18
RUN apk add nodejs npm
WORKDIR /app
COPY . .
WORKDIR /app/frontend
RUN npm i
RUN npm run build
WORKDIR /app/backend
RUN mkdir bin & go build -o ./bin/ main.go

ENTRYPOINT [ "./bin/main" ]
