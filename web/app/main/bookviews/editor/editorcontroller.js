angular.module('libraryOrganizer')
.controller('editorController', function($scope, $http, $mdDialog, book, $vm, viewType, username) {
	$scope.book = angular.copy(book);
	if (!$scope.book.title) {
		$vm.getSettingByName('Title', function(value) {
			$scope.book.title = value;
		});
	}
	if (!$scope.book.subtitle) {
		$vm.getSettingByName('Subtitle', function(value) {
			$scope.book.subtitle = value;
		});
	}
	if (!$scope.book.series) {
		$vm.getSettingByName('Series', function(value) {
			$scope.book.series = value;
		});
	}
	if (!$scope.book.volume) {
		$vm.getSettingByName('Volume', function(value) {
			$scope.book.volume = value;
		});
	}
	if (!$scope.book.imageurl) {
		$vm.getSettingByName('ImageURL', function(value) {
			$scope.book.imageurl = value;
		});
	}
	if (!$scope.book.publisher.publisher) {
		$vm.getSettingByName('Publisher', function(value) {
			$scope.book.publisher.publisher = value;
		});
	}
	if (!$scope.book.publisher.city) {
		$vm.getSettingByName('City', function(value) {
			$scope.book.publisher.city = value;
		});
	}
	if (!$scope.book.publisher.state) {
		$vm.getSettingByName('State', function(value) {
			$scope.book.publisher.state = value;
		});
	}
	if (!$scope.book.publisher.country) {
		$vm.getSettingByName('Country', function(value) {
			$scope.book.publisher.country = value;
		});
	}
	if (!$scope.book.originallypublished) {
		$vm.getSettingByName('Originally Published', function(value) {
			$scope.book.originallypublished = value;
		});
	}
	if (!$scope.book.editionpublished) {
		$vm.getSettingByName('Edition Published', function(value) {
			$scope.book.editionpublished = value;
		});
	}
	if (!$scope.book.dewey) {
		$vm.getSettingByName('Dewey', function(value) {
			$scope.book.dewey = value;
		});
	}
	if (!$scope.book.format) {
		$vm.getSettingByName('Format', function(value) {
			$scope.book.format = value;
		});
	}
	if (!$scope.book.pages) {
		$vm.getSettingByName('Pages', function(value) {
			$scope.book.pages = value;
		});
	}
	if (!$scope.book.width) {
		$vm.getSettingByName('Width', function(value) {
			$scope.book.width = value;
		});
	}
	if (!$scope.book.height) {
		$vm.getSettingByName('Height', function(value) {
			$scope.book.height = value;
		});
	}
	if (!$scope.book.depth) {
		$vm.getSettingByName('Depth', function(value) {
			$scope.book.depth = value;
		});
	}
	if (!$scope.book.weight) {
		$vm.getSettingByName('Weight', function(value) {
			$scope.book.weight = value;
		});
	}
	if (!$scope.book.primarylanguage) {
		$vm.getSettingByName('Primary Language', function(value) {
			$scope.book.primarylanguage = value;
		});
	}
	if (!$scope.book.secondarylanguage) {
		$vm.getSettingByName('Secondary Language', function(value) {
			$scope.book.secondarylanguage = value;
		});
	}
	if (!$scope.book.originallanguage) {
		$vm.getSettingByName('Original Language', function(value) {
			$scope.book.originallanguage = value;
		});
	}
	if (!$scope.book.isowned) {
		$vm.getSettingByName('Owned', function(value) {
			$scope.book.isowned = (value=="True");
		});
	}
	if (!$scope.book.isread) {
		$vm.getSettingByName('Read', function(value) {
			$scope.book.isread = (value=="True");
		});
	}
	if (!$scope.book.isreference) {
		$vm.getSettingByName('Reference', function(value) {
			$scope.book.isreference = (value=="True");
		});
	}
	if (!$scope.book.isshipping) {
		$vm.getSettingByName('Shipping', function(value) {
			$scope.book.isshipping = (value=="True");
		});
	}
	if (!$scope.book.isreading) {
		$vm.getSettingByName('Reading', function(value) {
			$scope.book.isreading = (value=="True");
		});
	}
	if (!$scope.book.edition) {
		$vm.getSettingByName('Edition', function(value) {
			$scope.book.edition = value;
		});
	}
	$scope.publishers = [];
	$scope.cities = [];
	$scope.states = [];
	$scope.countries = [];
	$scope.series = [];
	$scope.formats = [];
	$scope.languages = [];
	$scope.roles = [];
	$scope.deweys = [];
	$scope.genres = {};
	$scope.libraries = [];
	$scope.deweySearchText = $scope.book.dewey;
	$scope.updateLibraries = function() {
        $http.get('/libraries', {}).then(function(response) {
            for (l in response.data) {
	            if ((response.data[l].permissions&4)==4) {
	            	$scope.libraries.push(response.data[l])
	            }
	        }
            for (l in $scope.libraries) {
            	if ($scope.libraries[l].owner != username) {
            		$scope.libraries[l].display = $scope.libraries[l].name + " (" + $scope.libraries[l].owner + ")"
            	} else {
            		if (!$scope.book.library.id) {
            			$scope.book.library.id = $scope.libraries[l].id;
            			$scope.book.library.name = $scope.libraries[l].name;
            			$scope.book.library.permissions = $scope.libraries[l].permissions;
            			$scope.book.library.owner = $scope.libraries[l].owner;
            		}
	            	$scope.libraries[l].display = $scope.libraries[l].name
            	}
            }
        });
    };
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
			for (i in response.data) {
				$scope.deweys.push(response.data[i].dewey);
				$scope.genres[response.data[i].dewey] = response.data[i].genre;
			}
		});
	}
	$scope.updateDeweys();
	$scope.newContributor = {
		role: 'Role',
		name: {
			first: '',
			middles: '',
			last: 'Last'
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
		book.series = book.series?book.series:$scope.seriesSearchText;
		book.publisher.publisher = book.publisher.publisher?book.publisher.publisher:$scope.PublisherSearchText;
		book.publisher.city = book.publisher.city?book.publisher.city:$scope.CitySearchText;
		book.publisher.state = book.publisher.state?book.publisher.state:$scope.stateSearchText;
		book.publisher.country = book.publisher.country?book.publisher.country:$scope.countrySearchText;
		book.dewey = book.dewey?book.dewey:$scope.deweySearchText;
		book.format = book.format?book.format:$scope.formatSearchText;
		book.primarylanguage = book.primarylanguage?book.primarylanguage:$scope.primaryLanguageSearchText;
		book.secondarylanguage = book.secondarylanguage?book.secondarylanguage:$scope.secondaryLanguageSearchText;
		book.originallanguage = book.originallanguage?book.originallanguage:$scope.originalLanguageSearchText;
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
			} else if (viewType=='scanadd') {
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
			role: 'Role',
			name: {
				first: '',
				middles: '',
				last: 'Last'
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
	$scope.log = function(item) {
		console.log(item)
	}
	$scope.updateLibraries();
	$scope.getGenre = function() {
		if ($scope.deweySearchText == "FIC") {
			return 'Fiction';
		}
		return $scope.genres[$scope.deweySearchText]?$scope.genres[$scope.deweySearchText].replace(">", "\u003e"):'';
	}
});