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

.PHONY: deploy

deploy: clean build
	sls deploy --stage $$(OPT_BRANCH) --region $(REGION)

deploy-master: clean build
	sls deploy --stage master --region $(REGION)