.PHONY: build/fkot
build/fkot: go build -tags netgo -ldflags '-extldflags "-static"' -o ./moneyforward ./cmd/
	docker build -t moneyforward:latest .

.PHONY: run
run: docker run --name moneyforward --rm -d -it moneyforward:latest

.PHONY: stop
stop:
	docker stop moneyforward

.PHONY: clean
clean:
	$(RM)	./moneyforward
