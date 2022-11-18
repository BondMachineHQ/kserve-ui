FROM golang:1.19.3-buster

RUN apt update \
&& apt install -y git

RUN git clone https://github.com/kserve/kserve.git /kserve
COPY . /kserve
WORKDIR /kserve
RUN GOOS=darwin GOARCH=arm64 go build -mod vendor -o main ./main.go
CMD ["/kserve/main"]