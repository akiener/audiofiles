FROM golang:1.16.2 as go-builder

ENV GOLANGCI_VERSION 1.38.0
ENV GOLANGCI_BINARY golangci-lint-"$GOLANGCI_VERSION"-linux-amd64.deb

RUN apt update && apt install -y upx-ucl

RUN wget https://github.com/golangci/golangci-lint/releases/download/v$GOLANGCI_VERSION/$GOLANGCI_BINARY && \
	dpkg -i $GOLANGCI_BINARY && \
	rm $GOLANGCI_BINARY

FROM go-builder

ENV APP_HOME /audiofiles
WORKDIR $APP_HOME
COPY . .

RUN make lint
RUN make test
RUN make build

CMD .ci/audiofiles