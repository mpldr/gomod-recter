---
title: "How to deploy recter"
---

# How to deploy recter

see also: [configuration](./config.md)

## Using Docker

A docker container is provided as
[docker.io/mpldr/recter](https://hub.docker.com/r/mpldr/recter). Only
tagged versions are provided.

You can use a simple `docker-compose.yml`. Feel free to use the
following as a startingpoint and configure it to your needs.

```yaml
version: "3"

services:
  recter:
    image: docker.io/mpldr/recter:latest
    environment:
      - GLOG_LEVEL=warn
    ports:
      - 127.0.0.1:25000:25000
    volumes:
      - "./data:/data"
      - "./theme:/themes/bulma"
```

### Pitfalls

- If using TCP communication:
	- set the interface to run on to `0.0.0.0`
- If using UNIX-sockets
	- the permissions are taken from the parent directory, however
	  group and owner cannot be configured as there are no users
	  or groups other than root inside the container.
	- this leads to permissions always being xx6 which is not
	  recommended. Prefer to use the TCP interface if possible

## Without Docker

- get a binary
	- compile it yourself
	- get it from the [release page](https://git.sr.ht/~poldi1405/gomod-recter/refs)
	- create a package for your operating system
