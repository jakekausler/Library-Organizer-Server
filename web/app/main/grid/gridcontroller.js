angular.module('libraryOrganizer')
    .controller('gridController', function($scope, $mdToast, $http, $mdSidenav, $mdDialog) {
        $scope.sort = $scope.getParameterByName("sort", "dewey");
        $scope.numberToGet = parseInt($scope.getParameterByName("numbertoget", "50"));
        $scope.page = parseInt($scope.getParameterByName("page", "1"));
        $scope.numberOfBooks = 0;
        $scope.fromdewey = $scope.getParameterByName("fromdewey", "");
        $scope.todewey = $scope.getParameterByName("todewey", "");
        $scope.fromlexile = $scope.getParameterByName("fromlexile", "");
        $scope.tolexile = $scope.getParameterByName("tolexile", "");
        $scope.withcodes = $scope.getParameterByName("withcodes", "").split(",");
        if ($scope.withcodes.length > 0 && $scope.withcodes[0] == "") {
            $scope.withcodes.splice(0, 1)
        }
        $scope.frominterestlevel = $scope.getParameterByName("frominterestlevel", "")
        $scope.tointerestlevel = $scope.getParameterByName("tointerestlevel", "")
        $scope.fromar = $scope.getParameterByName("fromar", "")
        $scope.toar = $scope.getParameterByName("toar", "")
        $scope.fromlearningaz = $scope.getParameterByName("fromlearningaz", "")
        $scope.tolearningaz = $scope.getParameterByName("tolearningaz", "")
        $scope.fromguidedreading = $scope.getParameterByName("fromguidedreading", "")
        $scope.toguidedreading = $scope.getParameterByName("toguidedreading", "")
        $scope.fromdra = $scope.getParameterByName("fromdra", "")
        $scope.todra = $scope.getParameterByName("todra", "")
        $scope.fromgrade = $scope.getParameterByName("fromgrade", "")
        $scope.tograde = $scope.getParameterByName("tograde", "")
        $scope.fromfountaspinnell = $scope.getParameterByName("fromfountaspinnell", "")
        $scope.tofountaspinnell = $scope.getParameterByName("tofountaspinnell", "")
        $scope.fromage = $scope.getParameterByName("fromage", "")
        $scope.toage = $scope.getParameterByName("toage", "")
        $scope.fromreadingrecovery = $scope.getParameterByName("fromreadingrecovery", "")
        $scope.toreadingrecovery = $scope.getParameterByName("toreadingrecovery", "")
        $scope.frompmreaders = $scope.getParameterByName("frompmreaders", "")
        $scope.topmreaders = $scope.getParameterByName("topmreaders", "")
        $scope.filter = $scope.getParameterByName("filter", "");
        $scope.read = $scope.getParameterByName("read", "both");
        $scope.reference = $scope.getParameterByName("reference", "both");
        $scope.owned = $scope.getParameterByName("owned", "yes");
        $scope.loaned = $scope.getParameterByName("loaned", "no");
        $scope.shipping = $scope.getParameterByName("shipping", "no");
        $scope.reading = $scope.getParameterByName("reading", "no");
        $scope.gridSelectedLibraries = $scope.getParameterByName("gridselectedlibraries", "").split(',')
        $scope.libraries = [];
        $scope.output = [];
        $scope.isFiltersOpen = false;
        $scope.forms = {};
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
            if (!$scope.forms.sortAndFilter || $scope.forms.sortAndFilter.$valid) {
                var loadingName = $scope.guid();
                $scope.addToLoading(loadingName)
                $scope.gridSelectedLibraries = $scope.stringLibraryIds();
                var params = {
                        sort: $scope.sort,
                        numbertoget: $scope.numberToGet,
                        page: $scope.page,
                        fromdewey: $scope.fromdewey.toUpperCase(),
                        todewey: $scope.todewey.toUpperCase(),
                        fromlexile: $scope.fromlexile,
                        tolexile: $scope.tolexile,
                        withcodes: $scope.withcodes.join(","),
                        frominterestlevel: $scope.frominterestlevel+"",
                        tointerestlevel: $scope.tointerestlevel+"",
                        fromar: $scope.fromar?$scope.fromar:"",
                        toar: $scope.toar?$scope.toar:"",
                        fromlearningaz: $scope.fromlearningaz+"",
                        tolearningaz: $scope.tolearningaz+"",
                        fromguidedreading: $scope.fromguidedreading+"",
                        toguidedreading: $scope.toguidedreading+"",
                        fromdra: $scope.fromdra+"",
                        todra: $scope.todra+"",
                        fromgrade: $scope.fromgrade+"",
                        tograde: $scope.tograde+"",
                        fromfountaspinnell: $scope.fromfountaspinnell+"",
                        tofountaspinnell: $scope.tofountaspinnell+"",
                        fromage: $scope.fromage+"",
                        toage: $scope.toage+"",
                        fromreadingrecovery: $scope.fromreadingrecovery+"",
                        toreadingrecovery: $scope.toreadingrecovery+"",
                        frompmreaders: $scope.frompmreaders+"",
                        topmreaders: $scope.topmreaders+"",
                        filter: $scope.filter,
                        read: $scope.read,
                        reference: $scope.reference,
                        owned: $scope.owned,
                        loaned: $scope.loaned,
                        shipping: $scope.shipping,
                        reading: $scope.reading,
                        gridselectedlibraries: $scope.gridSelectedLibraries
                    }
                for (key in params) {
                    if (key != "filter" && params[key] == "") {
                        delete params[key]
                    }
                }
                $scope.setParameters(params)
                $http({
                    url: '/books',
                    method: 'GET',
                    params: {
                        sortmethod: $scope.sort,
                        numbertoget: $scope.numberToGet,
                        page: $scope.page,
                        fromdewey: $scope.fromdewey.toUpperCase(),
                        todewey: $scope.todewey.toUpperCase(),
                        fromlexile: $scope.fromlexile,
                        tolexile: $scope.tolexile,
                        lexilecode: $scope.lexilecode,
                        frominterestlevel: $scope.frominterestlevel,
                        tointerestlevel: $scope.tointerestlevel,
                        fromar: $scope.fromar,
                        toar: $scope.toar,
                        fromlearningaz: $scope.fromlearningaz,
                        tolearningaz: $scope.tolearningaz,
                        fromguidedreading: $scope.fromguidedreading,
                        toguidedreading: $scope.toguidedreading,
                        fromdra: $scope.fromdra,
                        todra: $scope.todra,
                        fromgrade: $scope.fromgrade,
                        tograde: $scope.tograde,
                        fromfountaspinnell: $scope.fromfountaspinnell,
                        tofountaspinnell: $scope.tofountaspinnell,
                        fromage: $scope.fromage,
                        toage: $scope.toage,
                        fromreadingrecovery: $scope.fromreadingrecovery,
                        toreadingrecovery: $scope.toreadingrecovery,
                        frompmreaders: $scope.frompmreaders,
                        topmreaders: $scope.topmreaders,
                        text: $scope.filter,
                        isread: $scope.read,
                        isreference: $scope.reference,
                        isowned: $scope.owned,
                        isloaned: $scope.loaned,
                        isshipping: $scope.shipping,
                        isreading: $scope.reading,
                        isbn: "",
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
                    $scope.removeFromLoading(loadingName);
                }, function(response) {
                    $mdToast.showSimple("Failed to get books");
                    $scope.removeFromLoading(loadingName);
                });
            }
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
        $scope.toggleFilters = function() {
            $scope.isFiltersOpen = !$scope.isFiltersOpen
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
                "dewey": {
                    "String": "",
                    "Valid": false
                },
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
                "contributors": [],
                "library": {},
                "lexile": {
                    "Int64": 0,
                    "Valid": false
                },
                "lexilecode": "",
                "interestlevel": {
                    "Int64": "",
                    "Valid": false
                },
                "ar": {
                    "Float64": "",
                    "Valid": false
                },
                "learningaz": {
                    "Int64": "",
                    "Valid": false
                },
                "guidedreading": {
                    "Int64": "",
                    "Valid": false
                },
                "dra": {
                    "Int64": "",
                    "Valid": false
                },
                "grade": {
                    "Int64": "",
                    "Valid": false
                },
                "age": {
                    "Int64": "",
                    "Valid": false
                },
                "fountaspinnell": {
                    "Int64": "",
                    "Valid": false
                },
                "readingrecovery": {
                    "Int64": "",
                    "Valid": false
                },
                "pmreaders": {
                    "Int64": "",
                    "Valid": false
                },
                "notes": ""
            }
            $scope.showEditDialog(ev, book, $scope, 'gridadd');
        }
        $scope.updateLibraries = function() {
            var loadingName = $scope.guid();
            $scope.addToLoading(loadingName)
            $http({
                url: '/libraries',
                method: 'GET'
            }).then(function(response) {
                $scope.libraries = response.data;
                var data = [];
                var libStructure = {};
                for (l in $scope.libraries) {
                    if (!libStructure[$scope.libraries[l].owner]) {
                        libStructure[$scope.libraries[l].owner] = [];
                    }
                    libStructure[$scope.libraries[l].owner].push({
                        id: $scope.libraries[l].id,
                        name: $scope.libraries[l].name,
                        children: [],
                        selected: $.inArray($scope.libraries[l].id+"", $scope.gridSelectedLibraries)!=-1
                    });
                }
                for (k in libStructure) {
                    if ((!$scope.gridSelectedLibraries || !$scope.gridSelectedLibraries[0]) && !$scope.currentLibraryId && k == $scope.username) {
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
                $scope.removeFromLoading(loadingName);
            }, function(response) {
                $mdToast.showSimple("Failed to get list of libraries");
                $scope.removeFromLoading(loadingName);
            });;
        };
        $scope.chooseLibrary = function($ev) {
            $scope.showLibraryChooserDialog($ev, $scope, true)
        }
        $scope.$watch('output', function() {
            $scope.updateRecieved();
        })
        $scope.updateLibraries();
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
        $scope.scan = function(ev) {
            $mdDialog.show({
                controller: 'scanController',
                templateUrl: 'web/app/main/grid/scan/scandialog.html',
                parent: angular.element(document.body),
                targetEvt: ev,
                clickOutsideToClose: true,
                fullscreen: false,
                locals: {
                    vm: $scope
                }
            });
        }
        $scope.searchByISBN = function(isbn) {
            // 1. Query for isbn in current libraries
            // 2a. If not found, return false and display
            // 2b. If found, close the dialog and navigate to the book
        }
        $scope.lexilecodes = ["AD","NC","HL","IG","GN","BR","NP"]
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
