.PHONY: test build-android

test:
	go test -race -v

build-android:
	gomobile bind -target android -o age.aar filippo.io/age github.com/MarinX/agemobile