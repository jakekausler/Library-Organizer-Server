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
        var sortLoadingName = $scope.vm.guid();
        $scope.vm.addToLoading(sortLoadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/sort',
			method: 'PUT',
			data: JSON.stringify($scope.stringifySortOrders())
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved sort method")
        	$scope.vm.removeFromLoading(sortLoadingName)
		}, function(response) {
			$mdDialog.cancel();
            $mdToast.showSimple("Failed to save sort method");
            $vm.removeFromLoading(loadingName);
        });
        var casesLoadingName = $scope.vm.guid();
        $scope.vm.addToLoading(casesLoadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/cases',
			method: 'PUT',
			data: JSON.stringify({editedcases: cases, toremoveids: $scope.toRemoveIds})
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved shelves")
        	$scope.vm.removeFromLoading(casesLoadingName)
		}, function(response) {
			$mdDialog.cancel();
            $mdToast.showSimple("Failed to save shelves");
            $vm.removeFromLoading(loadingName);
        });
        var seriesLoadingName = $scope.vm.guid();
        $scope.vm.addToLoading(seriesLoadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/series',
			method: 'PUT',
			data: JSON.stringify($scope.authorBasedSeries)
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved series types")
        	$scope.vm.removeFromLoading(seriesLoadingName)
		}, function(response) {
			$mdDialog.cancel();
            $mdToast.showSimple("Failed to save series types");
            $vm.removeFromLoading(loadingName);
        });
        var breaksLoadingName = $scope.vm.guid();
        $scope.vm.addToLoading(breaksLoadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/breaks',
			method: 'PUT',
			data: JSON.stringify($scope.breaks)
		}).then(function(response) {
			$scope.vm.updateCases();
			$mdDialog.cancel();
            $mdToast.showSimple("Successfully saved breaks")
        	$scope.vm.removeFromLoading(breaksLoadingName)
		}, function(response) {
			$mdDialog.cancel();
            $mdToast.showSimple("Failed to save breaks");
            $vm.removeFromLoading(loadingName);
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
	$scope.addShelf = function() {
		var casenumber = $scope.shelves[$scope.shelves.length-1]?$scope.shelves[$scope.shelves.length-1].casenumber+1:1;
		$scope.shelves.push({
			casenumber: casenumber,
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
        var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
		$http({
			url: '/information/series',
			method: 'GET'
		}).then(function(response){
			$scope.series = response.data;
        	$scope.vm.removeFromLoading(loadingName)
		}, function(response) {
            $mdToast.showSimple("Failed to get series list");
            $vm.removeFromLoading(loadingName);
        });
	}
    $scope.updateSeries();
	$scope.updateAuthorBasedSeries = function() {
        var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/series',
			method: 'GET'
		}).then(function(response) {
            if (response.data == null) {
 				response.data = [];
 			}
 			$scope.authorBasedSeries = response.data;
         	$scope.vm.removeFromLoading(loadingName)
        }, function(response) {
            $mdToast.showSimple("Failed to get special series list");
            $vm.removeFromLoading(loadingName);
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
    $scope.add = function() {
    	if ($scope.selectedTab == 0) {
    		$scope.addShelf();
    	} else if ($scope.selectedTab == 1) {
    		$scope.addBreak();
    	}
    }
    $scope.breaks = [];
	$scope.updateBreaks = function() {
        var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/breaks',
			method: 'GET'
		}).then(function(response){
			if (response.data == null) {
				response.data = [];
			}
			$scope.breaks = response.data;
        	$scope.vm.removeFromLoading(loadingName)
		}, function(response) {
            $mdToast.showSimple("Failed to get existing breaks");
            $vm.removeFromLoading(loadingName);
        });
	}
    $scope.updateBreaks();
    $scope.breaktypes = ["SHELF", "CASE"];
    $scope.valuetypes = ["DEWEY", "SERIES"];
	$scope.updateValueType =function(b) {
		if (b.valuetype == "DEWEY") {
			b.possiblevalues = $scope.deweys;
		}
		if (b.valuetype == "DEWEY") {
			b.possiblevalues = $scope.series;
		}
	}
	$scope.addBreak = function() {
		$scope.breaks.push({
			valuetype: 'DEWEY',
			value: '',
			breaktype: 'SHELF',
			libraryid: $scope.libraryid
		});
		$scope.numberOfCases++;
	};
	$scope.removeBreak = function(i) {
		$scope.breaks.splice(i, 1)
	}
	$scope.sortOrder = [];
	$scope.specialSortOrder = [];
	$scope.updateSortOrders = function() {
        var loadingName = $scope.vm.guid();
        $scope.vm.addToLoading(loadingName)
		$http({
			url: '/libraries/'+$scope.libraryid+'/sort',
			method: 'GET'
		}).then(function(response){
			var orders = response.data.split("||");
			var normalOrders = orders[0].split("--");
			for (o in normalOrders) {
				var order = normalOrders[o].split(":")
				$scope.sortOrder.push({
					label: order[0],
					method: order[1]
				})
			}
			var specialOrders = orders[1].split("--");
			for (o in specialOrders) {
				var order = specialOrders[o].split(":")
				$scope.specialSortOrder.push({
					label: order[0],
					method: order[1]
				})
			}
        	$scope.vm.removeFromLoading(loadingName)
		}, function(response) {
            $mdToast.showSimple("Failed to get sort orders");
            $vm.removeFromLoading(loadingName);
        });
	}
	$scope.updateSortOrders();
	$scope.stringifySortOrders = function() {
		var normalOrders = [];
		var specialOrders = [];
		for (o in $scope.sortOrder) {
			normalOrders.push($scope.sortOrder[o].label+":"+$scope.sortOrder[o].method)
		}
		for (o in $scope.specialSortOrder) {
			specialOrders.push($scope.specialSortOrder[o].label+":"+$scope.specialSortOrder[o].method)
		}
		return normalOrders.join("--")+"||"+specialOrders.join("--")
	}
})
