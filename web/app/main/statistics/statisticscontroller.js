angular.module('libraryOrganizer')
.controller('statisticsController', function($scope, $http) {
	$scope.statView = 'general';
	$scope.statSubView = 'bycounts';
	$scope.dimensions = {};$scope.updateDimensions = function() {
		$http({
			url: '/dimensions',
			method: 'POST',
		}).then(function(response){
			console.log(response.data);
			$scope.dimensions = response.data;
		});
	}
	$scope.updateDimensions();
});