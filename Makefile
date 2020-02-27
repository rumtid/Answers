test:
	set -eu;                                  \
	for i in */; do                           \
		pushd "$$i" > /dev/null;              \
		gofmt -w *.go && go vet && go test .; \
		popd > /dev/null;                     \
	done
