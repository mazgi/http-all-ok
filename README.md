# http-ok-whatever

[![publish-the-app-as-container-images](https://github.com/mazgi/http-ok-whatever/actions/workflows/publish-the-app-as-container-images.yaml/badge.svg)](https://github.com/mazgi/http-ok-whatever/actions/workflows/publish-the-app-as-container-images.yaml)

It's an HTTP stub app that returns 200 OK to all requests.

### How to build Docker images for Production

Run the `gh workflow run publish-app-production-assets` command.
Or, you are able to manually build as follows.

Using `docker compose build`:

```console
git diff-files --quiet\
 Dockerfile.d/go-app.production\
 workspace\
 compose.yaml && {
  local _V=$(git tag --sort=-version:refname | head -n 1)
  docker compose build\
   --build-arg=GIT_VERSION_TAG=${_V}\
   --build-arg=GIT_SHORT_SHA=$(git rev-parse --short HEAD)\
   --no-cache go-app.production
}
```
