FROM golang:1.17 AS build
WORKDIR /homework_1/
COPY source_code .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver homework1.go

FROM busybox
COPY --from=build /homework_1/httpserver /httpserver/httpserver
EXPOSE 8080
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]