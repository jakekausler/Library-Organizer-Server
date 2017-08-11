angular.module('libraryOrganizer')
.controller('viewController', function($scope, $mdDialog, $mdToast, $http, book, $vm, viewType, username) {
	$scope.book = book;
	$scope.vm = $vm;
	$scope.viewType = viewType;
	$scope.username = username;
	$scope.updateBook = function() {
		if ($scope.book.bookid) {
            var loadingName = $scope.vm.guid();
            $scope.vm.addToLoading(loadingName)
			$http({
				url: 'books/'+$scope.book.bookid,
				method: 'GET'
			}).then(function(response) {
				$scope.book = response.data;
				$scope.canEdit = (book.library.permissions&4)==4;
				$scope.canCheckout = (book.library.permissions&2)==2 && book.loanee.id == -1 && !book.isreading && book.isowned && !book.isshipping;
				$scope.canCheckin = book.loanee.id != -1 && ((book.library.permissions&4)==4 || book.loanee.username == $scope.username);
				$scope.vm.removeFromLoading(loadingName);
			}, function(response) {
	        	$mdToast.showSimple("Failed to retrieve book information");
	        	$vm.removeFromLoading(loadingName);
	        	$scope.cancel()
			})
		}
	}
	$scope.updateBook()
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
	    	}, function(response) {
	        	$mdToast.showSimple("Failed to check out book");
	        	$vm.removeFromLoading(loadingName);
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
	    	}, function(response) {
	        	$mdToast.showSimple("Failed to return book");
	        	$vm.removeFromLoading(loadingName);
	        	$scope.cancel()
	        })
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
	$scope.copy = function(ev) {
		var book = angular.copy($scope.book);
		book.loanee.id = -1;
		book.library.id = '';
		book.bookid = '';
		$scope.vm.showEditDialog(ev, book, $scope.vm, $scope.viewType);
	}
	$scope.averageRating = -1;
	$scope.userRating = 0;
	$scope.numRatings = 0;
	$scope.updateRating = function() {
		var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
    	$http({
    		url: 'books/'+$scope.book.bookid+"/ratings",
    		method: 'GET'
    	}).then(function(response) {
    		if (response.data) {
				$scope.numRatings = response.data.length;
    			$scope.averageRating = 0.0;
    			for (i in response.data) {
    				if ($scope.username==response.data[i].username) {
    					$scope.userRating = response.data[i].rating;
    				}
    				$scope.averageRating += response.data[i].rating
    			}
    			$scope.averageRating /= response.data.length;
    		} else {
    			$scope.averageRating = -1;
    		}
        	$scope.vm.removeFromLoading(loadingName);
    	}, function(response) {
        	$mdToast.showSimple("Failed to get ratings");
        	$vm.removeFromLoading(loadingName);
        })
	}
	$scope.updateRating()
	$scope.rate = function() {
		var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
    	$http({
    		url: 'books/'+$scope.book.bookid+"/ratings",
    		method: 'PUT',
    		data: JSON.stringify($scope.userRating)
    	}).then(function(response) {
			$mdToast.showSimple("Successfully rated book")
			$scope.updateRating();
        	$scope.vm.removeFromLoading(loadingName);
    	}, function(response) {
        	$mdToast.showSimple("Failed to rate book");
        	$vm.removeFromLoading(loadingName);
        })
	}
	$scope.reviews = [{
		username: $scope.username,
		review: '',
		bookid: $scope.book.bookid
	}];
	$scope.updateReviews = function() {
		var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
    	$http({
    		url: 'books/'+$scope.book.bookid+"/reviews",
    		method: 'GET'
    	}).then(function(response) {
    		if (response.data) {
				for (i in response.data) {
					if (response.data[i].username == $scope.username) {
						$scope.reviews[0] = response.data[i];
					} else {
						$scope.reviews.push(response.data[i])
					}
				}
    		} else {
    			$scope.reviews = [{
					username: $scope.username,
					review: '',
					bookid: $scope.book.bookid
				}];
    		}
        	$scope.vm.removeFromLoading(loadingName);
		}, function(response) {
        	$mdToast.showSimple("Failed to get reviews");
        	$vm.removeFromLoading(loadingName);
        })
    }
	$scope.saveReview = function() {
		var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
    	$http({
    		url: 'books/'+$scope.book.bookid+"/reviews",
    		method: 'PUT',
    		data: JSON.stringify($scope.reviews[0].review)
    	}).then(function(response) {
			$mdToast.showSimple("Successfully reviewed book")
			$scope.updateRating();
        	$scope.vm.removeFromLoading(loadingName);
    	}, function(response) {
        	$mdToast.showSimple("Failed to review book");
        	$vm.removeFromLoading(loadingName);
        })
	}
})