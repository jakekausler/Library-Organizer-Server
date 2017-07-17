angular.module('libraryOrganizer', ['ngMaterial', 'ng-fusioncharts', 'multiselect-searchtree'])
    .config(function($mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('indigo')
            .accentPalette('pink')
            .warnPalette('red')
            .backgroundPalette('indigo');
    })
    .controller('libraryOrganizerController', function($scope, $http, $timeout, $mdSidenav, $mdDialog) {
        $scope.getParameters = function() {
            var h = window.location.hash.slice(1);
            if (!h) {
                return {};
            }
            var hash;
            if (h.includes('%3F')) {
                hash = h.split('%3F');
            } else {
                hash = h.split('?');
            }
            var p = {};
            var ps = hash[1].split("&");
            for (v in ps) {
                var pm = ps[v].split("=");
                p[pm[0]] = pm[1];
            }
            return p;
        }
        $scope.parameters = $scope.getParameters();
        $scope.setParameters = function(params) {
            var h = window.location.hash.slice(1);
            if (!h) {
                h = '#state?';
            }
            var hash;
            if (h.includes('%3F')) {
                hash = h.split('%3F');
            } else {
                hash = h.split('?');
            }
            var p = {};
            var ps = hash[1].split("&");
            for (v in ps) {
                if (ps[v]) {
                    var pm = ps[v].split("=");
                    p[pm[0]] = pm[1];
                }
            }
            for (key in params) {
                p[key] = params[key];
            }
            var s = [];
            for (m in p) {
                s.push( m + "=" + p[m])
            }
            var qs = '?'+s.join('&');
            var newhash = hash[0]+qs
            window.location.hash = newhash
        }
        $scope.getParameterByName = function(name, def) {
            if ($scope.parameters[name]) {
                return $scope.parameters[name];
            }
            return def;
        }
        $scope.getSettingByName = function(name, callback) {
            $http({
                url: '/getsetting',
                method: 'POST',
                data: name
            }).then(function(response) {
                callback(response.data);
            });
        }
        $scope.display = $scope.getParameterByName("display", "grid");
        $scope.username = "";
        $scope.lastRecievedTime = new Date().getTime();
        $scope.parseFloat = function(v) {
            return parseFloat(v);
        }
        $scope.round = function(v, d) {
            return Math.round10(v, d)
        }
        $scope.showBookDialog = function(ev, book, vm, viewType) {
            $mdDialog.show({
                controller: 'viewController',
                templateUrl: 'web/app/main/bookviews/viewer/viewdialog.html',
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
        }
        $scope.settings = {};
        $scope.openSettings = function(ev, vm) {
            $mdDialog.show({
                controller: 'settingsController',
                templateUrl: 'web/app/main/settings/settingsdialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false,
                locals: {
                    vm: $scope
                }
            })
        }
        $scope.displayGrid = function() {
            $scope.setParameters({'display': 'grid'})
            $scope.display = 'grid';
        }
        $scope.displayStats = function() {
            $scope.setParameters({'display': 'stats'})
            $scope.display = 'stats';
        }
        $scope.displayShelves = function() {
            $scope.setParameters({'display': 'shelves'})
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
        $scope.showEditDialog = function(ev, book, vm, viewType) {
            $mdDialog.show({
                controller: 'editorController',
                templateUrl: 'web/app/main/bookviews/editor/editordialog.html',
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
        }
})