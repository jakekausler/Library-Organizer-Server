angular.module('libraryOrganizer')
    .controller('statisticsController', function($scope, $http) {
        $scope.statView = 'general';
        $scope.statSubView = 'bycounts';
        $scope.libraryids = [];
        $scope.dimensions = {};
        $scope.updateDimensions = function() {
            $http({
                url: '/dimensions',
                method: 'GET',
                params: {
                    libraryids: $scope.stringLibraryIds()
                }
            }).then(function(response) {
                console.log(response.data);
                $scope.dimensions = response.data;
            });
        }
        $scope.updateDimensions();
        $scope.stringLibraryIds = function() {
            return $scope.libraryids.join(',')
        }
    });
