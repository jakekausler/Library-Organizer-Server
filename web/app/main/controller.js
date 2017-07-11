var app = angular.module('libraryOrganizer', ['ngMaterial', 'ng-fusioncharts']).config(function($mdThemingProvider) {
	$mdThemingProvider.theme('default')
		.primaryPalette('indigo')
		.accentPalette('indigo')
		.warnPalette('red')
		.backgroundPalette('indigo');
});
app.controller('libraryOrganizerController', function($scope, $http, $timeout, $mdSidenav, $mdDialog) {
	$scope.display = "grid";

	$scope.statView = 'general';

	$scope.editingShelves = false;
	$scope.shelfSearchString = "";

	//Grid view functions
	$scope.sort = "dewey";
	$scope.numberToGet = 50;
	$scope.page = 1;
	$scope.numberOfBooks=0;
	$scope.fromdewey = "0";
	$scope.todewey = 'FIC';
	$scope.filter = "";
	$scope.read = 'both';
	$scope.reference = 'both';
	$scope.owned = 'yes';
	$scope.loaned = 'no';
	$scope.shipping = 'no';
	$scope.reading = 'no';

	$scope.publishers;
	$scope.cities;
	$scope.states;
	$scope.countries;
	$scope.series;
	$scope.formats;
	$scope.languages;
	$scope.roles;
	$scope.deweys;

	$scope.lastRecievedTime = new Date().getTime();

	$scope.updateRecieved = function() {
		console.log($scope.page);
		$http.get('/books', {
			params: {
				sortmethod: $scope.sort,
				numbertoget: $scope.numberToGet,
				page: $scope.page,
				fromdewey: $scope.fromdewey.toUpperCase(),
				todewey: $scope.todewey.toUpperCase(),
				text: $scope.filter,
				isread: $scope.read,
				isreference: $scope.reference,
				isowned: $scope.owned,
				isloaned: $scope.loaned,
				isshipping: $scope.shipping,
				isreading: $scope.reading
			}
		}).then(function(response){
			$scope.books = response.data.books;
			for (b in $scope.books) {
				for (c in $scope.books[b].contributors) {
					$scope.books[b].contributors[c].editing = false;
				}
			}
			$scope.numberOfBooks = response.data.numbooks;
		});
	};
	$scope.updateRecieved();
	$scope.updatePublishers = function() {
		$http({
			url: '/publishers',
			method: 'POST',
			// data: JSON.stringify({
			//	bookids: []
			// })
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
	$scope.previousPage = function() {
		$scope.page -= 1;
		$scope.updateRecieved();
	};
	$scope.nextPage = function() {
		$scope.page += 1;
		$scope.updateRecieved();
	};
	$scope.filterFunction = function(value, index, array) {
		formatDewey = function(dewey) {
			if (isNaN(dewey) && dewey.toUpperCase() !== 'FIC') {
				return 0;
			}
			if (isNaN(dewey) && dewey.toUpperCase() == 'FIC') {
				return 'FIC';
			}
			if (parseFloat(dewey) < 10) {
				dewey = '0'+parseFloat(dewey);
			}
			if (parseFloat(dewey) < 100) {
				dewey = '0'+parseFloat(dewey);
			}
			return dewey;
		};
		if (value.dewey < formatDewey($scope.fromdewey) || value.dewey > formatDewey($scope.todewey)) {
				return false;
		}
		if ($scope.read !== 'both') {
			if ($scope.read == 'yes') {
				if (!value.isread) {
					return false;
				}
			} else {
				if (value.isread) {
					return false;
				}
			}
		}
		if ($scope.reference !== 'both') {
			if ($scope.reference == 'yes') {
				if (!value.isreference) {
					return false;
				}
			} else {
				if (value.isreference) {
					return false;
				}
			}
		}
		if ($scope.owned !== 'both') {
			if ($scope.owned == 'yes') {
				if (!value.isowned) {
					return false;
				}
			} else {
				if (value.isowned) {
					return false;
				}
			}
		}
		if ($scope.loaned !== 'both') {
			if ($scope.loaned == 'yes') {
				if (!value.loanee.first && !value.loanee.middles && !value.loanee.last) {
					return false;
				}
			} else {
				if (value.loanee.first || value.loanee.middles || value.loanee.last) {
					return false;
				}
			}
		}
		if ($scope.shipping !== 'both') {
			if ($scope.shipping == 'yes') {
				if (!value.isshipping) {
					return false;
				}
			} else {
				if (value.isshipping) {
					return false;
				}
			}
		}
		if ($scope.reading !== 'both') {
			if ($scope.reading == 'yes') {
				if (!value.isreading) {
					return false;
				}
			} else {
				if (value.isreading) {
					return false;
				}
			}
		}
		if ($scope.filter !== "") {
			return value.title.toLowerCase().includes($scope.filter.toLowerCase()) || value.subtitle.toLowerCase().includes($scope.filter.toLowerCase()) || value.series.toLowerCase().includes($scope.filter.toLowerCase());
		}
		return true;
	};
	$scope.sortFunction = function(v1, v2) {
		if ($scope.sort == 'dewey') {
			v1 = v1.value;
			v2 = v2.value;
			if (!isNaN(v1) || (!isNaN(v2))) {
				return 0;
			}
			if (v1.dewey > v2.dewey) {
				return 1;
			} else if (v1.dewey < v2.dewey) {
				return -1;
			}
			var minname1 = "";
			var minname2 = "";
			for (i in v1.contributors) {
				contributor = v1.contributors[i];
				if (contributor.role == 'Author') {
					if (minname1 == "" || minname1 > contributor.name.last+contributor.name.first+contributor.name.middles) {
						minname1 = contributor.name.last+contributor.name.first+contributor.name.middles;
					}
				}
			}
			for (i in v2.contributors) {
				contributor = v2.contributors[i];
				if (contributor.role == 'Author') {
					if (minname2 == "" || minname2 > contributor.name.last+contributor.name.first+contributor.name.middles) {
						minname2 = contributor.name.last+contributor.name.first+contributor.name.middles;
					}
				}
			}
			if (minname1 > minname2) {
				return 1;
			} else if (minname1 < minname2) {
				return -1;
			}
			var titleChange = function(str) {
				if (str.startsWith("The ")) {
					return str.substring(4)+', The';
				}
				if (str.startsWith("An ")) {
					return str.substring(4)+', An';
				}
				if (str.startsWith("A ")) {
					return str.substring(4)+', A';
				}
				return str;
			};
			var series1 = titleChange(v1.series);
			var series2 = titleChange(v2.series);
			if (series1 > series2) {
				return 1;
			} else if (series1 < series2) {
				return -1;
			}
			if (v1.volume > v2.volume) {
				return 1;
			} else if (v1.volume < v2.volume) {
				return -1;
			}
			var title1 = titleChange(v1.title);
			var title2 = titleChange(v2.title);
			if (title1 > title2) {
				return 1;
			} else if (title1 < title2) {
				return -1;
			}
			var subtitle1 = titleChange(v1.subtitle);
			var subtitle2 = titleChange(v2.subtitle);
			if (subtitle1 > subtitle2) {
				return 1;
			} else if (subtitle1 < subtitle2) {
				return -1;
			}
			return 0;
		}
	};
	$scope.countPages = function() {
		if (isNaN($scope.numberToGet) || $scope.numberToGet <= 0) {
			return 0;
		}
		return Math.ceil($scope.numberOfBooks/$scope.numberToGet);
	};
	$scope.closeFiltersNav = function () {
		$mdSidenav('filterSideNav').close();
	};
	$scope.toggleFilters = function() {
		$mdSidenav('filterSideNav').open();
	};

	$scope.addBook = function(ev) {
		var book = {
			"bookid": "",
			"title": "",
			"subtitle": "",
			"originallypublished": "",
			"publisher": {
				"id": "",
				"publisher": "",
				"city": "",
				"state": "",
				"country": "",
				"parentcompany": ""
			},
			"isread": false,
			"isreference": false,
			"isowned": false,
			"isbn": "",
			"loanee": {
				"first": "",
				"middles": "",
				"last": ""
			},
			"dewey": "0",
			"pages": 0,
			"width": 0,
			"height": 0,
			"depth": 0,
			"weight": 0,
			"primarylanguage": "",
			"secondarylanguage": "",
			"originallanguage": "",
			"series": "",
			"volume": 0,
			"format": "",
			"edition": 0,
			"isreading": false,
			"isshipping": false,
			"imageurl": "",
			"spinecolor": "",
			"cheapestnew": 0,
			"cheapestused": 0,
			"editionpublished": "",
			"contributors": []
		}
		$scope.showEditorDialog(ev, book);
	}

	$scope.showEditorDialog = function(ev, book) {
		$scope.updatePublishers();
		$scope.updateCities();
		$scope.updateStates();
		$scope.updateCountries();
		$scope.updateSeries();
		$scope.updateFormats();
		$scope.updateLanguages();
		$scope.updateDeweys();
		$scope.updateRoles();
		$vm = $scope;
		$mdDialog.show({
			controller: function ($scope, $mdDialog) {
				$scope.book = angular.copy(book);
				$scope.publishers = $vm.publishers;
				$scope.cities = $vm.cities;
				$scope.states = $vm.states;
				$scope.countries = $vm.countries;
				$scope.formats = $vm.formats;
				$scope.series = $vm.series;
				$scope.deweys = $vm.deweys;
				$scope.languages = $vm.languages;
				$scope.roles = $vm.roles;
				$scope.publishersLength = $scope.publishers.length;
				$scope.citiesLength = $scope.cities.length;
				$scope.statesLength = $scope.states.length;
				$scope.countriesLength = $scope.countries.length;
				$scope.formatsLength = $scope.formats.length;
				$scope.seriesLength = $scope.series.length;
				$scope.deweysLength = $scope.deweys.length;
				$scope.languagesLength = $scope.languages.length;
				$scope.rolesLength = $scope.roles.length;
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
						$vm.updateRecieved();
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
						//todo remove book
						console.log(book);
					}, function() {});
				}
				$scope.cancel = function() {
					$mdDialog.cancel();
				};
				$scope.removeContributor = function(index) {
					$scope.book.contributors.splice(index, 1);
				}
				$scope.addContributor = function() {
					console.log($scope.newContributor);
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
				$scope.getCurrentDateString = function() {
					var currTime = new Date().getTime()
					if (currTime-$scope.lastRecievedTime>1000) {
						$scope.lastRecievedTime = currTime;
					}
					return $scope.lastRecievedTime;
				}
			},
			templateUrl: 'web/app/main/editordialog.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		})
	};

	$scope.displayGrid = function() {
		$scope.display = 'grid';
	}
	$scope.displayStats = function() {
		$scope.display = 'stats';
	}
	$scope.displayShelves = function() {
		$scope.display = 'shelves';
	}

	$scope.import = function(ev) {
		$vm = $scope;
		$mdDialog.show({
			controller: function ($scope, $mdDialog) {
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
			},
			templateUrl: 'web/app/main/importdialog.html',
			parent: angular.element(document.body),
			targetEvt: ev,
			clickOutsideToClose: true,
			fullscreen: false
		});
	}
	$scope.export = function() {
		$http.get('/exportbooks').then(function(data) {
			var anchor = angular.element('<a/>');
			anchor.attr({
				href: 'data:attachment/csv;charset=utf-8,' + encodeURI(data.data),
				target: '_blank',
				download: 'books.csv'
			})[0].click();
		});
		$http.get('/exportauthors').then(function(data) {
			var anchor = angular.element('<a/>');
			anchor.attr({
				href: 'data:attachment/csv;charset=utf-8,' + encodeURI(data.data),
				target: '_blank',
				download: 'authors.csv'
			})[0].click();
		});
	};
	$scope.setStatView = function(view) {
		$scope.statView = view;
	}
	//todo
	$scope.addShelf = function() {

	}
	//todo
	$scope.toggleEditShelves = function() {
		$scope.editingShelves = !$scope.editingShelves
	}
	//todo
	$scope.findBook = function() {
		
	}
	$scope.getCurrentDateString = function() {
		var currTime = new Date().getTime()
		if (currTime-$scope.lastRecievedTime>1000) {
			$scope.lastRecievedTime = currTime;
		}
		return $scope.lastRecievedTime;
	}

});
app.filter('boolFormat', function() {
	return function(x) {
		return x?'':'Not';
	};
});
app.directive('pagesInput', function() {
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
});
app.directive('positiveIntegerInput', function() {
	return {
		require: 'ngModel',
		link: function(scope, element, attr, mCtrl) {
			function validation(value) {
				mCtrl.$setValidity('positiveIntegerInput', !isNaN(value) && value > 0 && parseInt(value)==value);
				return value;
			}
			mCtrl.$parsers.push(validation);
		}
	}
});
app.directive('deweyInput', function() {
	return {
		require: 'ngModel',
		link: function(scope, element, attr, mCtrl) {
			function validation(value) {
				if (isNaN(value) && value.toUpperCase() != 'FIC') {
					mCtrl.$setValidity('deweyInput', false);
				} else {
					mCtrl.$setValidity('deweyInput', value.toUpperCase()=='FIC' || (!isNaN(value) && value >= 0 && value < 1000));
				}
				return value;
			}
			mCtrl.$parsers.push(validation);
		}
	}
});
app.directive('nonNegativeIntegerInput', function() {
	return {
		require: 'ngModel',
		link: function(scope, element, attr, mCtrl) {
			function validation(value) {
				mCtrl.$setValidity('nonNegativeIntegerInput', !isNaN(value) && value >= 0 && parseInt(value)==value);
				return value;
			}
			mCtrl.$parsers.push(validation);
		}
	}
});
app.directive('nonNegativeNumberInput', function() {
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
});
app.directive('yearInput', function() {
	return {
		require: 'ngModel',
		link: function(scope, element, attr, mCtrl) {
			function validation(value) {
				mCtrl.$setValidity('yearInput', !isNaN(value) && value.length == 4 && value >= 0 && parseInt(value)==value);
				return value;
			}
			mCtrl.$parsers.push(validation);
		}
	}
});
app.directive('isbnInput', function() {
	return {
		require: 'ngModel',
		link: function(scope, element, attr, mCtrl) {
			function validation(value) {
				valid = false;
				value = value.replace(/[^\dX]/gi, '');
				if(value.length == 10) {
					var chars = value.split('');
					if(chars[9].toUpperCase() == 'X') {
						chars[9] = 10;
					}
					var sum = 0;
					for(var i = 0; i < chars.length; i++) {
						sum += ((10-i) * parseInt(chars[i]));
					}
					valid = (sum % 11 == 0);
				} else if(value.length == 13) {
					var chars = value.split('');
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
				mCtrl.$setValidity('isbnInput', valid);
				return value;
			}
			mCtrl.$parsers.push(validation);
		}
	}
});
app.directive('appUploadFile', function() {
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
});