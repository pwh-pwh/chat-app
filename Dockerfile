FROM golang:1.21.5-bookworm
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o main ./
CMD ["/app/main"]