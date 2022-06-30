FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download && go mod verify

RUN go build -o main .

EXPOSE 8081

CMD [ "/app/main" ]
