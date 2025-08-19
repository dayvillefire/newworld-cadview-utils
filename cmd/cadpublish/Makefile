all: binary

clean:
	go clean -v

binary: clean
	go build -v

docker: binary
	docker build -t jbuchbinder/cadpublish .
	docker push jbuchbinder/cadpublish

