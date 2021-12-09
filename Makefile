VERSION:=$(shell git describe)
GOEXE:=$(shell go env GOEXE)
OUTFILE?=recter$(GOEXE)
CGO_ENABLED?=0 

build:
	go build $(GOTAGS) -o "$(OUTFILE)" -v -trimpath -modcacherw -mod=readonly -ldflags="-s -w -X main.Version=$(VERSION)"

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

run-goss-test: deps build
	cp contrib/config.toml config.toml
	./recter &> /dev/null &
	@sleep 1
	tools/goss$(GOEXE) -g contrib/goss.yml validate -r 30s
	@pkill -9 recter
	@sleep 1

deps:
	go build -o tools github.com/aelsabbahy/goss/cmd/goss

clean:
	rm recter
