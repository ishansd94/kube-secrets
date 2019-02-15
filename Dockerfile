##################################
# STEP 1 build executable binary #
##################################
ARG GO_VERSION=1.11
ARG APP
ARG EXECUTABLE

FROM golang:${GO_VERSION}-alpine AS builder

ARG APP
ARG EXECUTABLE

RUN apk update --no-cache && apk add --no-cache ca-certificates git curl tzdata && update-ca-certificates

# Create appuser.
RUN adduser -D -g '' appuser

#Installing Dep. Package manager for Go
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep version

COPY . ${GOPATH}/src/${APP}
WORKDIR ${GOPATH}/src/${APP}

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ${EXECUTABLE} cmd/${EXECUTABLE}/main.go


##################################
# STEP 2 Run the executable      #
##################################
ARG APP
ARG EXECUTABLE

FROM scratch

ARG APP
ARG EXECUTABLE

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/${APP}/${EXECUTABLE} /${EXECUTABLE}

# Use an unprivileged user.
USER appuser

EXPOSE 8000

# TODO replace executable name with arg, may have to use jinja templating for dockerfiles
ENTRYPOINT ["./secret"]