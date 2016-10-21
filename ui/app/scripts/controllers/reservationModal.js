'use strict';

/**
 * @ngdoc function
 * @name app.controller:ReservationModalCtrl
 * @description
 * # ReservationModalCtrl
 */
angular.module('app')
  .controller('ReservationModalCtrl', function ($uibModalInstance, ReservationsModel, calendarEvent) {
  var $rmCtrl = this;
  $rmCtrl.calendarEvent = calendarEvent;

  $rmCtrl.save = function () {

    ReservationsModel.create({
      'name': angular.element('#name').val(),
      'start_date': this.calendarEvent.startsAt.toJSON(),
      'end_date': this.calendarEvent.startsAt.toJSON(),
      'server_name': angular.element('#server').val()
    });
    $uibModalInstance.close();
  };

  $rmCtrl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
});
