# Use multistage builds.

# Create golang builder stage.
## Fetch OS updates and app dependencied.
FROM golang:1.16.3-alpine3.13 as builder
RUN apk update && apk add --no-cache git make
## Copy app source to image build directory.
RUN mkdir /build
ADD . /build/
## Build app from source in build directory.
WORKDIR /build
RUN make prep
RUN GOOS=linux make build

# Create final app image from builder stage.
FROM alpine:3.13
## Add application from builder stage to final image.
COPY --from=builder /build/devbox /app/
## Set directory to run in.
WORKDIR /app
## Set command to run.
ENTRYPOINT ["./devbox"]
