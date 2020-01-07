kill -9 $(lsof -t -i:8080)
git checkout master
git pull
npm install
sh ./run.sh