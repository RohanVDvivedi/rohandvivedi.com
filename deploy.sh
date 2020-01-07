# checkout master branch
git checkout master

# pull new changes on master
git pull

# kill the application that is running currently on port 80
kill -9 $(lsof -t -i:80)

# install for newly added go packages 
# go get

# install for newly added node packages 
npm install

# run the application again
sh ./run.sh