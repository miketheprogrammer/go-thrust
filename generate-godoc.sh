mkdir -p doc
rm -rf doc/*
mkdir doc/commands
mkdir doc/common
mkdir doc/menu
mkdir doc/window
mkdir doc/spawn
godocdown commands > doc/commands/README.md
godocdown common > doc/common/README.md
godocdown menu > doc/menu/README.md
godocdown window > doc/window/README.md
godocdown spawn > doc/spawn/README.md

