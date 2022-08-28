FROM alpine:latest

MAINTAINER Mick Hellstrom <embed@mickmake.com>

USER root
ARG GO_REPO_TOKEN
ENV GO_REPO_TOKEN ${GO_REPO_TOKEN}
ARG TZ
ENV TZ ${TZ}

COPY dist/GoWhen_linux_amd64_v1/GoWhen /usr/local/bin/GoWhen
COPY .ssh/ /root/.ssh/
RUN \
	chmod 500 /root/.ssh && \
	apk add --no-cache colordiff tzdata && \
	chmod a+x /usr/local/bin/GoWhen

#ENTRYPOINT ["/usr/local/bin/GoWhen"]
CMD ["/usr/local/bin/GoWhen"]

