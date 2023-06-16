.PHONY: test build-android

test:
	go test -race -v

build-android:
	@mkdir -p build/android
	gomobile bind -target android -androidapi 21 -o build/android/age.aar filippo.io/age github.com/MarinX/agemobile

build-ios:
	@mkdir -p build/ios
	gomobile bind -target ios -o build/ios/Age.xcframework filippo.io/age github.com/MarinX/agemobile
