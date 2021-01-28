#!bin/bash

heroku container:login
heroku container:push web -a threeo1
heroku container:release web -a threeo1