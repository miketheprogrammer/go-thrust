wget https://github.com/breach/thrust/releases/download/v$1/thrust-v$1-linux-x64.zip
mkdir -p ./vendor/linux/x64/v$1
mv thrust-v$1-linux-x64.zip ./vendor/linux/x64/v$1/hrust-v$1-linux-x64.zip
cd ./vendor/linux/x64/v$1/
unzip thrust-v$1-linux-x64.zip
rm thrust-v$1-linux-x64.zip