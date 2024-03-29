all: binary

clean:
	go clean -v

binary: clean
	GOARM=5 GOARCH=arm go build -v

copy: binary
	rsync -rvutpP cadpublish cadpublish.yaml jbuchbinder@manage:
