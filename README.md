# rohandvivedi.com
This is a web application hosted on rohandvivedi.com,
sole purpose : showcase projects and act as internet hosted resume
I plan to destribute this app in future, and so I am designing it with a good backend.

#### to setup :
 * ``sudo -E ./setup.sh &``
 * for production use  ``sudo -E ./deploy.sh prod &``

#### to deploy :
 * ``sudo -E ./deploy.sh &``
 * for production use  ``sudo -E ./deploy.sh prod &``

#### to access DB in console:
 * you need sqlite3 shell ``sudo apt-get install sqlite3``
 * ***you do not need to install sqlite3 library for deploying, go will get its own binaries, while setup.***
 * ``sqlite3 ./db/data.db``

#### to delete the database
 * ``rm ./db/data.db``

#### to change application configurations
 * use `dev_config.json` and `prod_config.json` according to the requirement, for your environment

#### to setup new owner
 * delete the database file
 * update the owner.json file
 * deploy
 * then use sqlite3 shell, to insert rest of the data (I know, this will become easier after the admin panel is developed)

#### Schema is public and defined as follows :
 * table: persons
   * id
   * fname
   * lname
   * email
   * ph_no
   * type ("owner" = Rohan, there can be only 1 owner)
 * table: socials
   * id
   * descr
   * username  (username corresponding to each profile listed)
   * profile_link
   * link_type (must be youtube, pdf, github, linkedin, facebook, this helps how to interpret the link)
   * person_id
 * table: pasts
   * id
   * organization
   * organization_link
   * past_type
   * position
   * team_or_research_title
   * descr
   * research_paper
   * from_date
   * to_date
   * person_id
 * table: projects
   * id
   * name
   * descr
   * progr_lang
   * libs_used
   * skill_set
   * project_owner (refers to person_id)
 * table: project_hyperlinks
   * id
   * name
   * href
   * link_type (must be caps for youtube, pdf, github, technopedia, this helps how to interpret it)
   * descr
   * project_id
 * table: project_categories
   * id
   * category_name
   * descr
 * table: project_category_project
   * project_category_id
   * project_id
 * table: person_project
   * person_id
   * project_id

