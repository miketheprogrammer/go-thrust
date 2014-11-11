tutorials:
	rm -f tutorial/basic_browser
	go build -o tutorial/bin/basic_browser tutorial/basic_browser.go
	go build -o tutorial/bin/basic_menu tutorial/basic_menu.go
	go build -o tutorial/bin/basic_menu_events tutorial/basic_menu_events.go


dist: build.release
	mkdir -p dist
	cd release && zip -r ../dist/go-thrust-v0.1.0-x64.zip go-thrust/*