FROM golang:1.19.3-buster as BUILD

RUN apt update \
&& apt install -y git

COPY . /kserve
WORKDIR /kserve
#RUN GOOS=darwin GOARCH=arm64 go build -mod vendor -o main ./main.go
RUN CGO_ENABLED=0 go build -mod vendor -o main ./main.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /kserve
COPY --from=BUILD /kserve/main .
CMD ["/kserve/main"]