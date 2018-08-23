REGION=ap-northeast-1
PROFILE=gemcook

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/notificate_all handlers/notificate_all/main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy-local
deploy-local: clean build
	sls deploy --stage local --aws-profile $(PROFILE) --region $(REGION)