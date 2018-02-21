FROM golang:1.10

WORKDIR /go/src/ct-authentication-server
COPY . .

RUN go get -d -v ./... \
    && go install -v ./... \
    && go get -u -d github.com/mattes/migrate/cli \
    && go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

EXPOSE 80

CMD ["ct-authentication-server"]
