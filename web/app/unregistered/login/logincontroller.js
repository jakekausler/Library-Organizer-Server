angular.module('unregistered')
.controller('loginController', function($scope, $mdDialog) {
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
});