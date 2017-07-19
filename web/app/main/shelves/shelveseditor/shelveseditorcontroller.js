angular.module('libraryOrganizer')
.controller('shelveseditorController', function ($scope, $mdDialog, $http, shelves, vm, libraryid) {
	$scope.vm = vm;
	$scope.libraryid = libraryid;
	$scope.shelves = angular.copy(shelves);
	$scope.toRemoveIds = [];
	for (i in $scope.shelves) {
		$scope.shelves[i].numberofshelves = $scope.shelves[i].shelves.length;
		$scope.shelves[i].shelfheight = $scope.shelves[i].shelves[0].height;
	}
	$scope.numberOfCases = $scope.shelves.length;
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
	$scope.save = function() {
		for (i in $scope.shelves) {
			$scope.shelves[i].numberofshelves = parseInt($scope.shelves[i].numberofshelves);
			$scope.shelves[i].shelfheight = parseInt($scope.shelves[i].shelfheight);
			$scope.shelves[i].width = parseInt($scope.shelves[i].width);
			$scope.shelves[i].paddingleft = parseInt($scope.shelves[i].paddingleft);
			$scope.shelves[i].paddingright = parseInt($scope.shelves[i].paddingright);
			$scope.shelves[i].spacerheight = parseInt($scope.shelves[i].spacerheight);
		}
		$http({
			url: '/savecases',
			method: 'POST',
			data: JSON.stringify({editedcases: $scope.shelves, toremoveids: $scope.toRemoveIds, libraryid: $scope.libraryid})
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
		});
	};
	$scope.moveShelfUp = function(cas) {
		var c = cas.casenumber;
		$scope.shelves[c-1].casenumber = c-1;
		$scope.shelves[c-2].casenumber = c;
		$scope.shelves.sort(function(a, b) {
			return (a.casenumber > b.casenumber) ? 1 : ((a.casenumber < b.casenumber) ? -1 : 0);
		});
	};
	$scope.moveShelfDown = function(cas) {
		var c = cas.casenumber;
		$scope.shelves[c-1].casenumber = c+1;
		$scope.shelves[c].casenumber = c;
		$scope.shelves.sort(function(a, b) {
			return (a.casenumber > b.casenumber) ? 1 : ((a.casenumber < b.casenumber) ? -1 : 0);
		});
	};
	$scope.addShelf = function(cas) {
		var c = cas.casenumber;
		for (i=c; i<$scope.shelves.length; i++) {
			$scope.shelves[i].casenumber = i+2;
		}
		$scope.shelves.push({
			casenumber: c+1,
			numberofshelves: '',
			width: '',
			shelfheight: '',
			paddingleft: '',
			paddingright: '',
			spacerheight: '',
			libraryid: $scope.libraryid
		});
		$scope.shelves.sort(function(a, b) {
			return (a.casenumber > b.casenumber) ? 1 : ((a.casenumber < b.casenumber) ? -1 : 0);
		});
		$scope.numberOfCases++;
	};
	$scope.removeShelf = function(cas) {
		var c = cas.casenumber;
		for (i=c; i<$scope.shelves.length; i++) {
			$scope.shelves[i].casenumber = i;
		}
		if ($scope.shelves[c-1].id) {
			$scope.toRemoveIds.push($scope.shelves[c-1].id);
		}
		$scope.shelves.splice(c-1,1)
		$scope.shelves.sort(function(a, b) {
			return (a.casenumber > b.casenumber) ? 1 : ((a.casenumber < b.casenumber) ? -1 : 0);
		});
		$scope.numberOfCases--;
	};
})