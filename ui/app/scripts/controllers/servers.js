'use strict';

/**
 * @ngdoc function
 * @name app.controller:ServersCtrl
 * @description
 * # ServersCtrl
 */
angular.module('app')
  .controller('ServersCtrl', function (ServersModel, $uibModal) {
    var vm = this;

    ServersModel.all().then(function (result) {
      vm.servers = result.data.servers;
    });

    vm.openNewServerModal = function() {
      vm.serverInstance = $uibModal.open({
        animation: true,
        templateUrl: '../views/partials/serverModal.html',
        controller: 'ServerModalCtrl',
        controllerAs: 'smCtrl',
        resolve: {
          server: function () {
            return { 'name': '' };
          }
        }
      });
    };

    vm.deleteServer = function(serverId) {
      ServersModel.destroy(serverId);
    };
  });
