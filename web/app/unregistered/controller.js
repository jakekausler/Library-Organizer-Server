var app = angular.module('unregistered', ['ngMaterial']).config(function($mdThemingProvider) {
	$mdThemingProvider.theme('default')
		.primaryPalette('indigo')
		.accentPalette('indigo')
		.warnPalette('red')
		.backgroundPalette('indigo');
});
app.controller('unregisteredController', function($scope, $http, $mdDialog) {
	$scope.login = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			controller: function ($scope, $mdDialog) {
				$scope.cancel = function() {
					$mdDialog.cancel()
				};
			},
			templateUrl: 'web/app/unregistered/login.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
	$scope.register = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			controller: function ($scope, $mdDialog) {
				$scope.cancel = function() {
					$mdDialog.cancel()
				};
			},
			templateUrl: 'web/app/unregistered/register.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
});