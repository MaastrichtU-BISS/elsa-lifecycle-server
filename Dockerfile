FROM golang:1.24

ADD . /app

WORKDIR /app

# install all go dependencies
RUN go mod download

# compile
RUN go build -o /app/main .

EXPOSE 8080
CMD [ "/app/main" ]