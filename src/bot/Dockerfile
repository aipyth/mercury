FROM golang:latest
RUN mkdir /bot
WORKDIR /bot
ADD . /bot
RUN go build -o main server.go
RUN ls -la
CMD [ "/bot/main" ]