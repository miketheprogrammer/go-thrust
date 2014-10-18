build.release:
	rm -rf release/vendor/darwin/x64/*
	touch release/vendor/darwin/x64/README.md
	rm -rf release/vendor/linux/x64/*
	touch release/vendor/linux/x64/README.md
	cp -rf tools release/
	rm -f release/Thrust
	go build -o release/Thrust
