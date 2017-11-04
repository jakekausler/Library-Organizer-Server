angular.module('libraryOrganizer')
    .directive('scan', function() {
        return function(scope, element, attrs) {
            element.bind("keydown keypress", function(event) {
                if (event.which === 13) {
                    scope.$apply(function() {
                        scope.$eval(attrs.scan);
                    });
                    event.preventDefault();
                }
            });
        };
    });