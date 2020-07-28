# rohandvivedi.com
This is a web application hosted on rohandvivedi.com,
sole purpose : showcase projects and act as internet hosted resume

to setup :
 * ``sudo -E bash ./setup.sh &``

to deploy :
 * ``sudo -E bash ./deploy.sh &``

*Note: you may replace **bash** with **sh***

to access DB in console:
 * you need sqlite3 shell ``sudo apt-get install sqlite3``
 * ***you do not need to install sqlite3 library for deploying, go will get its own binaries, while setup.***
 * ``sqlite3 ./db/data.db``