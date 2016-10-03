// Copyright IBM Corp. 2014,2015. All Rights Reserved.
// Node module: loopback-getting-started-intermediate
// This file is licensed under the MIT License.
// License text available at https://opensource.org/licenses/MIT
'use strict';

var async = require('async');

module.exports = function(app) {
  var env = (process.env.NODE_ENV || 'development');
  var isDevEnv = env === 'development' || env === 'test';
  if (!isDevEnv) {
    return;
  }

  // create all models
  async.parallel({
    users: async.apply(createUsers),
    servers: async.apply(createServers),
    updates: async.apply(createUpdater),
  }, function(err, results) {
    if (err) throw err;

    console.log('> dev/test models created successfully');
  });

  // create system users
  function createUsers(cb) {
    app.models.User.create([
      {email: 'admin@vmware.com', password: 'adminpassword'},
    ], cb);
  }

  // create servers
  function createServers(cb) {
    app.models.Server.create([
      {name: 'server1'},
      {name: 'server2'},
      {name: 'server3'},
    ], cb);
  }

  // create an updater entry
  function createUpdater(cb) {
    app.models.Updater.create([
      {active: true},
    ], cb);
  }
};
