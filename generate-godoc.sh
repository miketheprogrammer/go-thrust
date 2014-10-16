mkdir -p doc
godoc -html=true github.com/miketheprogrammer/thrust-go/commands > doc/commands.html
godoc -html=true github.com/miketheprogrammer/thrust-go/common > doc/common.html
godoc -html=true github.com/miketheprogrammer/thrust-go/menu > doc/menu.html
godoc -html=true github.com/miketheprogrammer/thrust-go/window > doc/window.html
godoc -html=true github.com/miketheprogrammer/thrust-go/spawn > doc/spawn.html

