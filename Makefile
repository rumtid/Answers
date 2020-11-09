test:
	for i in */; do                           \
		cd $$i;                               \
		gofmt -w *.go && go vet && go test .; \
		cd ..;                                \
	done
