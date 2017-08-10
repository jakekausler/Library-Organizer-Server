angular.module('libraryOrganizer')
.controller('exportController', function ($scope, $mdDialog, $http, vm) {
	$scope.vm = vm;
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
	$scope.export = function() {
        var booksloadingname = $scope.vm.guid();
        $scope.vm.addToLoading(booksloadingname)
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
            $scope.removeFromLoading(booksloadingname);
		}, function(response) {
        	$mdToast.showSimple("Failed to export books");
        	$vm.removeFromLoading(loadingName);
        });
        var authorsloadingname = $scope.vm.guid();
        $scope.vm.addToLoading(authorsloadingname)
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
            $scope.removeFromLoading(authorsloadingname);
		}, function(response) {
        	$mdToast.showSimple("Failed to export authors");
        	$vm.removeFromLoading(loadingName);
        });
		$mdDialog.cancel()
	};
});