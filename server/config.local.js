'use strict';

var GLOBAL_CONFIG = require('./config.json');

var env = (process.env.NODE_ENV || 'development');
var isDevEnv = env === 'development' || env === 'test';
var port = process.env.PORT || GLOBAL_CONFIG.port;
module.exports = {
  hostname: GLOBAL_CONFIG.hostname,
  restApiRoot: GLOBAL_CONFIG.restApiRoot,
  livereload: process.env.LIVE_RELOAD,
  isDevEnv: isDevEnv,
  port: port,
};
