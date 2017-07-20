angular.module('libraryOrganizer')
.controller('viewController', function($scope, $mdDialog, book, $vm, viewType, username) {
	$scope.book = book;
	$scope.vm = $vm;
	$scope.viewType = viewType;
	$scope.username = username;
	$scope.canEdit = (book.library.permissions&4)==4;
	$scope.canCheckout = (book.library.permissions&2)==2;
	$scope.checkout = function(ev) {
	    var d = $mdDialog.confirm()
	    	.title("Are you sure you would like to checkout this book?")
	    	.textContent("If you checkout this book, it will become unavailable for other users to checkout and will show as checked out to you.")
	    	.ariaLabel("Checkout")
	    	.targetEvent(ev)
	    	.ok("Yes")
	    	.cancel("Cancel");
	    $mdDialog.show(d).then(function() {
	    	
	    }, function() {
	    	$scope.cancel()
	    });
	}
	$scope.edit = function(ev) {
		$scope.vm.showEditDialog(ev, $scope.book, $scope.vm, $scope.viewType);
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