angular.module('unregistered', ['ngMaterial'])
.config(function($mdThemingProvider) {
	$mdThemingProvider.theme('default')
		.primaryPalette('indigo')
		.accentPalette('indigo')
		.warnPalette('red')
		.backgroundPalette('indigo');
})
.controller('unregisteredController', function($scope, $http, $mdDialog) {
	$scope.login = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			controller: 'loginController',
			templateUrl: 'web/app/unregistered/login/login.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
	$scope.register = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			controller: 'registerController',
			templateUrl: 'web/app/unregistered/register/register.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
});