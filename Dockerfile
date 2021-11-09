# syntax=docker/dockerfile:1
# golang:1.16-alpine
FROM golang:1.16-bullseye

WORKDIR /build

COPY . ./

RUN go build -o raindrops-go

RUN apt update && apt install -y netcat ncat iproute2 net-tools

CMD [ "/build/raindrops-go" ]

EXPOSE 80