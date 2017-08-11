angular.module('libraryOrganizer')
.controller('editorController', function($scope, $http, $mdDialog, $mdToast, book, $vm, viewType, username) {
	$scope.book = angular.copy(book);
	if (!$scope.book.tags) {
		$scope.book.tags = [];
	}
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
			$scope.book.dewey.String = value;
		});
		$vm.getSettingByName('Lexile', function(value) {
			$scope.book.lexile.Int64 = value;
		});
		$vm.getSettingByName('Interest Level', function(value) {
			$scope.book.interestlevel.Int64 = value;
		});
		$vm.getSettingByName('AR', function(value) {
			$scope.book.ar.Float64 = value;
		});
		$vm.getSettingByName('Learning AZ', function(value) {
			$scope.book.learningaz.Int64 = value;
		});
		$vm.getSettingByName('Guided Reading', function(value) {
			$scope.book.guidedreading.Int64 = value;
		});
		$vm.getSettingByName('DRA', function(value) {
			$scope.book.dra.Int64 = value;
		});
		$vm.getSettingByName('Grade', function(value) {
			$scope.book.grade.Int64 = value;6
		});
		$vm.getSettingByName('Fountas Spinnell', function(value) {
			$scope.book.fountaspinnell.Int64 = value;
		});
		$vm.getSettingByName('Age', function(value) {
			$scope.book.age.Int64 = value;
		});
		$vm.getSettingByName('Reading Recovery', function(value) {
			$scope.book.readingrecovery.Int64 = value;
		});
		$vm.getSettingByName('PM Readers', function(value) {
			$scope.book.pmreaders.Int64 = value;
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
	if (!$scope.book.lexile.Valid) {
		$scope.book.lexile.Int64 = "";
	}
	if (!$scope.book.interestlevel.Valid) {
		$scope.book.interestlevel.Int64 = "";
	}
	if (!$scope.book.ar.Valid) {
		$scope.book.ar.Float64 = "";
	}
	if (!$scope.book.learningaz.Valid) {
		$scope.book.learningaz.Int64 = "";
	}
	if (!$scope.book.guidedreading.Valid) {
		$scope.book.guidedreading.Int64 = "";
	}
	if (!$scope.book.dra.Valid) {
		$scope.book.dra.Int64 = "";
	}
	if (!$scope.book.grade.Valid) {
		$scope.book.grade.Int64 = "";
	}
	if (!$scope.book.fountaspinnell.Valid) {
		$scope.book.fountaspinnell.Int64 = "";
	}
	if (!$scope.book.age.Valid) {
		$scope.book.age.Int64 = "";
	}
	if (!$scope.book.readingrecovery.Valid) {
		$scope.book.readingrecovery.Int64 = "";
	}
	if (!$scope.book.pmreaders.Valid) {
		$scope.book.pmreaders.Int64 = "";
	}
	$scope.libraries = [];
	$scope.updateLibraries = function() {
        return $http({
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
	        }, function(response) {
	        	$mdToast.showSimple("Failed to retrieve list of libraries");
	        	$scope.libraries = [];
	        });
    };
    $scope.updateLibraries();
	$scope.updateDewey = function(d) {
        return $http({
				url: '/information/deweys/'+d,
				method: 'GET',
				params: {
					str: ''
				}
			}).then(function(response){
				return response.data;
			}, function(response) {
	        	$mdToast.showSimple("Failed to retrieve list of deweys");
	        	return [];
	        });
	};
	$scope.genre = "";
	$scope.$watch('book.dewey', function() {
        $scope.genre = $scope.updateDewey($scope.book.dewey);
    });
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
		book.lexile.Valid = isNaN(book.lexile.Int64)
		book.lexile.Int64 = isNaN(book.lexile.Int64)?0:parseInt(book.lexile.Int64);
		book.lexilecode = book.lexilecode?book.lexilecode:$scope.lexileSearchText;
		book.interestlevel.Valid = book.interestlevel.Int64 != ""
		book.interestlevel.Int64 = book.interestlevel.Int64==""?0:parseInt(book.interestlevel.Int64);
		book.ar.Valid = isNaN(book.ar.Float64)
		book.ar.Float64 = isNaN(book.ar.Float64)?0:parseFloat(book.ar.Float64);
		book.learningaz.Valid = book.learningaz.Int64 != ""
		book.learningaz.Int64 = book.learningaz.Int64==""?0:parseInt(book.learningaz.Int64);
		book.guidedreading.Valid = book.guidedreading.Int64 != ""
		book.guidedreading.Int64 = book.guidedreading.Int64==""?0:parseInt(book.guidedreading.Int64);
		book.dra.Valid = isNaN(book.dra.Int64)
		book.dra.Int64 = isNaN(book.dra.Int64)?0:parseInt(book.dra.Int64);
		book.grade.Valid = book.grade.Int64 != ""
		book.grade.Int64 = book.grade.Int64==""?0:parseInt(book.grade.Int64);
		book.fountaspinnell.Valid = book.fountaspinnell.Int64 != ""
		book.fountaspinnell.Int64 = book.fountaspinnell.Int64==""?0:parseInt(book.fountaspinnell.Int64);
		book.age.Valid = isNaN(book.age.Int64)
		book.age.Int64 = isNaN(book.age.Int64)?0:parseInt(book.age.Int64);
		book.readingrecovery.Valid = isNaN(book.readingrecovery.Int64)
		book.readingrecovery.Int64 = isNaN(book.readingrecovery.Int64)?0:parseInt(book.readingrecovery.Int64);
		book.pmreaders.Valid = isNaN(book.pmreaders.Int64)
		book.pmreaders.Int64 = isNaN(book.pmreaders.Int64)?0:parseInt(book.pmreaders.Int64);
		book.originallypublished = book.originallypublished+'-01-01';
		book.editionpublished = book.editionpublished+'-01-01';
		book.series = book.series?book.series:$scope.seriesSearchText;
		book.publisher.publisher = book.publisher.publisher?book.publisher.publisher:$scope.PublisherSearchText;
		book.publisher.city = book.publisher.city?book.publisher.city:$scope.CitySearchText;
		book.publisher.state = book.publisher.state?book.publisher.state:$scope.stateSearchText;
		book.publisher.country = book.publisher.country?book.publisher.country:$scope.countrySearchText;
		book.format = book.format?book.format:$scope.bindingSearchText;
		if (book.dewey.String == "0" || book.dewey.String == "00") {
			book.dewey.String = "000";
		}
		book.dewey.Valid = book.dewey.String != "" || book.dewey.String == "000"
		book.primarylanguage = book.primarylanguage?book.primarylanguage:$scope.primaryLanguageSearchText;
		book.secondarylanguage = book.secondarylanguage?book.secondarylanguage:$scope.secondaryLanguageSearchText;
		book.originallanguage = book.originallanguage?book.originallanguage:$scope.originalLanguageSearchText;
		var method = book.bookid ? 'PUT':'POST';
		console.log(book)
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
	$scope.query = function(f, str) {
		return $http({
				url: '/information/'+f.toLowerCase(),
				method: 'GET',
				params: {
					str: str
				}
			}).then(function(response){
				return response.data;
			}, function(response) {
	        	$mdToast.showSimple("Failed to retrieve list of " + f);
	        	return [];
	        });
	}
	$scope.log = function(item) {
		console.log(item)
	}
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
    $scope.lexilecodes = ["","AD","NC","HL","IG","GN","BR","NP"]
    $scope.interestlevels = [{
        name: "",
        value: ""
    }, {
        name: "LG",
        value: 0
    }, {
        name: "MG",
        value: 1
    }, {
        name: "MG+",
        value: 2
    }, {
        name: "UG",
        value: 3
    }]
    $scope.letters = [{
        name: "",
        value: ""
    }, {
        name: "A",
        value: 0
    }, {
        name: "B",
        value: 1
    }, {
        name: "C",
        value: 2
    }, {
        name: "D",
        value: 3
    }, {
        name: "E",
        value: 4
    }, {
        name: "F",
        value: 5
    }, {
        name: "G",
        value: 6
    }, {
        name: "H",
        value: 7
    }, {
        name: "I",
        value: 8
    }, {
        name: "J",
        value: 9
    }, {
        name: "K",
        value: 10
    }, {
        name: "L",
        value: 11
    }, {
        name: "M",
        value: 12
    }, {
        name: "N",
        value: 13
    }, {
        name: "O",
        value: 14
    }, {
        name: "P",
        value: 15
    }, {
        name: "Q",
        value: 16
    }, {
        name: "R",
        value: 17
    }, {
        name: "S",
        value: 18
    }, {
        name: "T",
        value: 19
    }, {
        name: "U",
        value: 20
    }, {
        name: "V",
        value: 21
    }, {
        name: "W",
        value: 22
    }, {
        name: "X",
        value: 23
    }, {
        name: "Y",
        value: 24
    }, {
        name: "Z",
        value: 25
    }]
    $scope.learningazlevels = [{
        name: "",
        value: ""
    }, {
        name: "aa",
        value: -1
    }, {
        name: "A",
        value: 0
    }, {
        name: "B",
        value: 1
    }, {
        name: "C",
        value: 2
    }, {
        name: "D",
        value: 3
    }, {
        name: "E",
        value: 4
    }, {
        name: "F",
        value: 5
    }, {
        name: "G",
        value: 6
    }, {
        name: "H",
        value: 7
    }, {
        name: "I",
        value: 8
    }, {
        name: "J",
        value: 9
    }, {
        name: "K",
        value: 10
    }, {
        name: "L",
        value: 11
    }, {
        name: "M",
        value: 12
    }, {
        name: "N",
        value: 13
    }, {
        name: "O",
        value: 14
    }, {
        name: "P",
        value: 15
    }, {
        name: "Q",
        value: 16
    }, {
        name: "R",
        value: 17
    }, {
        name: "S",
        value: 18
    }, {
        name: "T",
        value: 19
    }, {
        name: "U",
        value: 20
    }, {
        name: "V",
        value: 21
    }, {
        name: "W",
        value: 22
    }, {
        name: "X",
        value: 23
    }, {
        name: "Y",
        value: 24
    }, {
        name: "Z",
        value: 25
    }, {
        name: "Z1",
        value: 26
    }, {
        name: "Z2",
        value: 27
    }]
    $scope.grades = [{
        name: "",
        value: ""
    }, {
        name: "PK",
        value: 0
    }, {
        name: "K",
        value: 1
    }, {
        name: "1",
        value: 2
    }, {
        name: "2",
        value: 3
    }, {
        name: "3",
        value: 4
    }, {
        name: "4",
        value: 5
    }, {
        name: "5",
        value: 6
    }, {
        name: "6",
        value: 7
    }, {
        name: "7",
        value: 8
    }, {
        name: "8",
        value: 9
    }, {
        name: "9",
        value: 10
    }, {
        name: "10",
        value: 11
    }, {
        name: "11",
        value: 12
    }, {
        name: "12",
        value: 13
    }]
});
