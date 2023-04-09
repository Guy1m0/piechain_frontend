
set -x

cd ../


curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.5.0-beta

#cd fabric-samples
git apply ../fabric/fabric.patch
