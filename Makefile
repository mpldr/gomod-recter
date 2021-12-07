VERSION:=$(shell git describe)
GOEXE:=$(shell go env GOEXE)
OUTFILE?=recter$(GOEXE)
CGO_ENABLED?=0 

build:
	go build $(GOTAGS) -o "$(OUTFILE)" -v -trimpath -modcacherw -mod=readonly -ldflags="-s -w -extldflags=-static"

IMAGENAME=recter
REPO=mpldr
IMAGEFULLNAME=${REPO}/${IMAGENAME}:${VERSION}
docker:
	$(MAKE) build OUTFILE=.docker/recter GOTAGS="-tags docker"
	mkdir -p .docker/data
	cp -rf themes/ .docker/
	docker build -t ${IMAGEFULLNAME} .docker

push:
	docker tag ${IMAGEFULLNAME} ${REPO}/${IMAGENAME}:latest
	docker push ${IMAGEFULLNAME} ${REPO}/${IMAGENAME}:latest

clean:
	rm recter
