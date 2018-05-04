FROM jmfirth/webpack:6-slim as builder-frontend
WORKDIR /ignite-admin
COPY . .
RUN cd fe && yarn && yarn build

FROM golang:1.9 as builder-backend
ARG VERSION
WORKDIR /go/src/github.com/go-ignite/ignite-admin
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o ignite-admin 

FROM alpine
LABEL maintainer="go-ignite"
RUN apk --no-cache add ca-certificates tzdata sqlite \
			&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
			&& echo "Asia/Shanghai" >  /etc/timezone \
			&& apk del tzdata
# See https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
VOLUME /root/ignite/data

WORKDIR /root/ignite-admin
COPY --from=builder-backend /go/src/github.com/go-ignite/ignite-admin/ignite-admin ./
COPY --from=builder-backend /go/src/github.com/go-ignite/ignite-admin/conf ./conf
COPY --from=builder-frontend /ignite-admin/static ./static
RUN mv ./conf/config-temp.toml ./conf/config.toml

EXPOSE 8000
ENTRYPOINT ["./ignite-admin"]
