REGION=ap-northeast-1
PROFILE=gemcook

.PHONY: deploy-manual

deploy-manual:
	sls deploy --stage $(OPT_BRANCH) --region $(REGION) --aws-profile $(PROFILE)


.PHONY: deploy

deploy:
	sls deploy --stage $(OPT_BRANCH) --region $(REGION)
