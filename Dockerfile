FROM tinygo/tinygo:latest as build

COPY . /src
WORKDIR /src
# No git in the docker image
ENV GOFLAGS="-buildvcs=false"
# -x
RUN --mount=target=/home/tinygo/.cache,id=apt,type=cache,uid=1000,gid=1000 \
    cd /src && tinygo build -gc conservative -x -no-debug -scheduler tasks  \
    -o /tmp/sftp-server  ./cmd/sftp-server && ls -l /tmp/sftp-server /src /

RUN strip /tmp/sftp-server && ls -l /tmp/sftp-server

FROM scratch

# ~400k - compared to 4.3M (2.9 stripped)
COPY --from=build /tmp/sftp-server /bin/sftp-server
