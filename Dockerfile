FROM golang:1.21.6-alpine

ARG EXPORT_PORT

RUN apk --no-cache add ca-certificates git

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go api/api.go migrate
RUN go api/api.go seed

RUN go build -v api/api.go -o build/api

EXPOSE ${EXPORT_PORT}

CMD [ "./build/api", "serve" ]