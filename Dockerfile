FROM golang:1.19.3-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE  1234

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

CMD ["/app/main"]
