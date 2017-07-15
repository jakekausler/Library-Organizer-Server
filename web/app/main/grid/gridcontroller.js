angular.module('libraryOrganizer')
    .controller('gridController', function($scope, $http, $mdSidenav) {
        $scope.sort = "dewey";
        $scope.numberToGet = 50;
        $scope.page = 1;
        $scope.numberOfBooks = 0;
        $scope.fromdewey = "0";
        $scope.todewey = 'FIC';
        $scope.filter = "";
        $scope.read = 'both';
        $scope.reference = 'both';
        $scope.owned = 'yes';
        $scope.loaned = 'no';
        $scope.shipping = 'no';
        $scope.reading = 'no';
        $scope.libraries = [];
        $scope.output = [];
        $scope.stringLibraryIds = function() {
            var retval = "";
            for (o in $scope.output) {
                if ($scope.output[o].children.length == 0 && $scope.output[o].selected) {
                    retval += $scope.output[o].id + ",";
                } else {
                    for (l in $scope.output[o].children) {
                        if ($scope.output[o].children[l].selected) {
                            retval += $scope.output[o].children[l].id + ",";
                        }
                    }
                }
            }
            if (retval.endsWith(",")) {
                retval = retval.substring(0,retval.length-1)
            }
            return retval;
        }
        $scope.updateRecieved = function() {
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
                    isreading: $scope.reading,
                    libraryids: $scope.stringLibraryIds()
                }
            }).then(function(response) {
                $scope.books = response.data.books;
                for (b in $scope.books) {
                    for (c in $scope.books[b].contributors) {
                        $scope.books[b].contributors[c].editing = false;
                    }
                }
                $scope.numberOfBooks = response.data.numbooks;
            });
        };
        $scope.previousPage = function() {
            $scope.page -= 1;
            $scope.updateRecieved();
        };
        $scope.nextPage = function() {
            $scope.page += 1;
            $scope.updateRecieved();
        };
        $scope.countPages = function() {
            if (isNaN($scope.numberToGet) || $scope.numberToGet <= 0) {
                return 0;
            }
            return Math.ceil($scope.numberOfBooks / $scope.numberToGet);
        };
        $scope.closeFiltersNav = function() {
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
            $scope.showEditorDialog(ev, book, $scope, 'gridadd');
        }
        $scope.updateLibraries = function() {
            $http.get('/libraries', {}).then(function(response) {
                $scope.libraries = response.data;
                var data = [];
                var libStructure = {}
                for (l in $scope.libraries) {
                    if (!libStructure[$scope.libraries[l].owner]) {
                        libStructure[$scope.libraries[l].owner] = [];
                    }
                    libStructure[$scope.libraries[l].owner].push({
                        id: $scope.libraries[l].id,
                        name: $scope.libraries[l].name,
                        children: [],
                        selected: false
                    });
                }
                for (k in libStructure) {
                    if (!$scope.currentLibraryId && k == $scope.username) {
                        $scope.currentLibraryId = libStructure[k][0].id
                        libStructure[k][0].selected = true;
                    }
                    data.push({
                        id: "owner/"+k,
                        name: k,
                        children: libStructure[k],
                        selected: false
                    })
                }
                $scope.libraries = angular.copy(data);
                $scope.output = angular.copy($scope.libraries);
                $scope.updateRecieved();
            });
        };
        $scope.updateLibraries();
        $scope.chooseLibrary = function($ev) {
            $scope.showLibraryChooserDialog($ev, $scope, true)
        }
        $scope.$watch('output', function() {
            $scope.updateRecieved();
        })
        $scope.getBookIcon = function(book) {
            if (book.library.owner==$scope.username) {
                return "web/res/edit.svg";
            } else if ((book.library.permissions&4)==4) {
                return "web/res/edit.svg";
            } else if ((book.library.permissions&2)==2) {
                return "web/res/checkout.svg";
            } else if ((book.library.permissions&1)==1) {
                return "web/res/view.svg";
            }
        }
    });
