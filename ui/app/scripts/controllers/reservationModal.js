'use strict';

/**
 * @ngdoc function
 * @name app.controller:ReservationModalCtrl
 * @description
 * # ReservationModalCtrl
 */
angular.module('app')
  .controller('ReservationModalCtrl', function ($uibModalInstance, calendarEvent) {
  var $rmCtrl = this;
  $rmCtrl.calendarEvent = calendarEvent;

  $rmCtrl.save = function () {
    $uibModalInstance.close();
  };

  $rmCtrl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
