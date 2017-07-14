angular.module('libraryOrganizer', ['ngMaterial', 'ng-fusioncharts'])
    .config(function($mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('indigo')
            .accentPalette('indigo')
            .warnPalette('red')
            .backgroundPalette('indigo');
    })
    .controller('libraryOrganizerController', function($scope, $http, $timeout, $mdSidenav, $mdDialog) {
        $scope.display = "grid";
        $scope.lastRecievedTime = new Date().getTime();
        $scope.parseFloat = function(v) {
            return parseFloat(v);
        }
        $scope.round = function(v, d) {
            return Math.round10(v, d)
        }
        $scope.showEditorDialog = function(ev, book, vm, viewType) {
            $mdDialog.show({
                controller: 'editorController',
                templateUrl: 'web/app/main/editor/editordialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false,
                locals: {
                    book: book,
                    $vm: vm,
                    viewType: viewType
                }
            })
        };
        $scope.displayGrid = function() {
            $scope.display = 'grid';
        }
        $scope.displayStats = function() {
            $scope.display = 'stats';
        }
        $scope.displayShelves = function() {
            $scope.display = 'shelves';
        }
        $scope.import = function(ev) {
            $mdDialog.show({
                controller: 'importController',
                templateUrl: 'web/app/main/grid/import/importdialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false
            });
        }
        $scope.export = function(ev) {
            $mdDialog.show({
                controller: 'exportController',
                templateUrl: 'web/app/main/grid/export/exportdialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false
            });
        };
        $scope.getCurrentDateString = function() {
            var currTime = new Date().getTime()
            if (currTime - $scope.lastRecievedTime > 1000) {
                $scope.lastRecievedTime = currTime;
            }
            return $scope.lastRecievedTime;
        }
        $scope.OwnedLibraries = [];
        $scope.updateLibraries = function() {
            $http.get('/ownedlibraries', {}.then(function(response) {
                    $scope.libraries = response.data;
                });
            });
        $scope.updateLibraries();
    })
