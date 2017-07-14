angular.module('libraryOrganizer')
.controller('editorController', function($scope, $http, $mdDialog, book, $vm, viewType) {
	$scope.publishers = [];
	$scope.cities = [];
	$scope.states = [];
	$scope.countries = [];
	$scope.series = [];
	$scope.formats = [];
	$scope.languages = [];
	$scope.roles = [];
	$scope.deweys = [];
	$scope.updatePublishers = function() {
		$http({
			url: '/publishers',
			method: 'POST',
		}).then(function(response){
			$scope.publishers = response.data;
		});
	}
	$scope.updatePublishers();
	$scope.updateCities = function() {
		$http({
			url: '/cities',
			method: 'POST'
		}).then(function(response){
			$scope.cities = response.data;
		});
	}
	$scope.updateCities();
	$scope.updateStates = function() {
		$http({
			url: '/states',
			method: 'POST'
		}).then(function(response){
			$scope.states = response.data;
		});
	}
	$scope.updateStates();
	$scope.updateCountries = function() {
		$http({
			url: '/countries',
			method: 'POST'
		}).then(function(response){
			$scope.countries = response.data;
		});
	}
	$scope.updateCountries();
	$scope.updateSeries = function() {
		$http({
			url: '/series',
			method: 'POST'
		}).then(function(response){
			$scope.series = response.data;
		});
	}
	$scope.updateSeries();
	$scope.updateFormats = function() {
		$http({
			url: '/formats',
			method: 'POST'
		}).then(function(response){
			$scope.formats = response.data;
		});
	}
	$scope.updateFormats();
	$scope.updateLanguages = function() {
		$http({
			url: '/languages',
			method: 'POST'
		}).then(function(response){
			$scope.languages = response.data;
		});
	}
	$scope.updateLanguages();
	$scope.updateRoles = function() {
		$http({
			url: '/roles',
			method: 'POST'
		}).then(function(response){
			$scope.roles = response.data;
		});
	}
	$scope.updateRoles();
	$scope.updateDeweys = function() {
		$http({
			url: '/deweys',
			method: 'POST'
		}).then(function(response){
			$scope.deweys = response.data;
		});
	}
	$scope.updateDeweys();
	$scope.book = angular.copy(book);
	$scope.newContributor = {
		role: '',
		name: {
			first: '',
			middles: '',
			last: ''
		}
	};
	$scope.oldUrl = $scope.book.imageurl;
	$scope.pastingurl = false;
	$scope.save = function(book) {
		book.volume = parseFloat(book.volume);
		book.edition = parseInt(book.edition);
		book.pages = parseInt(book.pages);
		book.width = parseInt(book.width);
		book.height = parseInt(book.height);
		book.depth = parseInt(book.depth);
		book.weight = parseFloat(book.weight);
		book.originallypublished = book.originallypublished+'-01-01';
		book.editionpublished = book.editionpublished+'-01-01';
		$http({
			url: '/savebook',
			method: 'POST',
			data: JSON.stringify(book)
		}).then(function(response){
			if (viewType=='gridadd') {
				$vm.updateRecieved();
			} else if (viewType=='shelves') {
				$vm.updateCases();
			} else if (viewType=='grid') {
				$vm.$parent.$parent.updateRecieved();
			}
			$mdDialog.cancel();
		});
	}
	$scope.remove = function(book) {
		var confirm = $mdDialog.confirm()
			.title('Really Remove Book?')
			.textContent('Are you sure you want to remove this book? This cannot be undone.')
			.ariaLabel('Remove book')
			.ok('Yes')
			.cancel('Cancel');
		$mdDialog.show(confirm).then(function() {
			$http({
			url: '/deletebook',
			method: 'POST',
			data: book.bookid
		}).then(function(response) {
			if (viewType=='gridadd') {
				$vm.updateRecieved();
			} else if (viewType=='shelves') {
				$vm.updateCases();
			} else if (viewType=='grid') {
				$vm.$parent.$parent.updateRecieved();
			}
			$mdDialog.cancel();
		})
		}, function() {});
	}
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
	$scope.removeContributor = function(index) {
		$scope.book.contributors.splice(index, 1);
	}
	$scope.addContributor = function() {
		$scope.book.contributors.push(angular.copy($scope.newContributor));
		$scope.newContributor = {
			role: '',
			name: {
				first: '',
				middles: '',
				last: ''
			}
		};
	}
	$scope.cancelPastingURL = function() {
		$scope.pastingurl = false;
		$scope.book.imageurl = $scope.oldUrl;
	}
	$scope.query = function(arr, str) {
		arr.push(str);
		var results = str ? arr.filter(function(s) {
			return (angular.lowercase(s).indexOf(angular.lowercase(str)) !== -1);
		}) : arr;
		return results;
	}
});