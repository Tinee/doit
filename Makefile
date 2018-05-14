.PHONY: 

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

# dtest:
# 	docker build -t dockertest -f ./docker/authsvc/test.dockerfile .
# 	docker run dockertest

dbuild:
	docker build -t go-messenger/authsvc:${TAG} -f ./docker/authsvc/prod.dockerfile .

dtag:
	docker tag go-messenger/authsvc:${TAG} 743089793964.dkr.ecr.eu-west-2.amazonaws.com/go-messenger/authsvc:${TAG}

dpush:
	docker push 743089793964.dkr.ecr.eu-west-2.amazonaws.com/go-messenger/authsvc:${TAG}

ship: dbuild dtag dpush
	