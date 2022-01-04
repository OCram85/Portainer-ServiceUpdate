FROM alpine:3.15
LABEL maintainer="marco.blessing@googlemail.com"
ADD ServiceUpdater /app/
ENTRYPOINT "/app/ServiceUpdater"
