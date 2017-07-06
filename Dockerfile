FROM alpine
MAINTAINER Wendell Sun <iwendellsun@gmail.com>

WORKDIR /ignite-admin

ADD ignite-admin /ignite-admin/ignite-admin
ADD static /ignite-admin/static

EXPOSE 8000
ENTRYPOINT ["/bin/sh", "-c", "./ignite-admin"]
