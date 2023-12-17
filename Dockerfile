FROM golang:1.21 AS build-stage

WORKDIR /

COPY ./ /
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /ims

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /ims /ims
COPY --from=build-stage /configs /configs

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/ims"]