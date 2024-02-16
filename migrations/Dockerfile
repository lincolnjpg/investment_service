FROM alpine

ENV GOOSE_DIR=/goose

RUN mkdir -p ${GOOSE_DIR}/migrations

RUN apk add curl

WORKDIR ${GOOSE_DIR}

RUN curl -fsSL \
	https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=${GOOSE_DIR} sh

# command below is used to prevent the container to be closed just after the goose's installation
CMD ["tail", "-F", "any-file"]

