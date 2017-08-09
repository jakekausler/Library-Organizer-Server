angular.module('libraryOrganizer')
    .controller('statisticsController', function($scope, $http) {
        $scope.statSelectedLibraries = $scope.getParameterByName("statsselectedlibraries", "").split(',')
        $scope.libraries = [];
        $scope.output = [];
        $scope.stringLibraryIds = function() {
            var retval = "";
            for (o in $scope.output) {
                if ($scope.output[o].children.length == 0 && $scope.output[o].selected) {
                    retval += $scope.output[o].id + ",";
                } else {
                    for (l in $scope.output[o].children) {
                        if ($scope.output[o].children[l].selected) {
                            retval += $scope.output[o].children[l].id + ",";
                        }
                    }
                }
            }
            if (retval.endsWith(",")) {
                retval = retval.substring(0,retval.length-1)
            }
            return retval;
        }
        $scope.dimensions = {};
        $scope.updateDimensions = function() {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName)
            $scope.statSelectedLibraries = $scope.stringLibraryIds();
            $scope.setParameters({
                statsselectedlibraries: $scope.statSelectedLibraries
            })
            $http({
                url: '/information/dimensions',
                method: 'GET',
                params: {
                    libraryids: $scope.stringLibraryIds()
                }
            }).then(function(response) {
                $scope.dimensions = response.data;
                $scope.removeFromLoading(loadingName)
            }).then(function(response) {
                $mdToast.showSimple("Failed to get dimension data");
                $vm.removeFromLoading(loadingName);
            });
        }
        $scope.setStatView = function(view) {
            $scope.statView = view;
            $scope.setParameters({'statview': view})
            switch (view) {
            case 'general':
                $scope.setStatSubView('bycounts');
                break;
            case 'series':
                $scope.setStatSubView('byseries');
                break;
            case 'publishers':
                $scope.setStatSubView('bybooksperparent');
                break;
            case 'languages':
                $scope.setStatSubView('byprimary');
                break;
            case 'deweys':
                $scope.setStatSubView('bydeweys');
                break;
            case 'bindings':
                $scope.setStatSubView('bybindings');
                break;
            case 'contributors':
                $scope.setStatSubView('bycontributorstop');
                break;
            case 'dimensions':
                $scope.setStatSubView('byvolumes');
                break;
            case 'dates':
                $scope.setStatSubView('bydatesoriginal');
                break;
            case 'lexile':
                $scope.setStatSubView('bylexile');
                break;
            case 'tag':
                $scope.setStatSubView('bytag');
                break;
            }
        }
        $scope.setStatSubView = function(view) {
            if (view) {
                $scope.statSubView = view;
                $scope.setParameters({'statsubview': view})
            }
        }
        $scope.setStatView($scope.getParameterByName("statview", "general"));
        $scope.setStatSubView($scope.getParameterByName("statsubview", ""));
        $scope.setSelected = function(data) {
            for (d in data) {
                if (data[d].selected) {
                    $scope.setMatchedId(data[d].id, $scope.libraries)
                }
                $scope.setSelected(data[d].children)
            }
        };
        $scope.setMatchedId = function(id, data) {
            for (d in data) {
                if (data[d].id == id) {
                    data[d].selected = true;
                }
            }
        }
        $scope.updateLibraries = function() {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName)
            $http({
                url: '/libraries',
                method: 'GET'
            }).then(function(response) {
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
                        children: [],
                        selected: $.inArray($scope.libraries[l].id+"", $scope.statSelectedLibraries)!=-1
                    });
                }
                for (k in libStructure) {
                    if ((!$scope.statSelectedLibraries || !$scope.statSelectedLibraries[0]) && !$scope.currentLibraryId && k == $scope.username) {
                        $scope.currentLibraryId = libStructure[k][0].id
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
                $scope.output = angular.copy($scope.libraries);
                $scope.updateDimensions();
                $scope.removeFromLoading(loadingName)
            }).then(function(response) {
                $mdToast.showSimple("Failed to get list of libraries");
                $vm.removeFromLoading(loadingName);
            });
        };
        $scope.updateLibraries();
        $scope.chooseLibrary = function($ev) {
            $scope.showLibraryChooserDialog($ev, $scope, true)
        }
        $scope.$watch('output', function() {
            $scope.updateDimensions();
        })
    });
