GIT_TAG := $(shell git describe --tags --abbrev=0)
test_target := "./..."

.PHONY: test-coverage
test-coverage:
	go test -cover -coverprofile=coverage.out ./...; go tool cover -html=coverage.out -o coverage.html; rm coverage.out

# https://github.com/vektra/mockery is needed
.PHONY: create-mocks
create-mocks:
	mockery --all

# https://github.com/mfridman/tparse is needed
.PHONY: test-all
test-all:
	go test $(test_target) -cover -json | tparse -all

# https://github.com/mfridman/tparse is needed
.PHONY: test-unit
test-unit:
	go test --short $(test_target) -cover -json | tparse -all

.PHONY: publish-package
publish-package:
	GOPROXY=proxy.golang.org go list -m github.com/KarnerTh/mermerd@$(GIT_TAG)

erd: 
	find . -name "*.mmd" -exec mmdc -s 5 -i {} -o {}.png \;

push_aws_image: 
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 522507079211.dkr.ecr.us-east-1.amazonaws.com
	docker build -t mermerd --target base .
	docker tag mermerd:latest 522507079211.dkr.ecr.us-east-1.amazonaws.com/mermerd:latest
	docker push 522507079211.dkr.ecr.us-east-1.amazonaws.com/mermerd:latest
