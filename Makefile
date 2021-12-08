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
	$(MAKE) build OUTFILE=.docker/recter GOTAGS="-tags docker" CGO_ENABLED=0
	mkdir -p .docker/data .docker/etc/ssl/certs
	cp -rf themes/ .docker/
	docker run --rm -v $(PWD)/.docker:/data alpine /bin/ash -c "apk --no-cache add ca-certificates && cp /etc/ssl/cert.pem /data/etc/ssl/ && cp /etc/ssl/certs/ca-certificates.crt /data/etc/ssl/certs"
	docker build -t ${IMAGEFULLNAME} .docker

push:
	docker tag ${IMAGEFULLNAME} ${REPO}/${IMAGENAME}:latest
	docker push ${IMAGEFULLNAME}
	docker push ${REPO}/${IMAGENAME}:latest

clean:
	rm recter
