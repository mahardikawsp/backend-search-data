############################
# STEP 1 build executable binary
############################
FROM golang:1.18.5-buster AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apt update && apt install git -y && apt autoclean -y && apt autoremove -y
# Create appuser.
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR /app
COPY . .
# Fetch dependencies.
# Using go get.
RUN go clean && go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/apiservice && chmod +x /app/apiservice

############################
# STEP 2 build a small image
############################
FROM alpine
# Install dependency minimum for the image
RUN apk add gcompat libc6-compat
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /app /app
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta
# Use an unprivileged user.
USER appuser:appuser
# Run the hello binary.
ENTRYPOINT ["/app/apiservice"]
