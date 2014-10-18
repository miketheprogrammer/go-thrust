wget https://github.com/breach/thrust/releases/download/v$1/thrust-v$1-darwin-x64.zip
mkdir -p ./vendor/darwin/x64/v$1
mv thrust-v$1-darwin-x64.zip ./vendor/darwin/x64/v$1/thrust-v$1-darwin-x64.zip
cd ./vendor/darwin/x64/v$1/
unzip thrust-v$1-darwin-x64.zip
rm thrust-v$1-darwin-x64.zip