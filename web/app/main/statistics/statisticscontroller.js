angular.module('libraryOrganizer')
    .controller('statisticsController', function($scope, $http) {
        $scope.statView = 'general';
        $scope.statSubView = 'bycounts';
        $scope.libraryids = ['1'];
        $scope.stringLibraryIds = function() {
            return $scope.libraryids.join(',')
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
        $scope.updateDimensions();
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
            console.log(view)
            $scope.statSubView = view;
        }
    });
