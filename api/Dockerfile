FROM golang:1.9

WORKDIR /go/src/app

COPY ./vendor/ ./vendor/
COPY ./conf/ ./conf/
COPY ./*.go ./

#RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 3000

CMD ["go-wrapper", "run"] # ["app"]
