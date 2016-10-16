'use strict';

/**
 * @ngdoc function
 * @name app.controller:UserModalCtrl
 * @description
 * # UserModalCtrl
 */
angular.module('app')
  .controller('UserModalCtrl', function ($uibModalInstance, UsersModel, user) {
  var $umCtrl = this;
  $umCtrl.user = user;

  $umCtrl.save = function () {
    UsersModel.create($umCtrl.user);
    $uibModalInstance.close();
  };

  $umCtrl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
