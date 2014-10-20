mkdir -p doc
rm -rf doc/*
godocdown commands > doc/commands.md
godocdown common > doc/common.md
godocdown menu > doc/menu.md
godocdown window > doc/window.md
godocdown spawn > doc/spawn.md

