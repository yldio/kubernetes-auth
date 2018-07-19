FROM golang:1.10.3-alpine

RUN apk add --no-cache --update alpine-sdk

COPY . /go/src/github.com/yldio/kubernetes-auth
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep
WORKDIR /go/src/github.com/yldio/kubernetes-auth
RUN dep ensure -vendor-only
RUN make release-binary

FROM alpine:3.4
RUN apk add --update ca-certificates openssl

WORKDIR /go/src/github.com/yldio/kubernetes-github-auth
COPY --from=0 /go/src/github.com/yldio/kubernetes-auth/bin/kubernetes-auth /usr/local/bin/k8s-auth
WORKDIR /

ENTRYPOINT ["k8s-auth"]
