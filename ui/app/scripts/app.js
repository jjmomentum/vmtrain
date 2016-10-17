'use strict';

/**
 * @ngdoc overview
 * @name app
 * @description
 * # app
 *
 * Main module of the application.
 */
angular
  .module('app', [
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'mwl.calendar',
    'ui.bootstrap'
  ])
  .constant('ENDPOINT_URI', 'http://localhost:9292/api/1/')
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/dashboard.html',
        controller: 'DashboardCtrl',
        controllerAs: 'dashboard'
      })
      .when('/users', {
        templateUrl: 'views/users.html',
        controller: 'UsersCtrl',
        controllerAs: 'users'
      })
      .when('/servers', {
        templateUrl: 'views/servers.html',
        controller: 'ServersCtrl',
        controllerAs: 'servers'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
