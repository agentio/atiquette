local:
	go install ./...

lexgen:
	lexgenlite --build-file cmd/lexgenlite/build/at.json lexicons
	lexgenlite --build-file cmd/lexgenlite/build/xyz.json lexicons
	lexgenlite --max-string-length 1000000 --build-file cmd/lexgenlite/build/bsky.json lexicons
