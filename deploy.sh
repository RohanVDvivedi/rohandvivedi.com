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

# we own the Fucking server, kill any other application that is competing for 80 or 443 port

# kill the application that is running currently on port 80
APP_PIDS_RUNNING_ON_PORT_80=`sudo lsof -t -i tcp:80`
if [ $APP_PIDS_RUNNING_ON_PORT_80 ] 
then
    echo "sigkill on pids : " $APP_PIDS_RUNNING_ON_PORT_80
    kill -9 $APP_PIDS_RUNNING_ON_PORT_80
else
    echo "no process to be killed to acquire port tcp:80"
fi

# kill the application that is running currently on port 443
APP_PIDS_RUNNING_ON_PORT_443=`sudo lsof -t -i tcp:443`
if [ $APP_PIDS_RUNNING_ON_PORT_443 ] 
then
    echo "sigkill on pids : " $APP_PIDS_RUNNING_ON_PORT_443
    kill -9 $APP_PIDS_RUNNING_ON_PORT_443
else
    echo "no process to be killed to acquire port tcp:443"
fi

# run the application again
bash ./run.sh "$@"