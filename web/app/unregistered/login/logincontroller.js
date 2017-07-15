angular.module('unregistered')
.controller('logincontroller', function($scope, vm) {
	$scope.reset = function(ev) {
		vm.resetPassword(ev);
	}
})