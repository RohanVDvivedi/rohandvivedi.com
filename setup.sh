# update and upgrade system
apt-get update
apt-get upgrade

# setup golang

# setup nodejs and npm
apt-get install nodejs
apt-get install npm

# go to the appropriate parent directory
cd ~/go/src

# clone the repository
git clone https://github.com/RohanVDvivedi/rohandvivedi.com.git

# enter directory
cd ~/go/src/rohandvivedi.com

# run the application
sh ./deploy.sh