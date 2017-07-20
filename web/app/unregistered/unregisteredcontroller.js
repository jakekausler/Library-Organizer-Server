angular.module('unregistered', ['ngMaterial'])
.config(function($mdThemingProvider) {
	$mdThemingProvider.theme('default')
		.primaryPalette('indigo')
		.accentPalette('indigo')
		.warnPalette('red')
		.backgroundPalette('indigo');
})
.controller('unregisteredController', function($scope, $mdDialog) {
	$scope.login = function(ev) {
		$mdDialog.show({
			controller: 'logincontroller',
			templateUrl: 'web/app/unregistered/login/login.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false,
			locals: {
				vm: $scope
			}
		});
	};
	$scope.register = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			templateUrl: 'web/app/unregistered/register/register.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
	$scope.resetPassword = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			templateUrl: 'web/app/unregistered/reset/reset.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	};
});