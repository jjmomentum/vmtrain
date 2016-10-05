'use strict';

/**
 * @ngdoc function
 * @name app.controller:DashboardCtrl
 * @description
 * # DashboardCtrl
 * Controller of the app
 */
angular.module('app')
  .controller('DashboardCtrl', function ($scope) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    $scope.eventSources = [];
  });
