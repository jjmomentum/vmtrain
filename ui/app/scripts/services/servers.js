'use strict';

/**
 * @ngdoc function
 * @name app.service:ServersModel
 * @description
 * # ServersModel
 */

angular.module('app')
  .service('ServersModel', function ($http, ENDPOINT_URI) {
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
      return $http.post(getUrl(), server);
    };

    service.update = function (serverId, server) {
      return $http.put(getUrlForId(serverId), server);
    };

    service.destroy = function (serverId) {
      return $http.delete(getUrlForId(serverId));
    };
});
