FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
#设置环境
ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"
WORKDIR /go/src/bc-order

# Install dependencies
# RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY Makefile go.mod go.sum ./
# RUN make init && go mod download
COPY . .
#RUN make proto tidy build
RUN go build -o bc-order main.go
# Deployment container
FROM scratch

#COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/bc-order/bc-order /bc-order
ENTRYPOINT ["/bc-order"]
CMD []
