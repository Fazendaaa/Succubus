FROM golang:1.16.14-alpine3.15 as base

ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/Fazendaaa/Succubus

COPY go.* .
COPY main.go .

COPY cmd cmd/
COPY controllers controllers/
COPY pkg pkg/
COPY test test/

RUN [ "go", "test", "./..." ]
RUN [ "go", "build" ]
RUN [ "mv", "Succubus", "/usr/bin/succubus" ]

FROM alpine:3.15 as runner
LABEL author="fazenda"
LABEL project="succubus"

COPY --from=base /usr/bin/succubus /usr/bin/succubus

ENTRYPOINT [ "succubus" ]
