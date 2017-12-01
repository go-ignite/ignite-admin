FROM jmfirth/webpack:6-slim as builder-frontend
COPY . /app
RUN ls -lh
RUN yarn install \
			&& webpack

FROM golang:1.9 as builder-backend
WORKDIR /go/src/github.com/go-ignite/ignite-admin
COPY --from=builder-frontend /app .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ignite-admin .

FROM alpine
LABEL maintainer="iwendellsun@gmail.com"
RUN apk --no-cache add ca-certificates tzdata sqlite \
			&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
			&& echo "Asia/Shanghai" >  /etc/timezone \
			&& apk del tzdata
# See https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
VOLUME /root/ignite-admin/data

WORKDIR /root/ignite-admin
COPY --from=builder-backend /go/src/github.com/go-ignite/ignite-admin/ignite-admin ./
COPY --from=builder-backend /go/src/github.com/go-ignite/ignite-admin/static ./static
COPY --from=builder-backend /go/src/github.com/go-ignite/ignite-admin/conf ./conf

EXPOSE 8000
CMD ["/bin/sh", "-c", "./ignite-admin"]
