angular.module('libraryOrganizer')
    .controller('importController', function($scope, $mdDialog, vm) {
        $scope.vm = vm;
        $scope.cancel = function() {
            $mdDialog.cancel()
        };
        $scope.upload = function(file) {
            var loadingName = $scope.vm.guid();
            $scope.vm.addToLoading(loadingName)
            var formData = new FormData();
            formData.append('file', file);
            $.ajax({
                url: '/books/books',
                method: 'POST',
                data: formData,
                contentType: false,
                processData: false,
                success: function(response) {
                    $mdToast.showSimple("Successfully imported books")
                    $scope.vm.removeFromLoading(loadingName);
                },
                error: function(xhr, status, error) {
                    $mdToast.showSimple("Failed to import books")
                    $scope.vm.removeFromLoading(loadingName);
                }
            });
        };
    });