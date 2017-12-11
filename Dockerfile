FROM golang:1.9

RUN mkdir -p /go/src/github.com/blueshirts/madziki
WORKDIR /go/src/github.com/blueshirts/madziki

COPY ./vendor/ ./vendor/
COPY ./conf/ ./conf/
COPY ./handlers/ ./handlers/
COPY ./*.go ./
COPY ./api/ ./api/

#RUN go-wrapper download   # "go get -d -v ./..."
#RUN go-wrapper install

#CMD ["go-wrapper", "run"]
#CMD tail -f /dev/null
#RUN go run github.com/blueshirts/madziki
#RUN go test -v ./...

RUN go install github.com/blueshirts/madziki

#RUN go install
EXPOSE 3000

#CMD go run madziki
CMD /go/bin/madziki