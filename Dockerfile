FROM golang:alpine3.6 AS binary

ADD . /go/src/github.com/trollbot
WORKDIR /go/src/github.com/trollbot

ENV CGO_ENABLED=0

RUN apk add --no-cache git \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && go build -a -ldflags '-s' -installsuffix cgo -o trollBot


#FROM scratch
FROM alpine:3.6
WORKDIR /app
ENV SLACK_TOKEN <change_me>

RUN apk --no-cache add ca-certificates
COPY --from=binary /go/src/github.com/trollbot/trollBot /app
CMD ["/app/trollBot"]

ARG BUILD_DATE
ARG BUILD_VCS_REF
ARG BUILD_VERSION

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.vcs-url="https://github.com/mgvazquez/trollBot.git" \
      org.label-schema.vcs-ref=$BUILD_VCS_REF \
      org.label-schema.version=$BUILD_VERSION \
      com.microscaling.license=MIT