FROM golang:1.9

WORKDIR /go/src/ct-authentication-server
COPY . .

RUN go-wrapper download \
    && go-wrapper install \
    && go get -u -d github.com/mattes/migrate/cli \
    && go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

EXPOSE 80

CMD ["go-wrapper", "run"]
