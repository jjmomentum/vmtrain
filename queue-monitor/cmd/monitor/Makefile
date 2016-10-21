#
all: docker

monitor: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

docker: monitor
	docker build -t monitor --rm=true .

clean:
	go clean
	rm -f monitor
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/monitor/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/monitor/ { print $$1}') || echo -
	docker rmi monitor

run:
	docker run -d -p 8080:8080 monitor /monitor -l 8080 -t .

stop:
	killall monitor

.PHONY: monitor docker clean run stop 

