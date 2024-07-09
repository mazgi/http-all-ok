# ================
# Builder image
FROM ghcr.io/mazgi/go-app.development AS builder

ARG GIT_VERSION_TAG
ARG GIT_SHORT_SHA
COPY --from=workspace . /workspace/
WORKDIR /workspace/http-ok-whatever
RUN :\
  && : Build the app: appVersion=${GIT_VERSION_TAG}, appGitShortSha=${GIT_SHORT_SHA}\
  && eval "$(goenv init -)"\
  && go build -ldflags "-X main.appVersion=${GIT_VERSION_TAG} -X main.appGitShortSha=${GIT_SHORT_SHA}"\
  && :


# ================
# Runner image
FROM busybox AS runner

LABEL maintainer="docker@mazgi.com"
ARG RUNNER_USER_GID=10000
ARG RUNNER_USER_UID=10000
ENV PORT=8080
EXPOSE ${PORT}

RUN :\
  && : Create the user: UID=${RUNNER_USER_UID}, GID=${RUNNER_USER_GID}\
  && addgroup -g ${RUNNER_USER_GID} runner\
  && adduser -u ${RUNNER_USER_UID} -G runner -s /bin/sh -D runner\
  && :

COPY --chown=runner:runner --from=builder /workspace/http-ok-whatever/http-ok-whatever /bin/http-ok-whatever
CMD [ "http-ok-whatever" ]

HEALTHCHECK --interval=2s --timeout=1s --retries=2 --start-period=5s\
  CMD wget --quiet --output-document - "http://localhost:${PORT}/"
