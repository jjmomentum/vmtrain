// Copyright IBM Corp. 2015. All Rights Reserved.
// Node module: loopback-getting-started-intermediate
// This file is licensed under the MIT License.
// License text available at https://opensource.org/licenses/MIT

angular
  .module('app')
  .controller('AuthLoginController', ['$scope', 'AuthService', '$state',
      function($scope, AuthService, $state) {
    $scope.user = {
      email: 'admin@vmware.com',
      password: 'adminpassword'
    };

    $scope.login = function() {
      AuthService.login($scope.user.email, $scope.user.password)
        .then(function() {
          $state.go('all-servers');
        });
    };
  }])
  .controller('AuthLogoutController', ['$scope', 'AuthService', '$state',
      function($scope, AuthService, $state) {
    AuthService.logout()
      .then(function() {
        $state.go('all-servers');
      });
  }]);
