##################################
# STEP 1 build executable binary #
##################################
ARG GO_VERSION=1.11
ARG SSH_PRIVATE_KEY
ARG APP
ARG EXECUTABLE
ARG COMMIT
ARG BUILDTIME
ARG RELEASE

FROM golang:${GO_VERSION}-alpine AS builder

ARG SSH_PRIVATE_KEY
ARG APP
ARG EXECUTABLE
ARG COMMIT
ARG BUILDTIME
ARG RELEASE

RUN apk update --no-cache && apk add --no-cache ca-certificates git curl tzdata openssh openssl && update-ca-certificates

#adding out ssh key to container
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && chmod 400 /root/.ssh/id_rsa && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN eval `ssh-agent -s` && ssh-add /root/.ssh/id_rsa

RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Create appuser.
RUN adduser -D -g '' appuser

#Installing Dep. Package manager for Go
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep version

COPY . ${GOPATH}/src/${APP}
WORKDIR ${GOPATH}/src/${APP}

RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -a -installsuffix cgo \
        -ldflags="-w -s \
            -X ${APP}/internal/pkg/version.Release=${RELEASE} \
            -X ${APP}/internal/pkg/version.Commit=${Commit} \
            -X ${APP}/internal/pkg/version.BuildTime=${BUILDTIME}" \
        -o ${EXECUTABLE} \
        cmd/${EXECUTABLE}/main.go

RUN rm /root/.ssh/id_rsa

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

ENTRYPOINT ["./secret"]