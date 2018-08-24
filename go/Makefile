REGION=ap-northeast-1
PROFILE=gemcook

.PHONY: build

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/publish_all handlers/publish_all/main.go

.PHONY: clean

clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy-local

deploy-local: clean build
	sls deploy --stage local --aws-profile $(PROFILE) --region $(REGION)

.PHONY: deploy-develop

deploy-develop: clean build
	sls deploy --stage develop --aws-profile $(PROFILE) --region $(REGION)

.PHONY: deploy-release

deploy-release: clean build
	sls deploy --stage release --aws-profile $(PROFILE) --region $(REGION)

.PHONY: deploy-master

deploy-master: clean build
	sls deploy --stage master --aws-profile $(PROFILE) --region $(REGION)