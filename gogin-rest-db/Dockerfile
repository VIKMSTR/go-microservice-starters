#This Dockerfile is made for cross-platform builds
#When running on standard docker build, the BUILDPLATFORM,TARGETOS and TARGETARCH are set to the host platform (where do you run docker from)
#First stage - build the app
FROM --platform=$BUILDPLATFORM golang:1.18 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
#Building here for the platform
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -tags=nomsgpack -o service ./...

#The production image.
#Using alpine as the base image, might want to use scratch if security is an issue. Or debian-slim if you want more power.
FROM alpine:latest AS production
#GIN_MODE sets the build for the production mode.
ENV GIN_MODE=release
#PORT specifies on which port will the app listen.
ENV PORT=8000
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /app/service /bin/service
#EXPOSE the port for the outside world to be available
EXPOSE $PORT
#Run it
CMD ["/bin/service"]
