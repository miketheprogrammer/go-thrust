mkdir -p doc
godoc -html=true github.com/miketheprogrammer/go-thrust/commands > doc/commands.html
godoc -html=true github.com/miketheprogrammer/go-thrust/common > doc/common.html
godoc -html=true github.com/miketheprogrammer/go-thrust/menu > doc/menu.html
godoc -html=true github.com/miketheprogrammer/go-thrust/window > doc/window.html
godoc -html=true github.com/miketheprogrammer/go-thrust/spawn > doc/spawn.html

