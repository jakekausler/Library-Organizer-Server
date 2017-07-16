angular.module('libraryOrganizer')
.controller('shelvesController', function($scope, $http) {
	$scope.editingShelves = false;
	$scope.shelfSearchString = "";
	$scope.bookcases = [];
	$scope.libraries = [];
	$scope.output = [];
	$scope.currentLibraryId = $scope.getParameterByName("shelfselectedlibrary", "");
	$scope.canEditCurrentShelf = false;
	$scope.updateCases = function() {
		for (o in $scope.libraries) {
			for (l in $scope.libraries[o].children) {
				if ($scope.libraries[o].children[l].selected) {
					$scope.currentLibraryId = $scope.libraries[o].children[l].id;
            		$scope.canEditCurrentShelf = ($scope.libraries[o].children[l].permissions&4)==4;
				}
			}
		}
        $scope.setParameters({shelfselectedlibrary: $scope.currentLibraryId})
		$http.get('/cases', {
			params: {
				libraryid: $scope.currentLibraryId,
				sortmethod: 'DEWEY'
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
        $http.get('/libraries', {}).then(function(response) {
            $scope.libraries = response.data;
            var data = [];
            var libStructure = {}
            for (l in $scope.libraries) {
                if (!libStructure[$scope.libraries[l].owner]) {
                    libStructure[$scope.libraries[l].owner] = [];
                }
                libStructure[$scope.libraries[l].owner].push({
                    id: $scope.libraries[l].id,
                    name: $scope.libraries[l].name,
                    permissions: $scope.libraries[l].permissions,
                    children: [],
                    selected: $scope.libraries[l].id == $scope.currentLibraryId
                });
            }
            for (k in libStructure) {
            	if (!$scope.currentLibraryId && k == $scope.username) {
            		$scope.currentLibraryId = libStructure[k][0].id;
            		$scope.canEditCurrentShelf = (libStructure[k][0].permissions&4)==4;
            		libStructure[k][0].selected = true;
            	}
                data.push({
                    id: "owner/"+k,
                    name: k,
                    children: libStructure[k],
                    selected: false
                })
            }
            $scope.libraries = angular.copy(data);
            $scope.updateCases();
        });
    };
    $scope.updateLibraries();
    $scope.chooseLibrary = function($ev) {
    	$scope.showLibraryChooserDialog($ev, $scope, false)
    }
    $scope.$watch('output', function() {
        $scope.updateCases();
    })
});