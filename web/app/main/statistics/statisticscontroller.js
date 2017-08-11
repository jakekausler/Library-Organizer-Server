angular.module('libraryOrganizer')
    .controller('statisticsController', function($scope, $mdToast, $http) {
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
            }, function(response) {
                $mdToast.showSimple("Failed to get dimension data");
                $scope.removeFromLoading(loadingName);
            });
        }
        $scope.chartData = {};
        $scope.getChartData = function(chartType, caption, subcaption, labelDisplay, formatNumberScale, numberSuffix, decimals) {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName);
            $http({
                url: "information/statistics",
                method: 'GET',
                params: {
                    type: chartType,
                    libraryids: $scope.stringLibraryIds()
                }
            }).then(function(response) {
                $scope.chartData.chart = {
                    caption: caption?caption:undefined,
                    subcaption: subcaption?subcaption:undefined,
                    labelDisplay: labelDisplay?labelDisplay:undefined,
                    formatNumberScale: formatNumberScale?formatNumberScale:undefined,
                    numberSuffix: numberSuffix?numberSuffix:undefined,
                    decimals: decimals?decimals:undefined
                }
                if (response.data.total) {
                    $scope.chartData.chart.caption += " (Total: " + response.data.total + ")"
                }
                $scope.chartData.data = response.data.data;
                $scope.removeFromLoading(loadingName);
            }, function(response) {
                chartData = {}
                $mdToast.showSimple("Failed to update chart");
                $scope.removeFromLoading(loadingName);
            })
        }
        $scope.setStatView = function(view) {
            $scope.statView = view;
            $scope.setParameters({'statview': view})
            switch (view) {
            case 'general':
                $scope.setStatSubView('generalbycounts');
                break;
            case 'series':
                $scope.setStatSubView('series');
                break;
            case 'publishers':
                $scope.setStatSubView('publishersbooksperparent');
                break;
            case 'languages':
                $scope.setStatSubView('languagesprimary');
                break;
            case 'deweys':
                $scope.setStatSubView('deweys');
                break;
            case 'bindings':
                $scope.setStatSubView('formats');
                break;
            case 'contributors':
                $scope.setStatSubView('contributorstop');
                break;
            case 'dimensions':
                $scope.updateDimensions();
                $scope.setStatSubView('byvolumes');
                break;
            case 'dates':
                $scope.setStatSubView('datesoriginal');
                break;
            case 'lexile':
                $scope.setStatSubView('lexile');
                break;
            case 'tag':
                $scope.setStatSubView('tag');
                break;
            }
        }
        $scope.setStatSubView = function(view) {
            if (view) {
                $scope.statSubView = view;
                $scope.setParameters({'statsubview': view})
                var caption = "";
                var subcaption = "";
                var labelDisplay = "rotate";
                var formatNumberScale = "";
                var numberSuffix = "";
                var decimals = "";
                switch (view) {
                    case "generalbycounts":
                        formatNumberScale = "0";
                        caption = "Books By Count";
                        break;
                    case "generalbysize":
                        numberSuffix = " mm³";
                        decimals = "0";
                        caption = "Books by Size";
                        break;
                    case "generalbypages":
                        numberSuffix = " pages";
                        decimals = "0";
                        caption = "Books by Pages";
                        break;
                    case "publishersbooksperparent":
                        formatNumberScale = "0";
                        caption = "Books by Parent Company";
                        break;
                    case "publisherstopchildren":
                        formatNumberScale = "0";
                        caption = "Books by Top Publishers";
                        break;
                    case "publisherstoplocations":
                        formatNumberScale = "0";
                        caption = "Books by Top Locations";
                        break;
                    case "series":
                        formatNumberScale = "0";
                        caption = "Books by Series";
                        break;
                    case "languagesprimary":
                        formatNumberScale = "0";
                        caption = "Books by Primary Language";
                        break;
                    case "languagessecondary":
                        formatNumberScale = "0";
                        caption = "Books by Secondary Language";
                        break;
                    case "languagesoriginal":
                        formatNumberScale = "0";
                        caption = "Books by Original Language";
                        break;
                    case "deweys":
                        formatNumberScale = "0";
                        caption = "Books by Category";
                        break;
                    case "formats":
                        formatNumberScale = "0";
                        caption = "Books by Binding";
                        break;
                    case "contributorstop":
                        formatNumberScale = "0";
                        caption = "Books by Top Contributors";
                        break;
                    case "contributorstop":
                        formatNumberScale = "0";
                        caption = "Contributors by Role";
                        break;
                    case "datesoriginal":
                        formatNumberScale = "0";
                        caption = "Books by Original Publication Date";
                        break;
                    case "datespublication":
                        formatNumberScale = "0";
                        caption = "Books by Publication Date";
                        break;
                    case "lexile":
                        formatNumberScale = "0";
                        caption = "Books by Lexile Grade Level";
                        subcaption = "Taken from Common Core State Standards for English, Language Arts, Appendix A (Additional Information), NGA and CCSSO, 2012";                        break;
                        break;
                    case "tag":
                        formatNumberScale = "0";
                        caption = "Books by Tag";
                        break;
                }
                $scope.getChartData(view, caption, subcaption, labelDisplay, formatNumberScale, numberSuffix, decimals);
            }
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
                $scope.setStatView($scope.getParameterByName("statview", "general"));
                $scope.setStatSubView($scope.getParameterByName("statsubview", ""));
                $scope.removeFromLoading(loadingName)
            }, function(response) {
                $mdToast.showSimple("Failed to get list of libraries");
                $scope.removeFromLoading(loadingName);
            });
        };
        $scope.updateLibraries();
        $scope.chooseLibrary = function($ev) {
            $scope.showLibraryChooserDialog($ev, $scope, true)
        }
        $scope.$watch('output', function() {
            $scope.updateDimensions();
        });
    });
