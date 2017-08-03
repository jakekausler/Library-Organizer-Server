angular.module('libraryOrganizer')
    .controller('gridController', function($scope, $http, $mdSidenav, $mdDialog) {
        $scope.sort = $scope.getParameterByName("sort", "dewey");
        $scope.numberToGet = parseInt($scope.getParameterByName("numbertoget", "50"));
        $scope.page = parseInt($scope.getParameterByName("page", "1"));
        $scope.numberOfBooks = 0;
        $scope.fromdewey = $scope.getParameterByName("fromdewey", "0");
        $scope.todewey = $scope.getParameterByName("todewey", "FIC");
        $scope.fromlexile = $scope.getParameterByName("fromlexile", $scope.convertToLexile("0", ""));
        $scope.tolexile = $scope.getParameterByName("tolexile", $scope.convertToLexile("2000", ""));
        $scope.frominterestlevel = $scope.getParameterByName("frominterestlevel", "0")
        $scope.tointerestlevel = $scope.getParameterByName("tointerestlevel", "0")
        $scope.fromar = $scope.getParameterByName("fromar", "0")
        $scope.toar = $scope.getParameterByName("toar", "0")
        $scope.fromlearningaz = $scope.getParameterByName("fromlearningaz", "0")
        $scope.tolearningaz = $scope.getParameterByName("tolearningaz", "0")
        $scope.fromguidedreading = $scope.getParameterByName("fromguidedreading", "0")
        $scope.toguidedreading = $scope.getParameterByName("toguidedreading", "0")
        $scope.fromdra = $scope.getParameterByName("fromdra", "0")
        $scope.todra = $scope.getParameterByName("todra", "0")
        $scope.fromgrade = $scope.getParameterByName("fromgrade", "0")
        $scope.tograde = $scope.getParameterByName("tograde", "0")
        $scope.fromfountaspinnell = $scope.getParameterByName("fromfountaspinnell", "0")
        $scope.tofountaspinnell = $scope.getParameterByName("tofountaspinnell", "0")
        $scope.fromage = $scope.getParameterByName("fromage", "0")
        $scope.toage = $scope.getParameterByName("toage", "0")
        $scope.fromreadingrecovery = $scope.getParameterByName("fromreadingrecovery", "0")
        $scope.toreadingrecovery = $scope.getParameterByName("toreadingrecovery", "0")
        $scope.frompmreaders = $scope.getParameterByName("frompmreaders", "0")
        $scope.topmreaders = $scope.getParameterByName("topmreaders", "0")
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
                $scope.setParameters({
                        sort: $scope.sort,
                        numbertoget: $scope.numberToGet,
                        page: $scope.page,
                        fromdewey: $scope.fromdewey.toUpperCase(),
                        todewey: $scope.todewey.toUpperCase(),
                        fromlexile: $scope.fromlexile,
                        tolexile: $scope.tolexile,
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
                        filter: $scope.filter,
                        read: $scope.read,
                        reference: $scope.reference,
                        owned: $scope.owned,
                        loaned: $scope.loaned,
                        shipping: $scope.shipping,
                        reading: $scope.reading
                    })
                var fromlex = $scope.convertFromLexile($scope.fromlexile);
                var tolex = $scope.convertFromLexile($scope.tolexile);
                $http({
                    url: '/books',
                    method: 'GET',
                    params: {
                        sortmethod: $scope.sort,
                        numbertoget: $scope.numberToGet,
                        page: $scope.page,
                        fromdewey: $scope.fromdewey.toUpperCase(),
                        todewey: $scope.todewey.toUpperCase(),
                        fromlexile: fromlex[0],
                        fromlexilecode: fromlex[1],
                        tolexile: tolex[0],
                        tolexilecode: tolex[1],
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
                "contributors": [],
                "library": {},
                "lexile": 0,
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
            });
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
    });
