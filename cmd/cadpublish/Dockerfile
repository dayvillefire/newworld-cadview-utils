FROM alpine:latest
WORKDIR /app
RUN apk add gcompat
RUN mkdir -p /app
COPY ./cadpublish /app/
CMD ["/app/cadpublish", "--config", "/app/config.yaml"]

