'use strict';

/**
 * @ngdoc function
 * @name app.service:ReservationsModel
 * @description
 * # ReservationsModel
 */

angular.module('app')
  .service('ReservationsModel', function ($http, ENDPOINT_URI) {
    var service = this,
    path = 'reservations/';

    function getUrl() {
      return ENDPOINT_URI + path;
    }

    function getUrlForId(reservationId) {
      return getUrl(path) + reservationId;
    }

    service.all = function () {
      return $http.get(getUrl());
    };

    service.fetch = function (reservationId) {
      return $http.get(getUrlForId(reservationId));
    };

    service.create = function (reservation) {
      return $http.post(getUrl(), reservation);
    };

    service.update = function (reservationId, reservation) {
      return $http.put(getUrlForId(reservationId), reservation);
    };

    service.destroy = function (reservationId) {
      return $http.delete(getUrlForId(reservationId));
    };
});
