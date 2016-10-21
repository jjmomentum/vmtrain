'use strict';

/**
 * @ngdoc function
 * @name app.controller:UsersCtrl
 * @description
 * # UsersCtrl
 */
angular.module('app')
  .controller('UsersCtrl', function (UsersModel, $uibModal) {
    var vm = this;

    UsersModel.all().then(function (result) {
      vm.users = result.data.users;
    });

    vm.openNewUserModal = function() {
      vm.userInstance = $uibModal.open({
        animation: true,
        templateUrl: '../views/partials/userModal.html',
        controller: 'UserModalCtrl',
        controllerAs: 'umCtrl',
        resolve: {
          user: function () {
            return { 'name': '' };
          }
        }
      });
    };

    vm.deleteUser = function(userId) {
      UsersModel.destroy(userId);
    };
  });
