FROM golang:alpine

WORKDIR /usr/src/app

COPY . /usr/src/app


RUN go build -o user_reg

EXPOSE 8081

ENTRYPOINT ["./user_reg"]