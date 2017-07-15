angular.module('libraryOrganizer')
.controller('shelvesController', function($scope, $http) {
	$scope.editingShelves = false;
	$scope.shelfSearchString = "";
	$scope.bookcases = [];
	$scope.libraries = [];
	$scope.currentLibrary = 0;
	$scope.updateCases = function() {
		$http.get('/cases', {
			params: {
				libraryid: $scope.libraries[$scope.currentLibrary].id
			}
		}).then(function(response){
			$scope.bookcases = response.data;
		});
	}
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
    $scope.updateLibraries = function() {
        $http.get('/ownedlibraries', {}).then(function(response) {
            $scope.libraries = response.data;
			$scope.updateCases();
        });
    };
    $scope.updateLibraries();
});