tutorials:
	rm -rf tutorial/bin/*
	mkdir -p tutorial/bin
	# go build -o tutorial/bin/basic_window tutorial/basic_window.go
	# go build -o tutorial/bin/basic_menu tutorial/basic_menu.go
	# go build -o tutorial/bin/basic_menu_events tutorial/basic_menu_events.go
	# go build -o tutorial/bin/basic_session tutorial/basic_session.go
	# go build -o tutorial/bin/basic_multiple_windows tutorial/basic_multiple_windows.go
	go build -o tutorial/bin/advanced_session tutorial/advanced_session.go


dist: build.release
	mkdir -p dist
	cd release && zip -r ../dist/go-thrust-v0.1.0-x64.zip go-thrust/*