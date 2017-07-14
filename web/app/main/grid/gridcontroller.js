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
        $scope.libraryids = [];
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
                    libraries: $scope.libraryids.stringLibraryIds()
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
        $scope.updateRecieved();
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
        $scope.stringLibraryIds = function() {
            return $scope.libraryids.join(',')
        }
    });
