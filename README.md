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

to delete the database
 * ``rm ./db/data.db``

to setup new owner
 * delete the database file
 * update the owner.json file
 * deploy
 * then use sqlite3 shell, to insert rest of the data (I know, this will become easier after the admin panel is developed)

The Schema is public
Schema :
 * table: persons (this mostly would carry only 1 entry of the owner of this project)
   * id
   * fname
   * lname
   * email
   * ph_no
   * type ("owner" = Rohan, there can be only 1 owner)
 * table: socials
   * id
   * descr
   * profile_link
   * link_type (must be youtube, pdf, github, linkedin, facebook, this helps how to interpret the link)
   * person_id
 * table: pasts
   * id
   * organization
   * organization_link
   * team
   * descr
   * from_date
   * to_date
   * person_id
 * table: projects
   * id
   * name
   * descr
   * github_link
   * youtube_link
   * image_link (additional github, youtube and images may be added in the project_hyperlinks table)
   * project_owner (refers to person_id)
 * table: project_hyperlinks
   * id
   * href
   * descr
   * link_type (must be youtube, pdf, github, technopedia, this helps how to interpret it)
   * project_id
 * project_categories
   * id
   * category_name
   * descr
 * project_category_project
   * project_category_id
   * project_id
 * table: person_project
   * person_id
   * project_id
