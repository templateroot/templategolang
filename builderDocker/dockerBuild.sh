cd ..
CGO_ENABLED=0 GOOS=linux  go build  -ldflags "-w -s" -buildvcs=false -o builderDocker/
cd builderDocker
sudo cp /etc/ssl/certs/ca-certificates.crt ./
sudo chmod 666 ca-certificates.crt
cp ../WJGOTEMPLATE1.cfg ./
docker build --tag template1 .
#docker build --tag dockerid/template1 .
rm ca-certificates.crt
rm wjgotemplate1 WJGOTEMPLATE1.cfg

# remove all [none] tag images
docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi --force
