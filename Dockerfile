FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep

# Set ENV
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin/

ENV DB_NAME root

RUN mkdir -p $GOPATH/src/templatetest22
WORKDIR $GOPATH/src/templatetest22/
RUN go get github.com/pilu/fresh

CMD ["/bin/sh", "start.sh"] 


