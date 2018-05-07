TAG = master

all:
	$(MAKE) -j 3 \
		helloworld \
		_build/linux-amd64/helloworld.$(TAG).linux-amd64.tar.gz \
		_build/darwin-amd64/helloworld.$(TAG).darwin-amd64.tar.gz

_build/linux-amd64/helloworld.%.linux-amd64.tar.gz: _build/linux-amd64/helloworld
	cd $(dir $@) && tar czf $(notdir $@) $(notdir $<)

_build/darwin-amd64/helloworld.%.darwin-amd64.tar.gz: _build/darwin-amd64/helloworld
	cd $(dir $@) && tar czf $(notdir $@) $(notdir $<)

_build/linux-amd64/helloworld:
	mkdir -p $(dir $@)
	GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(TAG)" -o $@

_build/darwin-amd64/helloworld:
	mkdir -p $(dir $@)
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.version=$(TAG)" -o $@

helloworld:
	go build -ldflags="-X main.version=$(TAG)" -o $@

clean:
	rm -rf helloworld _build

.PHONY: all clean
