FROM golang:alpine


ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64


WORKDIR /build


COPY . .

RUN go env -w  GOPROXY=https://goproxy.cn,direct

RUN go build -o app .


WORKDIR /dist


RUN cp /build/app .

COPY ./config-dev.yml .


EXPOSE 8098

CMD ["/dist/app", "conf/config-dev.yml"]