angular.module('libraryOrganizer')
.controller('importController', function ($scope, $mdDialog) {
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
	$scope.upload = function(file) {
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
			},
			error: function(xhr, status, error) {
				$mdToast.showSimple("Failed to import books")
			}
		});
	};
});