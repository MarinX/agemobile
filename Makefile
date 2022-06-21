.PHONY: test build-android

test:
	go test -race -v

build-android:
	gomobile bind -target android -o build/android/age.aar filippo.io/age github.com/MarinX/agemobile

build-ios:
	gomobile bind -target ios -o build/ios/Age.xcframework filippo.io/age github.com/MarinX/agemobile