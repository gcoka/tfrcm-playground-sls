PROFILE=gemcook

.PHONY: deploy-manual

deploy-manual:
	sls deploy --stage $(OPT_BRANCH) --aws-profile $(PROFILE)


.PHONY: deploy

deploy:
	sls deploy --stage $(OPT_BRANCH)
