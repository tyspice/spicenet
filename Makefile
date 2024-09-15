
build:
	@templ generate
	@npx tailwindcss -i ./assets/styles/input.css -o ./public/assets/styles.css
	@go build -o ./bin/spicenet main.go

tailwind:
	@npx tailwindcss -i ./assets/styles/input.css -o ./public/assets/styles.css --watch

templ:
	@templ generate --watch

clean:
	rm -rf ./bin