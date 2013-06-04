#### Example of a less Go server on heroku with a proper Go workspace

<pre>
$ heroku create <example_app_name>
$ heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go 
$ git push heroku master
</pre>

#### This example compiles on heroku

The directory structure has two binaries, and a `foobar.com/shared` shared package that both to import.

    src/foobar.com/frontend/  <-- there's a binary in here, the web one
    src/foobar.com/shared/    <-- NO BINARY in here, just a package with shared code
    src/foobar.com/worker/    <-- there's a second binary in here

.godir must be set to base path. In this case, foobar.com

Symlinks are used to expose the frontend, shared, and worker packages to the
heroku buildpack.
