angular.module('unregistered')
.controller('registerController', function($scope, $mdDialog) {
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
});