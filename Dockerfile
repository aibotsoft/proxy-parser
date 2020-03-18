#FROM scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY dist/ .

CMD ["/service"]