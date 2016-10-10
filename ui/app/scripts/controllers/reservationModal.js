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
    console.log('save!');
    $uibModalInstance.close();
  };

  $rmCtrl.cancel = function () {
    console.log('cancel!');
    $uibModalInstance.dismiss('cancel');
  };
});
