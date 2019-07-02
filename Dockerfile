FROM alpine AS build
RUN apk update && apk add alpine-sdk go git
ENV GOPATH=/go
COPY cmd /go/src/cmd
WORKDIR /go/src
RUN cd cmd/server && go build --ldflags '-s -w -extldflags "-static"' --buildmode pie

FROM scratch
COPY --from=build /go/src/cmd/server/server /server
CMD [ "/server" ]
