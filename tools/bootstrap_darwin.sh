wget https://github.com/breach/thrust/releases/download/v$1/thrust-v$1-darwin-x64.zip
mkdir -p $2/vendor/darwin/x64/v$1
mv thrust-v$1-darwin-x64.zip $2/vendor/darwin/x64/v$1/thrust-v$1-darwin-x64.zip
cd $2/vendor/darwin/x64/v$1/
unzip thrust-v$1-darwin-x64.zip
rm thrust-v$1-darwin-x64.zip