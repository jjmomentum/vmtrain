'use strict';
module.exports = function(Updater) {
  /**
   * Examine queue for waiting reservations and update backing store.
   * @param {Function(Error)} callback
   */
  Updater.checkForUpdates = function(callback) {
    Updater.findById(1, function(err, updater) {
      if (updater && updater.active) {
        // TODO: Read from queue any new reservation requests.
        console.log('TODO: Check reservation queue.');
      }
    });
    callback(null);
  };
};
