FROM golang:1.9

WORKDIR /go/src/ct-sso
COPY . .

RUN go-wrapper download && go-wrapper install

EXPOSE 80

CMD ["go-wrapper", "run"]
