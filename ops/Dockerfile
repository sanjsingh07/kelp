FROM golang:1.13.8
LABEL maintainer="Demon <Github: @sanjsingh07>"

# # install yarn
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt update
RUN apt -y install yarn
RUN apt install -y nfs-common
RUN apt install -y cifs-utils

#install zip to fix unzip ccxt issue
RUN apt install -y zip

# set working dir
WORKDIR /go/src/github.com/stellar/kelp

#gopath
ENV GOPATH /go

# install glide
RUN go get -u github.com/Masterminds/glide

# install kelp
RUN git clone https://github.com/orunpay/kelp .
RUN git checkout deployment
RUN git fetch --tags
RUN glide install

# install astielectron
RUN go get -u github.com/asticode/go-astilectron-bundler/...
RUN go install github.com/asticode/go-astilectron-bundler/astilectron-bundler

# run build script
RUN ./scripts/build.sh

# set ulimit
RUN ulimit -n 10000

# use command line arguments from invocation of docker run against this ENTRYPOINT command - https://stackoverflow.com/a/40312311/1484710
ENTRYPOINT ["./bin/kelp"]
# sample command to run this container as a daemon process:
# docker run -v `pwd`/ops:/go/src/github.com/stellar/kelp/bin/ops -d -p 8000:8000 -p 3000:3000 <image> server -p 8000
# this assumes that you are running ccxt on port 3000 outside this kelp docker container
# the three port numbers (8011 in the example above) must be the same and must be specified
# the first part of the -v argument is the directory where you want to save the kelp configs and kelp logs from the bots created in this container