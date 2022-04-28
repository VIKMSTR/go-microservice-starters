#FROM --platform=linux/amd64 golang:1.18 as builder
FROM --platform=$BUILDPLATFORM golang:1.18 as builder
#ENV GIN_MODE release
#ENV PORT 8000
### We create an /app directory in which
## we'll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
## We want to build our application's binary executable
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -tags=nomsgpack -o service ./...

## the lightweight scratch image we'll
## run our application within
#FROM --platform=linux/amd64 alpine:latest AS production
FROM alpine:latest AS production
ENV GIN_MODE=release
ENV PORT=8000
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /app/service /bin/service


## we can then kick off our newly compiled
## binary exectuable!!

#RUN export GIN_MODE=release && export PORT=8000
EXPOSE $PORT
CMD ["/bin/service"]
