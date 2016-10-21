'use strict';

/**
 * @ngdoc function
 * @name app.service:ServersModel
 * @description
 * # ServersModel
 */

angular.module('app')
  .service('ServersModel', function ($http, $route, ENDPOINT_URI) {
    var service = this,
    path = 'servers/';

    function getUrl() {
      return ENDPOINT_URI + path;
    }

    function getUrlForId(userId) {
      return getUrl(path) + userId;
    }

    service.all = function () {
      return $http.get(getUrl());
    };

    service.fetch = function (serverId) {
      return $http.get(getUrlForId(serverId));
    };

    service.create = function (server) {
      return $http.post(getUrl(), server).then(function successCallback() {
        $route.reload();
      });
    };

    service.update = function (serverId, server) {
      return $http.put(getUrlForId(serverId), server).then(function successCallback() {
        $route.reload();
      });
    };

    service.destroy = function (serverId) {
      return $http.delete(getUrlForId(serverId)).then(function successCallback() {
        $route.reload();
      });
    };
});
