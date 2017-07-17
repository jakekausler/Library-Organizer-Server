angular.module('libraryOrganizer')
.controller('viewController', function($scope, $http, $mdDialog, book, $vm, viewType, username) {
	$scope.book = book;
	$scope.vm = $vm;
	$scope.viewType = viewType;
	$scope.username = username;
	$scope.canEdit = (book.library.permissions&4)==4;
	$scope.canCheckout = (book.library.permissions&2)==2;
	$scope.edit = function(ev) {
	    $mdDialog.show({
	        controller: 'editorController',
	        templateUrl: 'web/app/main/bookviews/editor/editordialog.html',
	        parent: angular.element(document.body),
	        targetEvt: ev,
	        clickOutsideToClose: true,
	        fullscreen: false,
	        locals: {
	            book: $scope.book,
	            $vm: $scope.vm,
	            viewType: $scope.viewType,
	            username: $scope.username
	        }
	    })
	}
	$scope.checkout = function(ev) {
	    $mdDialog.show({
	        controller: 'checkoutController',
	        templateUrl: 'web/app/main/bookviews/checkout/checkoutdialog.html',
	        parent: angular.element(document.body),
	        targetEvt: ev,
	        clickOutsideToClose: true,
	        fullscreen: false,
	        locals: {
	            book: $scope.book,
	            $vm: $scope.vm,
	            viewType: $scope.viewType
	        }
	    })
	}
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
	$scope.contributors = '';
	$scope.getContributors = function() {
		var contrib = []
		for (c in $scope.book.contributors) {
			c = $scope.book.contributors[c];
			contrib.push(c.name.first + " " + c.name.middles.replace(";"," ") + " " + c.name.last + " (" + c.role + ")")
		}
		$scope.contributors = contrib.join(', ');
	}
	$scope.originallypublished = $scope.book.originallypublished;
	if ($scope.book.editionpublished==$scope.originallypublished || !$scope.book.editionpublished) {
		$scope.editionpublished = $scope.originallypublished;
		$scope.originallypublished = '';
	}
	if ($scope.originallypublished == 0) {
		$scope.originallypublished = '';
	}
	if ($scope.editionpublished == 0) {
		$scope.editionpublished = '';
	}
	$scope.getContributors();
})