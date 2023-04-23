# Example of using a scratch base image.


FROM golang:1.20.3-alpine3.17 as builder

# Need to create user and group now as adduser are not available in scratch. Will copy the
# passwd and group files to the final image.

# Create appuser.
ENV USER=appuser
ENV UID=1000

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /go/src

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

RUN chown appuser:appuser ./app
RUN chown appuser:appuser -R ./static


FROM scratch
#FROM alpine:3.17

# Copy the passwd file and specify USER.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER appuser:appuser

# Copy ca-certificates from builder. In some cases this will be needed.
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /go/src/app .
COPY --from=builder /go/src/static ./static

EXPOSE 8080

CMD [ "./app" ]