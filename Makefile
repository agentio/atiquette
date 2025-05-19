local:
	go install ./...

lexgen:
	lexgenlite --max-string-length 1000000 --build-file build/bsky.json lexicons
