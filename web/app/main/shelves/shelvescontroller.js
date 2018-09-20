angular.module('libraryOrganizer')
    .controller('shelvesController', function($scope, $mdToast, $http, $mdDialog) {
        $scope.editingShelves = false;
        $scope.shelfSearchString = "";
        $scope.bookcases = [];
        $scope.libraries = [];
        $scope.output = [];
        $scope.currentLibraryId = $scope.getParameterByName("shelfselectedlibrary", "");
        $scope.canEditCurrentShelf = false;
        $scope.updateCases = function() {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName);
            for (var o in $scope.libraries) {
                for (var l in $scope.libraries[o].children) {
                    if ($scope.libraries[o].children[l].selected) {
                        $scope.currentLibraryId = $scope.libraries[o].children[l].id;
                        $scope.canEditCurrentShelf = ($scope.libraries[o].children[l].permissions & 4) == 4;
                    }
                }
            }
            $scope.setParameters({
                shelfselectedlibrary: $scope.currentLibraryId
            });
            $http({
                $mdToast,
                url: '/libraries/' + $scope.currentLibraryId + '/cases',
                method: 'GET'
            }).then(function(response) {
                $scope.bookcases = response.data;
                if (!$scope.bookcases) {
                    $scope.bookcases = [];
                }
                $scope.removeFromLoading(loadingName);
            }, function(response) {
                $mdToast.showSimple("Failed to get library cases");
                $scope.removeFromLoading(loadingName);
            });
        };
        $scope.editShelves = function(ev) {
            $mdDialog.show({
                controller: 'shelveseditorController',
                templateUrl: 'web/app/main/shelves/shelveseditor/shelveseditordialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false,
                multiple: true,
                locals: {
                    vm: $scope,
                    shelves: $scope.bookcases,
                    libraryid: $scope.currentLibraryId
                }
            });
        };
        $scope.findBook = function(ev) {
            if (ev.key !== "Enter") {
                return;
            }
            if (ev.target.value == "" || !ev.target.value) {
                UnhighlightMatches();
                return;
            }
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName);
            $http({
                url: '/libraries/' + $scope.currentLibraryId + '/search',
                method: 'GET',
                params: {
                    text: ev.target.value
                }
            }).then(function(response) {
                console.log(response);
                HighlightMatches(response.data);
                $scope.removeFromLoading(loadingName);
            }, function(response) {
                $mdToast.showSimple("Failed to search");
                $scope.removeFromLoading(loadingName);
            });
        };
        HighlightMatches = function(matches) {
            var currentMatch = 0;
            console.log($scope.bookcases);
            for (var i = 0; i < $scope.bookcases.length; i++) {
                for (var j = 0; j < $scope.bookcases[i].shelves.length; j++) {
                    if ($scope.bookcases[i].shelves[j].books) {
                    for (var k = 0; k < $scope.bookcases[i].shelves[j].books.length; k++) {
                        if (currentMatch < matches.length && matches[currentMatch].case == i && matches[currentMatch].shelf == j && matches[currentMatch].book == k) {
                            $scope.bookcases[i].shelves[j].books[k].highlight = true;
                            currentMatch++;
                        } else {
                            $scope.bookcases[i].shelves[j].books[k].highlight = false;
                        }
                    }
                    }
                }
            }
        };
        UnHighlightMatches = function() {
            for (var i = 0; i < $scope.bookcases.length; i++) {
                for (var j = 0; j < $scope.bookcases[i].shelves.length; j++) {
                    for (var k = 0; k < $scope.bookcases[i].shelves[j].books.length; k++) {
                        $scope.bookcases[i].shelves[j].books[k].highlight = undefined;
                    }
                }
            }
        };
        $scope.updateLibraries = function() {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName);
            $http({
                url: '/libraries',
                method: 'GET'
            }).then(function(response) {
                $scope.libraries = response.data;
                var data = [];
                var libStructure = {};
                for (var l in $scope.libraries) {
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
                for (var k in libStructure) {
                    if (!$scope.currentLibraryId && k == $scope.username) {
                        $scope.currentLibraryId = libStructure[k][0].id;
                        $scope.canEditCurrentShelf = (libStructure[k][0].permissions & 4) == 4;
                        libStructure[k][0].selected = true;
                    }
                    data.push({
                        id: "owner/" + k,
                        name: k,
                        children: libStructure[k],
                        selected: false
                    });
                }
                $scope.libraries = angular.copy(data);
                $scope.updateCases();
                $scope.removeFromLoading(loadingName);
            }, function(response) {
                $mdToast.showSimple("Failed to get list of libraries");
                $scope.removeFromLoading(loadingName);
            });
        };
        $scope.updateLibraries();
        $scope.chooseLibrary = function($ev) {
            $scope.showLibraryChooserDialog($ev, $scope, false);
        };
        $scope.$watch('output', function() {
            if ($scope.libraries.length > 0) {
                $scope.updateCases();
            }
        });
    });
