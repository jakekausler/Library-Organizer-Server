angular.module('libraryOrganizer')
.controller('librarychooserController', function($scope, $http, $mdDialog, vm, multiselect) {
	$scope.vm = vm;
	$scope.multiselect = multiselect;
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
});