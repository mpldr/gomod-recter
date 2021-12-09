---
title: "Documentation for recter"
---

# Documentation for recter

## What is recter

Recter allows you to easily set up your own forwarding for Go-Modules. This
enables you to host your code on any platform you want and even migrate them,
without breaking the importpath.

## What can it do

- turn `github.com/your-username/some-obscure-wordplay-with-go` into
  `your.domain.tld/some-project`
- allow you to generate shields that automatically display the latest version:
  ![example](https://img.shields.io/badge/dynamic/json?color=green&label=Version&query=%24.latest_version&url=https%3A%2F%2Fmpldr.codes%2Frecter%2Fapi%2Fversions%2Flatest&style=flat-square&logo=git&color=F05032)
- setup a portal listing all your projects
- allow godoc renderers to link to your sourcefile

## What can't it do, and why.

- create multi-element paths
	- one of the very ideas was to simplify import paths. Having
	  multi-element-paths runs contrary to this idea.
- detect *anything* from your git repo (like the default branch) that is not
  exposed somewhere else
	- versions can be retrieved from Go Proxies, so these are easy.
	  Reimplementing `git ls-remote` is considered out of scope.
- generate SSL certificates
	- nginx and caddy are very mature reverse proxies and can be configured
	  way better and more detailed. This is to encourage the use of a
	  reverse proxy.

## How to deploy it

see [deployment](./deployment.md)
