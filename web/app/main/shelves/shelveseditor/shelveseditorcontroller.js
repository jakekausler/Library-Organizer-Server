angular.module('libraryOrganizer')
.controller('shelveseditorController', function ($scope, $mdDialog, $mdToast, $http, shelves, vm, libraryid) {
	$scope.vm = vm;
	$scope.libraryid = libraryid;
	$scope.shelves = angular.copy(shelves);
	$scope.toRemoveIds = [];
	for (i in $scope.shelves) {
		$scope.shelves[i].numberofshelves = $scope.shelves[i].shelves.length;
		$scope.shelves[i].shelfheight = $scope.shelves[i].shelves[0].height;
	}
	$scope.numberOfCases = $scope.shelves.length;
	$scope.defaultWidth = 0;
	$scope.defaultNumberOfShelves = 0;
	$scope.defaultPaddingLeft = 0;
	$scope.defaultPaddingRight = 0;
	$scope.defaultShelfHeight = 0;
	$scope.defaultSpacerSize = 0;
	$scope.vm.getSettingByName('Case Width', function(value) {
		$scope.defaultWidth = value;
	});
	$scope.vm.getSettingByName('Number of Shelves', function(value) {
		$scope.defaultNumberOfShelves = value;
	});
	$scope.vm.getSettingByName('Padding Left', function(value) {
		$scope.defaultPaddingLeft = value;
	});
	$scope.vm.getSettingByName('Padding Right', function(value) {
		$scope.defaultPaddingRight = value;
	});
	$scope.vm.getSettingByName('Shelf Height', function(value) {
		$scope.defaultShelfHeight = value;
	});
	$scope.vm.getSettingByName('Spacer Size', function(value) {
		$scope.defaultSpacerSize = value;
	});
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
	$scope.save = function() {
		var cases = [];
		for (i in $scope.shelves) {
			$scope.shelves[i].numberofshelves = parseInt($scope.shelves[i].numberofshelves);
			$scope.shelves[i].shelfheight = parseInt($scope.shelves[i].shelfheight);
			$scope.shelves[i].width = parseInt($scope.shelves[i].width);
			$scope.shelves[i].paddingleft = parseInt($scope.shelves[i].paddingleft);
			$scope.shelves[i].paddingright = parseInt($scope.shelves[i].paddingright);
			$scope.shelves[i].spacerheight = parseInt($scope.shelves[i].spacerheight);
			cases.push({
				id: $scope.shelves[i].id,	
				casenumber: $scope.shelves[i].casenumber,
				numberofshelves: $scope.shelves[i].numberofshelves,
				shelfheight: $scope.shelves[i].shelfheight,
				width: $scope.shelves[i].width,
				paddingleft: $scope.shelves[i].paddingleft,
				paddingright: $scope.shelves[i].paddingright,
				spacerheight: $scope.shelves[i].spacerheight
			})
		}
		$http({
			url: '/libraries/'+$scope.libraryid+'/cases',
			method: 'PUT',
			data: JSON.stringify({editedcases: cases, toremoveids: $scope.toRemoveIds})
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved shelves")
		});
		$http({
			url: '/libraries/'+$scope.libraryid+'/series',
			method: 'PUT',
			data: JSON.stringify($scope.authorBasedSeries)
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved misc")
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
		$scope.shelves.push({
			casenumber: $scope.shelves[$scope.shelves.length-1].casenumber+1,
			numberofshelves: $scope.defaultNumberOfShelves,
			width: $scope.defaultWidth,
			shelfheight: $scope.defaultShelfHeight,
			paddingleft: $scope.defaultPaddingLeft,
			paddingright: $scope.defaultPaddingRight,
			spacerheight: $scope.defaultSpacerSize,
			libraryid: $scope.libraryid
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
	$scope.series = [];
	$scope.authorBasedSeries = [];
	$scope.toAddSeries = "";
	$scope.updateSeries = function() {
		$http({
			url: '/information/series',
			method: 'GET'
		}).then(function(response){
			$scope.series = response.data;
		});
	}
    $scope.updateSeries();
	$scope.updateAuthorBasedSeries = function() {
		$http({
			url: '/libraries/'+$scope.libraryid+'/series',
			method: 'GET'
		}).then(function(response){
			if (response.data == null) {
				response.data = [];
			}
			$scope.authorBasedSeries = response.data;
		});
	}
    $scope.updateAuthorBasedSeries();
    $scope.addToSeries = function() {
    	if ($scope.toAddSeries != "" && !$scope.authorBasedSeries.includes($scope.toAddSeries)) {
    		$scope.authorBasedSeries.push($scope.toAddSeries)
    	}
    	$scope.toAddSeries = "";
    }
    $scope.removeFromSeries = function(series) {
    	if ($scope.authorBasedSeries.includes(series)) {
    		$scope.authorBasedSeries.splice($scope.authorBasedSeries.indexOf(series), 1);
    	}
    }
})
