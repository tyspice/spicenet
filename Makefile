.PHONY: build run tailwind templ clean

TAILWIND_WATCH = npx tailwindcss -i ./assets/styles/input.css -o ./public/assets/styles.css --watch
TEMPL_WATCH = templ generate --watch

build:
	@templ generate
	@npx tailwindcss -i ./assets/styles/input.css -o ./public/assets/styles.css
	@go build -o ./bin/spicenet main.go

run:
	@echo "Starting background processes..."
	@./run-background

tailwind:
	@$(TAILWIND_WATCH)

templ:
	@$(TEMPL_WATCH)

clean:
	@echo "Stopping background processes and cleaning binaries"
	@-kill `cat .command1.pid` 2>/dev/null || true
	@-kill `cat .command2.pid` 2>/dev/null || true
	@-kill `cat .command3.pid` 2>/dev/null || true
	@rm -f .command1.pid .command2.pid .command3.pid
	rm -rf ./bin