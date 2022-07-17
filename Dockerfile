# 1 - generate a new binary for use
FROM golang:1-bullseye as build
RUN apt-get update && \
    apt-get install wget
RUN wget https://github.com/ImageMagick/ImageMagick/releases/download/7.1.0.43/ImageMagick--gcc-x86_64.AppImage && \
    mv ImageMagick--gcc-x86_64.AppImage magick &&\
    chmod +x magick && \
    mv magick /usr/bin/
RUN mkdir /opt/app
WORKDIR /opt/app
COPY . .
RUN make build

# 2 - use the newly built image
FROM debian:bullseye
RUN apt-get update && \
    apt-get install -qqy x11-apps && \
    rm -rf /var/lib/apt/lists/*
ENV DISPLAY :0
COPY --from=build /usr/bin/magick /usr/bin
COPY --from=build /opt/app/build/clingy /usr/bin
RUN groupadd --gid 1001 clingy && \
    useradd --uid 1001 --gid 1001 -m clingy && \
    usermod -a -G users clingy && \
    chown -R clingy /home/clingy
USER clingy
RUN mkdir /home/clingy/images
WORKDIR /home/clingy
ENTRYPOINT ["/usr/local/bin/clingy"]