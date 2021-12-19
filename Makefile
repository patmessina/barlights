APP=barlights

rpi:
	docker run --rm -v "${PWD}":/usr/src/$(APP) --platform linux/arm/v6 \
		-w /usr/src/$(APP) ws2811-builder:v6 go build -o "bin/$(APP)-armv6" -v
