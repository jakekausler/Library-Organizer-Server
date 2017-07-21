angular.module('libraryOrganizer')
.controller('editorController', function($scope, $http, $mdDialog, book, $vm, viewType, username) {
	$scope.book = angular.copy(book);
	if (!$scope.book.bookid && !$scope.book.title) {
		$vm.getSettingByName('Title', function(value) {
			$scope.book.title = value;
		});
		$vm.getSettingByName('Subtitle', function(value) {
			$scope.book.subtitle = value;
		});
		$vm.getSettingByName('Series', function(value) {
			$scope.book.series = value;
		});
		$vm.getSettingByName('Volume', function(value) {
			$scope.book.volume = value;
		});
		$vm.getSettingByName('ImageURL', function(value) {
			$scope.book.imageurl = value;
		});
		$vm.getSettingByName('Publisher', function(value) {
			$scope.book.publisher.publisher = value;
		});
		$vm.getSettingByName('City', function(value) {
			$scope.book.publisher.city = value;
		});
		$vm.getSettingByName('State', function(value) {
			$scope.book.publisher.state = value;
		});
		$vm.getSettingByName('Country', function(value) {
			$scope.book.publisher.country = value;
		});
		$vm.getSettingByName('Originally Published', function(value) {
			$scope.book.originallypublished = value;
		});
		$vm.getSettingByName('Edition Published', function(value) {
			$scope.book.editionpublished = value;
		});
		$vm.getSettingByName('Dewey', function(value) {
			$scope.book.dewey = value;
		});
		$vm.getSettingByName('Format', function(value) {
			$scope.book.format = value;
		});
		$vm.getSettingByName('Pages', function(value) {
			$scope.book.pages = value;
		});
		$vm.getSettingByName('Width', function(value) {
			$scope.book.width = value;
		});
		$vm.getSettingByName('Height', function(value) {
			$scope.book.height = value;
		});
		$vm.getSettingByName('Depth', function(value) {
			$scope.book.depth = value;
		});
		$vm.getSettingByName('Weight', function(value) {
			$scope.book.weight = value;
		});
		$vm.getSettingByName('Primary Language', function(value) {
			$scope.book.primarylanguage = value;
		});
		$vm.getSettingByName('Secondary Language', function(value) {
			$scope.book.secondarylanguage = value;
		});
		$vm.getSettingByName('Original Language', function(value) {
			$scope.book.originallanguage = value;
		});
		$vm.getSettingByName('Owned', function(value) {
			$scope.book.isowned = (value=="True");
		});
		$vm.getSettingByName('Read', function(value) {
			$scope.book.isread = (value=="True");
		});
		$vm.getSettingByName('Reference', function(value) {
			$scope.book.isreference = (value=="True");
		});
		$vm.getSettingByName('Shipping', function(value) {
			$scope.book.isshipping = (value=="True");
		});
		$vm.getSettingByName('Reading', function(value) {
			$scope.book.isreading = (value=="True");
		});
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
        $http({
        	url: '/libraries'
        }).then(function(response) {
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
			url: '/information/publishers',
			method: 'GET',
		}).then(function(response){
			$scope.publishers = response.data;
		});
	}
	$scope.updatePublishers();
	$scope.updateCities = function() {
		$http({
			url: '/information/cities',
			method: 'GET'
		}).then(function(response){
			$scope.cities = response.data;
		});
	}
	$scope.updateCities();
	$scope.updateStates = function() {
		$http({
			url: '/information/states',
			method: 'GET'
		}).then(function(response){
			$scope.states = response.data;
		});
	}
	$scope.updateStates();
	$scope.updateCountries = function() {
		$http({
			url: '/information/countries',
			method: 'GET'
		}).then(function(response){
			$scope.countries = response.data;
		});
	}
	$scope.updateCountries();
	$scope.updateSeries = function() {
		$http({
			url: '/information/series',
			method: 'GET'
		}).then(function(response){
			$scope.series = response.data;
		});
	}
	$scope.updateSeries();
	$scope.updateFormats = function() {
		$http({
			url: '/information/formats',
			method: 'GET'
		}).then(function(response){
			$scope.formats = response.data;
		});
	}
	$scope.updateFormats();
	$scope.updateLanguages = function() {
		$http({
			url: '/information/languages',
			method: 'GET'
		}).then(function(response){
			$scope.languages = response.data;
		});
	}
	$scope.updateLanguages();
	$scope.updateRoles = function() {
		$http({
			url: '/information/roles',
			method: 'GET'
		}).then(function(response){
			$scope.roles = response.data;
		});
	}
	$scope.updateRoles();
	$scope.updateDeweys = function() {
		$http({
			url: '/information/deweys',
			method: 'GET'
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
		var method = book.id ? 'PUT':'POST';
		$http({
			url: '/books',
			method: method,
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
				url: '/books',
				method: 'DELETE',
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