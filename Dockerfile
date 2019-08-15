FROM alpine

RUN apk update && apk add curl

COPY ./demoapp /demoapp

ENTRYPOINT ["/demoapp"]
