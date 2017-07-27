angular.module('libraryOrganizer')
.controller('viewController', function($scope, $mdDialog, $mdToast, $http, book, $vm, viewType, username) {
	$scope.book = book;
	$scope.vm = $vm;
	$scope.viewType = viewType;
	$scope.username = username;
	$scope.canEdit = (book.library.permissions&4)==4;
	$scope.canCheckout = (book.library.permissions&2)==2 && book.loanee.id == -1;
	$scope.canCheckin = book.loanee.id != -1 && ((book.library.permissions&4)==4 || book.loanee.username == $scope.username);
	$scope.checkout = function(ev) {
	    var d = $mdDialog.confirm()
	    	.title("Are you sure you would like to checkout this book?")
	    	.textContent("If you checkout this book, it will become unavailable for other users to checkout and will show as checked out to you.")
	    	.ariaLabel("Checkout")
	    	.targetEvent(ev)
	    	.ok("Yes")
	    	.cancel("Cancel");
	    $mdDialog.show(d).then(function() {
            var loadingName = $scope.vm.guid();
            $scope.vm.addToLoading(loadingName)
	    	$http({
	    		url: 'books/checkout',
	    		method: 'PUT',
	    		data: $scope.book.bookid
	    	}).then(function(response) {
				if ($scope.viewType=='gridadd') {
					$scope.vm.updateRecieved();
				} else if ($scope.viewType=='shelves') {
					$scope.vm.updateCases();
				} else if ($scope.viewType=='grid') {
					$scope.vm.$parent.$parent.updateRecieved();
				} else if ($scope.viewType=='scanadd') {
					$scope.vm.$parent.$parent.updateRecieved();
				}
				$mdToast.showSimple("Successfully checked out book")
            	$scope.vm.removeFromLoading(loadingName);
	    		$scope.cancel()
	    	})
	    }, function() {
	    	$scope.cancel()
	    });
	}
	$scope.checkin = function(ev) {
	    var d = $mdDialog.confirm()
	    	.title("Are you sure you would like to return this book?")
	    	.textContent("If you return this book, it will become available for other users to checkout and no longer show as checked out to you.")
	    	.ariaLabel("Checkin")
	    	.targetEvent(ev)
	    	.ok("Yes")
	    	.cancel("Cancel");
	    $mdDialog.show(d).then(function() {
            var loadingName = $scope.vm.guid();
            $scope.vm.addToLoading(loadingName)
	    	$http({
	    		url: 'books/checkin',
	    		method: 'PUT',
	    		data: $scope.book.bookid
	    	}).then(function(response) {
				if ($scope.viewType=='gridadd') {
					$scope.vm.updateRecieved();
				} else if ($scope.viewType=='shelves') {
					$scope.vm.updateCases();
				} else if ($scope.viewType=='grid') {
					$scope.vm.$parent.$parent.updateRecieved();
				} else if ($scope.viewType=='scanadd') {
					$scope.vm.$parent.$parent.updateRecieved();
				}
				$mdToast.showSimple("Returned book")
            	$scope.vm.removeFromLoading(loadingName);
	    		$scope.cancel()
	    	})
	    }, function() {
			$mdToast.showSimple("Failed to return book")
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
	$scope.copy = function(ev) {
		var book = angular.copy($scope.book);
		book.loanee.id = -1;
		book.library.id = '';
		book.bookid = '';
		$scope.vm.showEditDialog(ev, book, $scope.vm, $scope.viewType);
	}
	$scope.averageRating = -1;
})