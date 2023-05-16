
#set -x

cd ../

curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh


./install-fabric.sh --fabric-version 2.5.0 d s b
#curl -sSL https://bit.ly/2ysbOFE | bash -s 2.5.1

#cd fabric-samples
#git apply fabric.patch
