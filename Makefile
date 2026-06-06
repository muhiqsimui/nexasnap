.PHONY: dev generator sync-api build clean

dev:
	cd astro && npm run dev

generator:
	cd generator && go run ./cmd

sync-api:
	cd astro && npm run sync-api

build: generator sync-api
	cd astro && npm run build

clean:
	rm -rf generated/api/v1/*.json generated/docs/*.json
