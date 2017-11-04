angular.module('libraryOrganizer')
    .filter('boolFormat', function() {
        return function(x) {
            return x ? '' : 'Not';
        };
    })
    .directive('pagesInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    mCtrl.$setValidity('pagesInput', !isNaN(value) && value > 0 && value <= scope.countPages());
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('positiveIntegerInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    mCtrl.$setValidity('positiveIntegerInput', !isNaN(value) && value > 0 && parseInt(value) == value);
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('deweyInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    if (isNaN(value) && value.toUpperCase() != 'FIC') {
                        mCtrl.$setValidity('deweyInput', false);
                    } else {
                        mCtrl.$setValidity('deweyInput', value.toUpperCase() == 'FIC' || (!isNaN(value) && value >= 0 && value < 1000));
                    }
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('nonNegativeIntegerInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    mCtrl.$setValidity('nonNegativeIntegerInput', !isNaN(value) && value >= 0 && parseInt(value) == value);
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('nonNegativeNumberInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    mCtrl.$setValidity('nonNegativeNumberInput', !isNaN(value) && value >= 0);
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('yearInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    mCtrl.$setValidity('yearInput', !isNaN(value) && value.length == 4 && value >= 0 && parseInt(value) == value);
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('isbnInput', function() {
        return {
            require: 'ngModel',
            link: function(scope, element, attr, mCtrl) {
                function validation(value) {
                    if (value.length == 0) {
                        return true;
                    }
                    valid = false;
                    value = value.replace(/[^\dX]/gi, '');
                    if (value.length == 10) {
                        var chars = value.split('');
                        if (chars[9].toUpperCase() == 'X') {
                            chars[9] = 10;
                        }
                        var sum = 0;
                        for (var i = 0; i < chars.length; i++) {
                            sum += ((10 - i) * parseInt(chars[i]));
                        }
                        valid = (sum % 11 == 0);
                    } else if (value.length == 13) {
                        var chars = value.split('');
                        var sum = 0;
                        for (var i = 0; i < chars.length; i++) {
                            if (i % 2 == 0) {
                                sum += parseInt(chars[i]);
                            } else {
                                sum += parseInt(chars[i]) * 3;
                            }
                        }
                        valid = (sum % 10 == 0);
                    }
                    mCtrl.$setValidity('isbnInput', valid);
                    return value;
                }
                mCtrl.$parsers.push(validation);
            }
        }
    })
    .directive('appUploadFile', function() {
        var directive = {
            template: '<input id="fileInput" type="file" accept=".csv" class="ng-hide"><md-button id="uploadButton" class="icon import-icon" aria-label="attach_file"></md-button><md-input-container md-no-float ng-hide="true"><input id="textInput" ng-model="fileName" type="text" placeholder="No file chosen" ng-readonly="true"></md-input-container>',
            link: function(scope, element, attrs) {
                var input = $(element[0].querySelector('#fileInput'));
                var button = $(element[0].querySelector('#uploadButton'));
                var textInput = $(element[0].querySelector('#textInput'));

                if (input.length && button.length && textInput.length) {
                    button.click(function(e) {
                        input.click();
                    });
                    textInput.click(function(e) {
                        input.click();
                    });
                }

                input.on('change', function(e) {
                    var files = e.target.files;
                    if (files[0]) {
                        scope.upload(files[0]);
                    }
                });
            }
        };
        return directive;
    })