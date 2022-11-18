FROM golang:1.19.3-buster

RUN apt update \
&& apt install -y git

COPY . /kserve
WORKDIR /kserve
#RUN GOOS=darwin GOARCH=arm64 go build -mod vendor -o main ./main.go
RUN go build -mod vendor -o main ./main.go

FROM alpine

COPY --from=0 /kserve/main /kserve/main
CMD ["/kserve/main"]