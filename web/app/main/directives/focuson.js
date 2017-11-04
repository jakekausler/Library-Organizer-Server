angular.module('libraryOrganizer')
    .directive('focusOn', function($timeout) {
        return {
            restrict: 'A',
            link: function($scope, $element, $attr) {
                $scope.$watch($attr.focusOn, function(_focusVal) {
                    $timeout(function() {
                        _focusVal ? $element[0].focus() :
                            $element[0].blur();
                    });
                });
            }
        }
    })