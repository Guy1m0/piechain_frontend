
set -x

cd ../


curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.2 1.5.2

cd fabric-samples
git apply ../fabric/fabric.patch
