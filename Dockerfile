FROM golang:stretch
FROM node:current-slim

RUN mkdir -p /kelp

WORKDIR /kelp

COPY . .

EXPOSE 3000 8000

CMD ["./bin/kelp", "server", "--no-electron", "-c", "./examples/configs/trader/sample_custom_config.cfg"]