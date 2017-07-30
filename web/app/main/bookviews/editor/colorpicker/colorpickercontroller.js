angular.module('libraryOrganizer')
.controller('colorpickerController', function($scope, $mdDialog, spineColor) {
	$scope.spineColor = spineColor;
	$scope.changed = false;
})