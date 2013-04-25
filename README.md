#### Example of a less trivial Go server on heroku

<pre>
$ heroku create <example_app_name>
$ heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go 
$ git push heroku master
</pre>

#### This example does not currently compile on heroku

The directory structure has two binaries, and a `foobar.com/shared` shared package that both try to import.

    foobar.com/frontend/  <-- there's a binary in here, the web one
    foobar.com/shared/    <-- NO BINARY in here, just a package with shared code
    foobar.com/worker/    <-- there's a second binary in here

This directory structure no longer compiles, because the moment I specify a ".godir" that's anything other than the root of my directory structure, the go compiler can no longer find the code under foobar.com/shared anymore.

