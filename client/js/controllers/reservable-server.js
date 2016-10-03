// Copyright IBM Corp. 2015. All Rights Reserved.
// Node module: loopback-getting-started-intermediate
// This file is licensed under the MIT License.
// License text available at https://opensource.org/licenses/MIT

angular
  .module('app')
  .controller('AllServersController', ['$scope', 'Server', function($scope,
      Server) {
    $scope.servers = Server.find({
    });
  }]);
