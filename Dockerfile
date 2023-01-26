FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goecho


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goecho /app/

EXPOSE 80

HEALTHCHECK --interval=5s --timeout=3s CMD curl --fail http://localhost:80/health || exit 1

ENTRYPOINT ./goecho