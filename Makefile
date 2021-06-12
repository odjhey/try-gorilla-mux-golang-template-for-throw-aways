PID = ./.main.pid
GO_FILES = $(wildcard *.go)
APP = ./build/main

watch: start
	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$'  . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@echo "actually do nothing"

build: $(GO_FILES)
	@go build -o $(APP)

$(APP): $(GO_FILES)
	@go build $? -o $@

start:
	# @sh -c "$(APP) & echo $$! > $(PID)"
	@./build/main & echo $$! > $(PID)

restart: kill before build start

.PHONY: start serve restart kill before # let's go to reserve rules names
