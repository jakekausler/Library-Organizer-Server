angular.module('libraryOrganizer')
.controller('exportController', function ($scope, $mdDialog, $http) {
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
	$scope.export = function() {
		$http({
			url: '/books/books',
			method: 'GET'
		}).then(function(data) {
			var anchor = angular.element('<a/>');
			anchor.attr({
				href: 'data:attachment/csv;charset=utf-8,' + encodeURI(data.data),
				target: '_blank',
				download: 'books.csv'
			})[0].click();
		});
		$http({
			url: '/books/contributors',
			method: 'GET'
		}).then(function(data) {
			var anchor = angular.element('<a/>');
			anchor.attr({
				href: 'data:attachment/csv;charset=utf-8,' + encodeURI(data.data),
				target: '_blank',
				download: 'authors.csv'
			})[0].click();
		});
		$mdDialog.cancel()
	};
});