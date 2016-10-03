'use strict';

module.exports = function(app, cb) {
  var delay = 5000;

  var timeoutId = setInterval(function() {
    app.models.Updater.checkForUpdates(function() {});
  }, delay);

  /*
   * The `app` object provides access to a variety of LoopBack resources such as
   * models (e.g. `app.models.YourModelName`) or data sources (e.g.
   * `app.datasources.YourDataSource`). See
   * http://docs.strongloop.com/display/public/LB/Working+with+LoopBack+objects
   * for more info.
   */
  process.nextTick(cb); // Remove if you pass `cb` to an async function yourself
};
