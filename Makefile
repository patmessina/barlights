APP=barlights
PLATFORM=linux/arm/v6
BUILD_CONTAINER_NAME=ws2811-builder:v6

rpi:
	docker run --rm -v "${PWD}":/usr/src/$(APP) --platform linux/arm/v6 \
		-w /usr/src/$(APP) ws2811-builder:v6 go build -o "bin/$(APP)-armv6" -v

docker-build-container:
	docker buildx build --platform $(PLATFORM) --tag $(BUILD_CONTAINER_NAME) --file ./Dockerfile .


clean:
	go clean
	rm -rf bin/
