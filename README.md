# Explore Strava routes

Browse routes on a map in a web page.

![Strava API Explorer](https://raw.githubusercontent.com/shavit/cautious-spork/master/doc/screenshot.png)

## Installation
1. Make sure you have `$GOPATH` configured in your machine.
2. Go to `$GOPATH/src/github.com/robfig`.
3. Download revel `git clone https://github.com/robfig/revel`.
4. Build it `cd $GOPATH/src/github.com/robfig/revel/revel`, then `go build`.
5. Install it `go install`.
6. Type `revel` and you should see a response.

Set the environment vailables in the `.profile` file. Run the app with `revel run maps`

This is a HTML5 application, built with [Brunch](http://brunch.io).

## Getting started
* Install (if you don't have them):
    * [Node.js](http://nodejs.org): `brew install node` on OS X
    * [Brunch](http://brunch.io): `npm install -g brunch`
    * Brunch plugins and app dependencies: `npm install`
* Run:
    * `brunch watch --server` — watches the project with continuous rebuild. This will also launch HTTP server with [pushState](https://developer.mozilla.org/en-US/docs/Web/Guide/API/DOM/Manipulating_the_browser_history).
    * `brunch build --production` — builds minified project for production
* Learn:
    * `public/` dir is fully auto-generated and served by HTTP server.  Write your code in `app/` dir.
    * Place static files you want to be copied from `app/assets/` to `public/`.
    * [Brunch site](http://brunch.io), [Getting started guide](https://github.com/brunch/brunch-guide#readme)
