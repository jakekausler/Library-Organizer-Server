angular.module('libraryOrganizer', ['ngMaterial', 'ng-fusioncharts', 'multiselect-searchtree'])
    .config(function($mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('indigo')
            .accentPalette('indigo')
            .warnPalette('red')
            .backgroundPalette('indigo');
    })
    .controller('libraryOrganizerController', function($scope, $http, $timeout, $mdSidenav, $mdDialog) {
        $scope.display = "grid";
        $scope.username = "";
        $scope.lastRecievedTime = new Date().getTime();
        $scope.parseFloat = function(v) {
            return parseFloat(v);
        }
        $scope.round = function(v, d) {
            return Math.round10(v, d)
        }
        $scope.showBookDialog = function(ev, book, vm, viewType) {
            console.log(book.library)
            if ((book.library.permissions&4)==4) {
                $scope.showEditorDialog(ev, book, vm, viewType);
            } else if ((book.library.permissions&2)==2) {
                $scope.showCheckOutDialog(ev, book, vm, viewType);
            } else if ((book.library.permissions&1)==1) {
                $scope.showViewDialog(ev, book, vm, viewType);
            }
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
                    viewType: viewType,
                    username: $scope.username
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
        $scope.updateUsername = function() {
            $http.get('/username', {}).then(function(response) {
                $scope.username = response.data;
            })
        }
        $scope.updateUsername()
        $scope.showLibraryChooserDialog = function(ev, vm, multiselect) {
            $mdDialog.show({
                controller: 'librarychooserController',
                templateUrl: 'web/app/main/librarychooser/librarychooser.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false,
                locals: {
                    vm: vm,
                    multiselect: multiselect
                }
            })
        };
    })
