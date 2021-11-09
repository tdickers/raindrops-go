# export DOCKER_BUILDKIT=1; make
all: bin/raindrops-go
.PHONY: bin/raindrops-go
bin/raindrops-go:
	@docker build -t raindrops-go:latest . 
run: .PHONY
	@docker run --rm --cap-add=NET_ADMIN --name raindrops-go -ti -p8080\:80 raindrops-go\:latest