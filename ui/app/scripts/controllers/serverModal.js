'use strict';

/**
 * @ngdoc function
 * @name app.controller:ServerModalCtrl
 * @description
 * # ServerModalCtrl
 */
angular.module('app')
  .controller('ServerModalCtrl', function ($uibModalInstance, ServersModel, server) {
  var $smCtrl = this;
  $smCtrl.server = server;

  $smCtrl.save = function () {
    ServersModel.create($smCtrl.server);
    $uibModalInstance.close();
  };

  $smCtrl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
