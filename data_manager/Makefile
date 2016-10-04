#
all: docker

q3-training-journal: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

docker: q3-training-journal
	docker build -t q3-training-journal --rm=true .

clean:
	go clean
	rm -f q3-training-journal
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/q3-training-journal/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/q3-training-journal/ { print $$1}') || echo -
	docker rmi q3-training-journal

run:
	docker run -d -p 8080:8080 q3-training-journal /q3-training-journal -l 8080 -t .

stop:
	killall q3-training-journal

.PHONY: q3-training-journal docker clean run stop
