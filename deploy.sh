# go the application directory
cd ~/go/src/rohandvivedi.com

# checkout master branch
git checkout master

# pull new changes on master
git pull

# install for newly added go packages 
go get github.com/mattn/go-sqlite3

# install for newly added node packages 
npm install --no-optional

# kill the application that is running currently on port 80
APP_PIDS_RUNNING_ON_PORT_80=`sudo lsof -t -i tcp:80`
if [ $APP_PIDS_RUNNING_ON_PORT_80 ] 
then
    echo "sigkill on pids : " $APP_PIDS_RUNNING_ON_PORT_80
    kill -9 $APP_PIDS_RUNNING_ON_PORT_80
else
    echo "no process to be killed to acquire port tcp:80"
fi

# run the application again
sh ./run.sh