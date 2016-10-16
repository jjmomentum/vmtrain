'use strict';

/**
 * @ngdoc function
 * @name app.controller:NavCtrl
 * @description
 * # NavCtrl
 */
angular.module('app')
  .controller('NavCtrl', function ($scope, $location) {
    $scope.isActive = function (viewLocation) {
      return viewLocation === $location.path();
    };
  });
