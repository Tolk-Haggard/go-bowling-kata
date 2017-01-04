.PHONY : goconvey test deps

${GOPATH}/bin/glide:
	@go get github.com/Masterminds/glide
deps: ${GOPATH}/bin/glide
	glide install
goconvey:
	goconvey -excludedDirs "uats,vendor" -timeout 10s
test:
	go test $$(go list ./... | grep -v 'vendor\|uats')
race:
	go build -race
