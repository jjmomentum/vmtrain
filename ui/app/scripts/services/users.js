'use strict';

/**
 * @ngdoc function
 * @name app.service:UsersModel
 * @description
 * # UsersModel
 */

angular.module('app')
  .service('UsersModel', function ($http, $route, ENDPOINT_URI) {
    var service = this,
    path = 'users/';

    function getUrl() {
      return ENDPOINT_URI + path;
    }

    function getUrlForId(userId) {
      return getUrl(path) + userId;
    }

    service.all = function () {
      return $http.get(getUrl());
    };

    service.fetch = function (userId) {
      return $http.get(getUrlForId(userId));
    };

    service.create = function (user) {
      return $http.post(getUrl(), user).then(function successCallback(response) {
        $route.reload();
      });
    };

    service.update = function (userId, user) {
      return $http.put(getUrlForId(userId), user).then(function successCallback(response) {
        $route.reload();
      });
    };

    service.destroy = function (userId) {
      return $http.delete(getUrlForId(userId)).then(function successCallback(response) {
        $route.reload();
      });
    };
});
