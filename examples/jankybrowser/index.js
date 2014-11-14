"use strict";

var argv = require('minimist')(process.argv.slice(2));

var async = require('async');

var api = null;
var window = null;

var url = __dirname + "/test.html";

var fs = require("fs");
var http = require("http");
var port = 21024;


async.series([
  function(cb) {
    var server = http.createServer(function(req, res) {
      if(req.url === "/") {
        req.url = "/browser.html";
      }

      req.url = __dirname + req.url;
      console.log("attempting to load %s", req.url);
      fs.createReadStream(req.url).pipe(res);
    });

    server.listen(port, '127.0.0.1', function() {
      console.log("listening on %s", port);
      cb(null);
    });
  },
  function(cb_) {
    require('node-thrust')(function(err, a) {
      api = a;
      return cb_(err);
    }, argv.thrust_path || null);
  },
  function(cb_) {
    var root = "http://127.0.0.1:" + port;

    if(argv.ext) root += ("/" + argv.ext);

    window = api.window({
      root_url: root,
      size: {
        width: 1024,
        height: 768
      }
    });

    window.on("closed", function() {
      process.exit(0);
    });

    return cb_();
  },
  function(cb_) {
    window.show(cb_);
  }
], function(err) {
  if(err) {
    console.log('FAILED');
    console.error(err);
  }
  console.log('OK');
});
