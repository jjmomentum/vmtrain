'use strict';

/**
 * @ngdoc function
 * @name app.controller:DashboardCtrl
 * @description
 * # DashboardCtrl
 * Controller of the app
 */
angular.module('app')
  .controller('DashboardCtrl', function (moment, calendarConfig, $uibModal, ServersModel, ReservationsModel) {
    var vm = this;
    vm.calendarView = 'month';
    vm.viewDate = new Date();
    vm.isCellOpen = true;
    vm.servers = [];
    vm.reservations = [];
    vm.reservationMap = {};
    vm.events =[];

    function getDaysInMonth() {
      var date = moment().startOf('month').toDate();
      var month = date.getMonth();
      var days = [];
      while (date.getMonth() === month) {
        days.push(moment(date).startOf('day').toDate());
        date.setDate(date.getDate() + 1);
      }
      return days;
    };

    ServersModel.all().then(function (result) {
      vm.servers = result.data.servers;
      ReservationsModel.all().then(function (result) {
        vm.reservations = result.data.reservations;
        // Loops reservations and build map based of day for easier access
        vm.reservationMap = {};
        angular.forEach(vm.reservations, function(reservation) {
          if (!vm.reservationMap[reservation.start_date]) {
            vm.reservationMap[reservation.start_date] = {};
          }
          vm.reservationMap[reservation.start_date][reservation.server_name] = reservation.name;
        });
        var days = getDaysInMonth();
        // Loop days to add servers
        angular.forEach(days, function(day) {
          // Loop servers to create events
          angular.forEach(vm.servers, function(server) {
            if (vm.reservationMap[day.toJSON()] && vm.reservationMap[day.toJSON()][server.name]) {
              // Add reserved server
              var reserverName = vm.reservationMap[day.toJSON()][server.name];
              vm.events.push({
                title: server.name + ' - ' + reserverName,
                color: calendarConfig.colorTypes.important,
                startsAt: moment(day).startOf('day').toDate(),
                allDay: true,
                details: {
                  name: reserverName,
                  server: server.name
                }
              });
            } else {
              // Add free server
              vm.events.push({
                title: server.name,
                color: calendarConfig.colorTypes.success,
                startsAt: moment(day).startOf('day').toDate(),
                allDay: true,
                actions: actions,
                details: {
                  name: null,
                  server: server.name
                }
              });
            }
          });
        });
      });
    });

    vm.viewChangeClicked = function() {
      // disable changing view from month
      return false;
    };

    var actions = [{
      label: '<i class="glyphicon glyphicon-pencil"></i>',
      onClick: function(args) {
        vm.openReservationModal(args.calendarEvent);
      }
    }];

    // vm.events = [
    //   {
    //     title: 'Server 1 - Tom',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Tom',
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 2 - Free',
    //     color: calendarConfig.colorTypes.success,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: null,
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 3 - Barrett',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Barrett',
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 4 - Luis',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Luis',
    //       server: 1
    //     }
    //   },
    //         {
    //     title: 'Server 5 - Free',
    //     color: calendarConfig.colorTypes.success,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: null,
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 6 - Eric',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Eric',
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 7 - Jimmy',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Jimmy',
    //       server: 1
    //     }
    //   },
    //   {
    //     title: 'Server 8 - Devin',
    //     color: calendarConfig.colorTypes.important,
    //     startsAt: moment().startOf('day').toDate(),
    //     endsAt: moment().endOf('day').toDate(),
    //     allDay: true,
    //     actions: actions,
    //     details: {
    //       name: 'Devin',
    //       server: 1
    //     }
    //   },
    //   // {
    //   //   title: 'A non all day egent',
    //   //   color: calendarConfig.colorTypes.important,
    //   //   startsAt: moment().startOf('day').add(7, 'hours').toDate(),
    //   //   endsAt: moment().startOf('day').add(19, 'hours').toDate(),
    //   //   draggable: true,
    //   //   resizable: true
    //   // }
    // ];

    vm.openReservationModal = function(calendarEvent) {
      vm.reservationInstance = $uibModal.open({
        animation: true,
        templateUrl: '../views/partials/reservationModal.html',
        controller: 'ReservationModalCtrl',
        controllerAs: 'rmCtrl',
        resolve: {
          calendarEvent: function () {
            return calendarEvent;
          }
        }
      });
    };
  });
