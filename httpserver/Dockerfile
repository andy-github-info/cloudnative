FROM golang:alpine3.14 AS build
COPY src  httpserver/src
WORKDIR httpserver/src
ENV CGO_ENABLED=0
RUN  go build main.go

FROM scratch
WORKDIR httpserver/
COPY --from=build /go/httpserver/src/main ./
EXPOSE 80
ENTRYPOINT ["./main"]
