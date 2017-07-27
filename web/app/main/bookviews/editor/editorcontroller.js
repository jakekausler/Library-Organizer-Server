angular.module('libraryOrganizer')
.controller('editorController', function($scope, $http, $mdDialog, $mdToast, book, $vm, viewType, username) {
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
		$vm.getSettingByName('Lexile', function(value) {
			$scope.book.lexile = $vm.convertToLexile(value);
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
	if (!isNaN($scope.book.lexile)) {
		$scope.book.lexile = $vm.convertToLexile($scope.book.lexile);
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
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
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
            $vm.removeFromLoading(loadingName);
        });
    };
	$scope.updatePublishers = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/publishers',
			method: 'GET',
		}).then(function(response){
			$scope.publishers = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updatePublishers();
	$scope.updateCities = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/cities',
			method: 'GET'
		}).then(function(response){
			$scope.cities = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateCities();
	$scope.updateStates = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/states',
			method: 'GET'
		}).then(function(response){
			$scope.states = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateStates();
	$scope.updateCountries = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/countries',
			method: 'GET'
		}).then(function(response){
			$scope.countries = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateCountries();
	$scope.updateSeries = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/series',
			method: 'GET'
		}).then(function(response){
			$scope.series = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateSeries();
	$scope.updateFormats = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/formats',
			method: 'GET'
		}).then(function(response){
			$scope.formats = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateFormats();
	$scope.updateLanguages = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/languages',
			method: 'GET'
		}).then(function(response){
			$scope.languages = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateLanguages();
	$scope.updateRoles = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/roles',
			method: 'GET'
		}).then(function(response){
			$scope.roles = response.data;
            $vm.removeFromLoading(loadingName);
		});
	}
    $scope.updateRoles();
	$scope.updateDeweys = function() {
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$http({
			url: '/information/deweys',
			method: 'GET'
		}).then(function(response){
			for (i in response.data) {
				$scope.deweys.push(response.data[i].dewey);
				$scope.genres[response.data[i].dewey] = response.data[i].genre;
			}
            $vm.removeFromLoading(loadingName);
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
        var loadingName = $vm.guid();
        $vm.addToLoading(loadingName)
		$scope.convertIsbn();
		book.volume = parseFloat(book.volume);
		book.edition = parseInt(book.edition);
		book.pages = parseInt(book.pages);
		book.width = parseInt(book.width);
		book.height = parseInt(book.height);
		book.depth = parseInt(book.depth);
		book.weight = parseFloat(book.weight);
		book.lexile = parseInt($vm.convertFromLexile(book.lexile));
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
				$vm.updateRecieved();
			}
			if (method == "PUT") {
				$mdToast.showSimple("Successfully added book")
			} else {
				$mdToast.showSimple("Successfully saved book")
			}
            $vm.removeFromLoading(loadingName);
			$mdDialog.cancel();
		}, function(response) {
			if (method == "PUT") {
				$mdToast.showSimple("Failed to add book")
			} else {
				$mdToast.showSimple("Failed to save book")
			}
            $vm.removeFromLoading(loadingName);
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
	        var loadingName = $vm.guid();
	        $vm.addToLoading(loadingName)
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
				$mdToast.showSimple("Successfully removed book")
            	$vm.removeFromLoading(loadingName);
				$mdDialog.cancel();
			}, function(response) {
				$mdToast.showSimple("Failed to remove book")
            	$vm.removeFromLoading(loadingName);
			})
		});
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
	$scope.getGenre = function() {
		if ($scope.book.dewey == "FIC") {
			return 'Fiction';
		}
		return $scope.genres[$scope.book.dewey]?$scope.genres[$scope.book.dewey].replace(">", "\u003e"):'';
	}
    $scope.updateLibraries();
    $scope.convertIsbn = function() {
    	if ($scope.book.isbn.length == 10 && $scope.isValidIsbn($scope.book.isbn)) {
    		$scope.book.isbn = $scope.isbn10to13($scope.book.isbn);
    	}
    }
    $scope.isbn10to13 = function(isbn10) {
    	var chars = isbn10.split("");
	    chars.unshift("9", "7", "8");
	    chars.pop();
	    var i = 0;
	    var sum = 0;
	    for (i = 0; i < 12; i += 1) {
	        sum += chars[i] * ((i % 2) ? 3 : 1);
	    }
	    var check_digit = (10 - (sum % 10)) % 10;
	    chars.push(check_digit);
	    var isbn13 = chars.join("");
	    return isbn13;
    }
    $scope.isValidIsbn = function(isbn) {
    	if (isbn.length == 0) {
			return true;
		}
		valid = false;
		isbn = isbn.replace(/[^\dX]/gi, '');
		if(isbn.length == 10) {
			var chars = isbn.split('');
			if(chars[9].toUpperCase() == 'X') {
				chars[9] = 10;
			}
			var sum = 0;
			for(var i = 0; i < chars.length; i++) {
				sum += ((10-i) * parseInt(chars[i]));
			}
			valid = (sum % 11 == 0);
		} else if(isbn.length == 13) {
			var chars = isbn.split('');
			var sum = 0;
			for (var i = 0; i < chars.length; i++) {
				if(i % 2 == 0) {
					sum += parseInt(chars[i]);
				} else {
					sum += parseInt(chars[i]) * 3;
				}
			}
			valid = (sum % 10 == 0);
		}
		return valid;
    }
});