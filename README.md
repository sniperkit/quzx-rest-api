# demas/cowl-services

# Introduction

# Installation

Automated builds of the image available on [Dockerhub](...) and is the recommended method of installation.

```bash
docker pull demas/cowl-services
```

Alternatively you can build the image locally.

```bash
docker build -t demas/cowl-services .
```

# Quick Start

The quickies way to get started is using [docker-compose](https://docs.docker.com/compose/).

Alternatively, you can manually launch the services container:

```bash
docker run --name cowl-services -d \
           --env 'TWICONKEY=...' \
           --env 'TWICONSEC=...' \
           --env 'TWIACCTOK=...' \
           --env 'TWIACCTOKSEC=...' \
           --env 'USER=...'\
           --env 'PASS=...' \
           --env 'SECRET=...' \
           --publish 4000:4000 \
           --restart=always \
           demas/cowl-services

```