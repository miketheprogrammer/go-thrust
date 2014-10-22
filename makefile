build.release:
	rm -rf release/go-thrust/*
	mkdir -p release/go-thrust
	go build -o release/go-thrust/Thrust

build.tutorials:
	rm -rf tutorial/vendor/darwin/x64/*
	touch tutorial/vendor/darwin/x64/README.md
	rm -rf tutorial/vendor/linux/x64/*
	touch tutorial/vendor/linux/x64/README.md
	rm -rf tutorial/tools
	cp -rf tools tutorial
	rm -f tutorial/basic_browser
	go build -o tutorial/basic_browser tutorial/basic_browser.go


dist.darwin: build.release
	mkdir -p dist
	cd release && zip -r ../dist/go-thrust-v0.1.0-darwin-x64.zip go-thrust/*

dist.linux: build.release
	mkdir -p dist
	cd release/
	cd release && zip -r ../dist/go-thrust-v0.1.0-linux-x64.zip go-thrust/*
