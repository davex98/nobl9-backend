openapi:
	oapi-codegen -generate types -o random-generator/controllers/openapi_types.gen.go -package controllers api/swagger.yml
	oapi-codegen -generate chi-server -o random-generator/controllers/openapi_api.gen.go -package controllers api/swagger.yml
format:
	go fmt ./...

build_image:
	docker build . -t nobl9-backend --file docker/prod/Dockerfile
dev_env:
	cd docker/dev && docker-compose -f docker-compose.yml up
run_app:
	make build_image
	docker run -p 8080:8080 nobl9-backend:latest
push_image:
ifdef version
	 docker tag nobl9-backend:latest jakuburghardt/nobl9-backend:$(version) && docker push jakuburghardt/nobl9-backend:$(version)
else
	@echo 'provide version, eg: make push_image version=v5'
endif
run_app_from_repository:
ifdef version
	docker run -p 8080:8080 jakuburghardt/nobl9-backend:$(version)
else
	@echo 'specify version you wanna run, eg: make run_app_from_repository version=v5'
endif
############# TESTS #############
unit_test:
	cd random-generator && go test -race ./...
e2e_test:
	make build_image
	docker run -d -p 8080:8080 --name nobl9-backend nobl9-backend:latest
	npm install
	npm run test
	docker stop nobl9-backend
	docker container rm nobl9-backend
test_all:
	make unit_test
	make e2e_test