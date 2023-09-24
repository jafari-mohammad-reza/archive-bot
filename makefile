build:
	@go build -o ./dist/bot
run: build
	@./dist/bot
dev:
	@/go/bin/reflex -r '\.go$$' -s -- sh -c "go build -buildvcs=false -o ./dist/bot && ./dist/bot"