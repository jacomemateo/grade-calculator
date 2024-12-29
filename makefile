
build:
	tailwindcss -i ./static/css/base.css -o ./static/css/output.css
	templ generate
	go build -o tmp/main app/main.go

run:
	air