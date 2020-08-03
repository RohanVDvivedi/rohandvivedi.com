# rohandvivedi.com
This is a web application hosted on rohandvivedi.com,
sole purpose : showcase projects and act as internet hosted resume
I plan to destribute this app in future, and so I am designing it with a good backend.

to setup :
 * ``sudo -E bash ./setup.sh &``

to deploy :
 * ``sudo -E bash ./deploy.sh &``

*Note: you may replace **bash** with **sh***

to access DB in console:
 * you need sqlite3 shell ``sudo apt-get install sqlite3``
 * ***you do not need to install sqlite3 library for deploying, go will get its own binaries, while setup.***
 * ``sqlite3 ./db/data.db``

The Schema is public
Schema :
 * table: people (this mostly would carry only 1 entry of the owner of this project)
   * id
   * fname
   * lname
   * email
   * ph_no
   * linkedin
   * github
   * type ("owner" = Rohan, there can be only 1 owner)
 * table: projects
   * id
   * name
   * descr
   * type (embedded system, systems programming, robotics, etc)
 * table: project_hyperlinks
   * id
   * href
   * descr
   * type (must be youtube, pdf_document, github, technopedia, this helps in deciding the)
   * project_id
 * table: person_project
   * person_id
   * project_id 
