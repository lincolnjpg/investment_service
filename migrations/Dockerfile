FROM alpine

ENV GOOSE_DIR=/goose

RUN mkdir -p ${GOOSE_DIR}/migrations

RUN apk add curl

WORKDIR ${GOOSE_DIR}

RUN curl -fsSL \
	https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
	GOOSE_INSTALL=${GOOSE_DIR} sh

CMD ["tail", "-F", "non-existent-file"]

