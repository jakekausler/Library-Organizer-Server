angular.module('libraryOrganizer')
.controller('shelvesController', function($scope, $http) {
	$scope.editingShelves = false;
	$scope.shelfSearchString = "";
	$scope.bookcases = [];
	$scope.libraryid = '1';
	$scope.updateCases = function() {
		$http.get('/cases', {
			params: {
				libraryid: $scope.libraryid
			}
		}).then(function(response){
			$scope.bookcases = response.data;
		});
	}
	$scope.updateCases();
	//todo
	$scope.addShelf = function() {

	}
	//todo
	$scope.toggleEditShelves = function() {
		$scope.editingShelves = !$scope.editingShelves
	}
	//todo
	$scope.findBook = function() {
		
	}
});