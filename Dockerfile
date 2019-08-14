FROM denismakogon/gocv-alpine:4.0.1-buildstage as build-stage

#install dep

RUN apk add wget && \
    wget https://raw.githubusercontent.com/golang/dep/master/install.sh && \
    chmod 755 install.sh && \
    ./install.sh

COPY . $GOPATH/src/classifier-example

WORKDIR $GOPATH/src/classifier-example

RUN dep ensure

RUN go build -o detect

FROM denismakogon/gocv-alpine:4.0.1-runtime

EXPOSE 8080

COPY --from=build-stage /go/src/classifier-example/detect detect

COPY classifiers ./classifiers

RUN ls

ENTRYPOINT ["./detect"]