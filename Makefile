.PHONY: build-for-raspberry-pi
build-for-raspberry-pi:
	GOOS=linux GOARCH=arm GOARM=6 go build -tags netgo -installsuffix netgo -ldflags '-extldflags "-static"' -o build/go-moneyforward ./cmd/

.PHONY: clean
clean:
	$(RM)	./build/*
