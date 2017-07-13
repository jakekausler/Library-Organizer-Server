angular.module('libraryOrganizer')
.controller('importController', function ($scope, $mdDialog) {
	$scope.cancel = function() {
		$mdDialog.cancel()
	};
	$scope.upload = function(file) {
		var formData = new FormData();
		formData.append('file', file);
		$.ajax({
			url: '/import',
			method: 'POST',
			data: formData,
			contentType: false,
			processData: false
		});
	};
});