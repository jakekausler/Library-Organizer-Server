angular.module('libraryOrganizer')
    .controller('statisticsController', function($scope, $http) {
        $scope.statView = 'general';
        $scope.statSubView = 'bycounts';
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
            $http({
                url: '/dimensions',
                method: 'GET',
                params: {
                    libraryids: $scope.stringLibraryIds()
                }
            }).then(function(response) {
                $scope.dimensions = response.data;
            });
        }
        $scope.setStatView = function(view) {
            $scope.statView = view;
            switch (view) {
            case 'general':
                $scope.statSubView = 'bycounts';
                break;
            case 'series':
                $scope.statSubView = 'byseries';
                break;
            case 'publishers':
                $scope.statSubView = 'bybooksperparent';
                break;
            case 'languages':
                $scope.statSubView = 'byprimary';
                break;
            case 'deweys':
                $scope.statSubView = 'bydeweys';
                break;
            case 'formats':
                $scope.statSubView = 'byformats';
                break;
            case 'contributors':
                $scope.statSubView = 'bycontributorstop';
                break;
            case 'dimensions':
                $scope.statSubView = 'byvolumes';
                break;
            case 'dates':
                $scope.statSubView = 'bydatesoriginal';
                break;
            }
        }
        $scope.setStatSubView = function(view) {
            $scope.statSubView = view;
        }
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
                        children: [],
                        selected: false
                    });
                }
                for (k in libStructure) {
                    if (!$scope.currentLibraryId && k == $scope.username) {
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
