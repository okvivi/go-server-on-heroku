#### Example of a less trivial Go server on heroku

<pre>
$ heroku create <example_app_name>
$ heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go 
$ git subtree push --prefix src heroku master
</pre>

#### A few non trivial observations

The .godir directory is meant to point not only to where you want the app to run, but also
to where you want it to compile. Think of it as the root of your 'src/' tree. All the code
under .godir will be copied under $GOPATH/src.

I'm using `git subtree` to push code to heroku because this is also a scenario with a more
complicated repository, not the standard 'one repo per binary' but instead something where
you can keep more than just your go files.
