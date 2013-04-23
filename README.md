#### Example Go server on heroku

`$ heroku create <example_app_name>
$ heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go 
$ git push heroku master
`

The error I get from this is
`$ git push heroku master                                                                                                                                        [15:31:00]
Counting objects: 14, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (8/8), done.
Writing objects: 100% (14/14), 1.25 KiB, done.
Total 14 (delta 0), reused 3 (delta 0)

-----> Fetching custom git buildpack... done
-----> Go app detected
-----> Installing Go 1.0.3... done
       Installing Virtualenv... done
       Installing Mercurial... done
       Installing Bazaar... done
-----> Running: go get -tags heroku ./...
package okvivi/models: unrecognized import path "okvivi/models"
 !     Heroku push rejected, failed to compile Go app

To git@heroku.com:gofe-example.git
 ! [remote rejected] master -> master (pre-receive hook declined)
error: failed to push some refs to 'git@heroku.com:gofe-example.git'`

