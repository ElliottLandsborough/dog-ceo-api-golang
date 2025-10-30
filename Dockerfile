FROM public.ecr.aws/docker/library/golang:1.25.3 as builder

ARG BUILD_CONTROLLER

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY app ./app/
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -tags lambda.norpc -o ../bin/function ./app/controllers/${BUILD_CONTROLLER}

FROM public.ecr.aws/lambda/provided:al2 as runtime

COPY --from=builder /go/bin/function ./function

ENTRYPOINT [ "./function" ]
