angular.module('libraryOrganizer')
.controller('settingsController', function($scope, $mdDialog, $http, vm) {
    $scope.settings = [];
	$scope.pushSettings = function () {
		var settings = [];
		for (group in $scope.settings) {
			for (s in $scope.settings[group].settings) {
				setting = {
					group: $scope.settings[group].group,
					name: $scope.settings[group].settings[s].setting,
					value: $scope.settings[group].settings[s].value,
					valuetype: $scope.settings[group].settings[s].valuetype,
					possiblevalues: $scope.settings[group].settings[s].possiblevalues
				}
				settings.push(setting)
			}
		}
		$http({
			url: '/settings',
			method: 'PUT',
			data: JSON.stringify(settings)
		});
	}
    $scope.updateSettings = function() {
        $scope.settings = [];
        $http({
            url: '/settings',
            method: 'GET'
        }).then(function(response) {
            var settinggroups = {}
            response.data.forEach(function(v, i) {
                if (!settinggroups[v.group]) {
                    settinggroups[v.group] = [];
                }
                settinggroups[v.group].push({
                    setting: v.name,
                    value: v.value,
                    valuetype: v.valuetype,
                    possiblevalues: v.possiblevalues?v.possiblevalues:[]
                })
            })
            for (group in settinggroups) {
                var settings = [];
                for (setting in settinggroups[group]) {
                    settings.push({
                        setting: settinggroups[group][setting].setting,
                        value: settinggroups[group][setting].value,
                        valuetype: settinggroups[group][setting].valuetype,
                        possiblevalues: settinggroups[group][setting].possiblevalues
                    })
                }
                $scope.settings.push({
                    group: group,
                    settings: settings
                })
            }
        })
    }
    $scope.updateSettings();
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
    $scope.ownedLibraries = [];
    $scope.selectedLibrary = 0;
    $scope.updateOwnedLibraries = function() {
        $http({
            url: '/libraries/owned',
            method: 'GET'
        }).then(function(response) {
            $scope.ownedLibraries = response.data;
            for (i in $scope.ownedLibraries) {
                $scope.ownedLibraries[i].editusers = [];
                $scope.ownedLibraries[i].checkoutusers = [];
                $scope.ownedLibraries[i].viewusers = [];
                for (u in $scope.ownedLibraries[i].user) {
                    var pem = $scope.ownedLibraries[i].user[u].permission;
                    if ((pem&4)==4) {
                        $scope.ownedLibraries[i].editusers.push($scope.ownedLibraries[i].user[u])
                    }
                    if ((pem&2)==2) {
                        $scope.ownedLibraries[i].checkoutusers.push($scope.ownedLibraries[i].user[u])
                    }
                    if ((pem&1)==1) {
                        $scope.ownedLibraries[i].viewusers.push($scope.ownedLibraries[i].user[u])
                    }
                }
            }
        });
    };
    $scope.updateOwnedLibraries();
    $scope.queryUsers = function(users, criteria) {
        console.log(users)
        return criteria ? users.filter(function(s) {
            return (angular.lowercase(s.username).indexOf(angular.lowercase(criteria)) !== -1);
        }) : [];
    }
    $scope.users = [];
    $scope.updateUsers = function() {
        $http({
            url: '/getusers',
            method: 'GET'
        }).then(function(response) {
            $scope.users = response.data;
        })
    }
    $scope.updateUsers();
    $scope.removeLibrary = function(libraryid) {
        for (i in $scope.ownedLibraries) {
            if ($scope.ownedLibraries[i].id == libraryid) {
                $scope.ownedLibraries.splice(i, 1);
                return;
            }
        }
    }
    $scope.addToEdit = function(library) {
        if (!$scope.containsUser(library.editusers, library.toAddEdit)) {
            library.editusers.push(library.toAddEdit);
            library.toAddView = library.toAddEdit;
            $scope.addToView(library);
        }
        library.toAddEdit = null;
    }
    $scope.removeUserFromEdit = function(userid, library) {
        for (i in library.editusers) {
            if (library.editusers[i].id == userid) {
                library.editusers.splice(i, 1);
                return;
            }
        }
    }
    $scope.addToCheckout = function(library) {
        if (!$scope.containsUser(library.checkoutusers, library.toAddCheckout)) {
            library.checkoutusers.push(library.toAddCheckout);
            library.toAddView = library.toAddCheckout;
            $scope.addToView(library);
        }
        library.toAddCheckout = null;
    }
    $scope.removeUserFromCheckout = function(userid, library) {
        for (i in library.checkoutusers) {
            if (library.checkoutusers[i].id == userid) {
                library.checkoutusers.splice(i, 1);
                return;
            }
        }
    }
    $scope.addToView = function(library) {
        if (!$scope.containsUser(library.viewusers, library.toAddView)) {
            library.viewusers.push(library.toAddView);
        }
        library.toAddView = null;
    }
    $scope.removeUserFromView = function(userid, library) {
        for (i in library.viewusers) {
            if (library.viewusers[i].id == userid) {
                library.viewusers.splice(i, 1);
                $scope.removeUserFromEdit(userid, library);
                $scope.removeUserFromCheckout(userid, library);
                return;
            }
        }
    }
    $scope.addLibrary = function() {
        $scope.ownedLibraries.push({
            editusers: [],
            checkoutusers: [],
            viewusers: [],
            id: -1,
            name: 'New Library'
        });
        $scope.selectedLibrary = $scope.ownedLibraries.length-1;
    }
    $scope.saveLibraries = function() {
        for (i in $scope.ownedLibraries) {
            $scope.ownedLibraries[i].user = [];
            for (j in $scope.ownedLibraries[i].editusers) {
                $scope.ownedLibraries[i].user.push({
                    id: $scope.ownedLibraries[i].editusers[j].id,
                    username: $scope.ownedLibraries[i].editusers[j].username,
                    permission: 4
                })
            }
            for (j in $scope.ownedLibraries[i].checkoutusers) {
                added = false;
                for (k in $scope.ownedLibraries[i].user) {
                    if ($scope.ownedLibraries[i].user[k].id==$scope.ownedLibraries[i].checkoutusers[j].id) {
                        $scope.ownedLibraries[i].user[k].permission += 2;
                        added = true;
                    }
                }
                if (!added) {
                    $scope.ownedLibraries[i].user.push({
                        id: $scope.ownedLibraries[i].checkoutusers[j].id,
                        username: $scope.ownedLibraries[i].checkoutusers[j].username,
                        permission: 2
                    })
                }
            }
            for (j in $scope.ownedLibraries[i].viewusers) {
                added = false;
                for (k in $scope.ownedLibraries[i].user) {
                    if ($scope.ownedLibraries[i].user[k].id==$scope.ownedLibraries[i].viewusers[j].id) {
                        $scope.ownedLibraries[i].user[k].permission += 1;
                        added = true;
                    }
                }
                if (!added) {
                    $scope.ownedLibraries[i].user.push({
                        id: $scope.ownedLibraries[i].viewusers[j].id,
                        username: $scope.ownedLibraries[i].viewusers[j].username,
                        permission: 1
                    })
                }
            }
        }
        $http({
            url: '/libraries/owned',
            method: 'PUT',
            data: $scope.ownedLibraries
        }).then(function(response){});
    }
    $scope.containsUser = function(arr, user) {
        for (i in arr) {
            if (arr[i].id==user.id) {
                return true;
            }
        }
        return false;
    }
});