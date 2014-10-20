mkdir -p doc
mkdir -p doc/commands
mkdir -p doc/common
mkdir -p doc/menu
mkdir -p doc/window
mkdir -p doc/spawn
mkdir -p doc/dispatcher
godocdown commands > doc/commands/README.md
godocdown common > doc/common/README.md
godocdown menu > doc/menu/README.md
godocdown window > doc/window/README.md
godocdown spawn > doc/spawn/README.md
godocdown dispatcher > doc/dispatcher/README.md

# Tutorials
mkdir -p doc/tutorials/basic_browser

