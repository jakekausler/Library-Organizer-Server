angular.module('libraryOrganizer')
    .controller('scanController', function($scope, $mdDialog, $http, vm) {
        $scope.GOOGLE_BOOKS_API_BASE = 'https://www.googleapis.com';
        $scope.GOOGLE_BOOKS_API_BOOK = '/books/v1/volumes';
        $scope.WORLDCAT_API_BASE = 'http://xisbn.worldcat.org';
        $scope.WORLDCAT_API_BOOK = '/webservices/xid/isbn';
        $scope.vm = vm;
        $scope.isbn = '';
        $scope.results = {
            online: [],
            inlibrary: []
        };
        $scope.cancel = function() {
            $mdDialog.cancel()
        };
        $scope.scan = function() {
            if ($scope.isbn) {
                var loadingName = $scope.vm.guid();
                $scope.vm.addToLoading(loadingName)
                var parts = $scope.isbn.split(' ');
                if (parts.length == 1) {
                    $scope.isbn = parts[0];
                } else if (parts.length == 2) {
                    $scope.isbn = parts[1];
                } else if (parts.length == 3) {
                    $scope.isbn = parts[1];
                }
                $scope.searchByISBN($scope.isbn);
                $http({
                    url: $scope.GOOGLE_BOOKS_API_BASE + $scope.GOOGLE_BOOKS_API_BOOK + '?q=' + $scope.isbn
                }).then(function(response) {
                    response.data.items.forEach(function(v, i) {
                        $scope.results.online.push(v)
                    })
                    $scope.vm.removeFromLoading(loadingName);
                }, function(response) {
                    $mdToast.showSimple("Failed to retrieve books");
                    $vm.removeFromLoading(loadingName);
                });
            }
        };
        $scope.selectOnline = function(ev, result) {
            var book = {
                "bookid": "",
                "title": result.volumeInfo.title ? result.volumeInfo.title : '',
                "subtitle": result.volumeInfo.subtitle ? result.volumeInfo.subtitle : '',
                "originallypublished": result.volumeInfo.publishedDate ? result.volumeInfo.publishedDate.substring(0, 4) : '',
                "publisher": {
                    "id": "",
                    "publisher": result.volumeInfo.publisher ? result.volumeInfo.publisher : '',
                    "city": "",
                    "state": "",
                    "country": "",
                    "parentcompany": ""
                },
                "isread": false,
                "isreference": false,
                "isowned": true,
                "isbn": $scope.isbn,
                "dewey": {
                    "String": "",
                    "Valid": false
                },
                "pages": result.volumeInfo.pageCount ? result.volumeInfo.pageCount : 0,
                "width": result.volumeInfo.dimensions && result.volumeInfo.dimensions.width ? result.volumeInfo.dimensions.width : 0,
                "height": result.volumeInfo.dimensions && result.volumeInfo.dimensions.depth ? result.volumeInfo.dimensions.depth : 0,
                "depth": result.volumeInfo.dimensions && result.volumeInfo.dimensions.thickness ? result.volumeInfo.dimensions.thickness : 0,
                "weight": 0,
                "primarylanguage": result.volumeInfo.language && $scope.languages[result.volumeInfo.language] ? $scope.languages[result.volumeInfo.language] : "",
                "secondarylanguage": "",
                "originallanguage": result.volumeInfo.language && $scope.languages[result.volumeInfo.language] ? $scope.languages[result.volumeInfo.language] : "",
                "series": "",
                "volume": 0,
                "format": "",
                "edition": 0,
                "isreading": false,
                "isshipping": false,
                "imageurl": result.volumeInfo.imageLinks && result.volumeInfo.imageLinks.thumbnail ? result.volumeInfo.imageLinks.thumbnail : "",
                "spinecolor": "",
                "cheapestnew": 0,
                "cheapestused": 0,
                "editionpublished": result.volumeInfo.publishedDate ? result.volumeInfo.publishedDate.substring(0, 4) : '',
                "contributors": $scope.getAuthors(result.volumeInfo.authors),
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
            $scope.vm.showEditDialog(ev, book, $scope.vm, 'scanadd');
        };
        $scope.selectInLibrary = function(ev, result) {
            $scope.vm.showBookDialog(ev, result, $scope.vm, 'grid')
        }
        $scope.getAuthors = function(authors) {
            var contributors = [];
            if (authors) {
                authors.forEach(function(v, i) {
                    var first, middles, last = "";
                    v = v.replace('.', '');
                    v = v.split(' ');
                    if (v.length == 1) {
                        last = v[0];
                    } else if (v.length == 2) {
                        first = v[0];
                        last = v[1];
                    } else if (v.length > 2) {
                        first = v[0];
                        last = v[v.length - 1];
                        middles = v.slice(1, -1).join(';');
                    }
                    contributors.push({
                        role: 'Author',
                        name: {
                            first: first,
                            middles: middles,
                            last: last
                        }
                    });
                });
            }
            return contributors;
        };
        $scope.searchByISBN = function(isbn) {
            var loadingName = $scope.vm.guid();
            $scope.vm.addToLoading(loadingName)
            $http({
                url: '/books',
                method: 'GET',
                params: {
                    sortmethod: $scope.vm.sort,
                    numbertoget: -1,
                    page: 1,
                    fromdewey: $scope.vm.fromdewey.toUpperCase(),
                    todewey: $scope.vm.todewey.toUpperCase(),
                    fromlexile: $scope.vm.fromlexile,
                    tolexile: $scope.vm.tolexile,
                    lexilecode: $scope.vm.lexilecode,
                    text: $scope.vm.filter,
                    isread: $scope.vm.read,
                    isreference: $scope.vm.reference,
                    isowned: $scope.vm.owned,
                    isloaned: $scope.vm.loaned,
                    isshipping: $scope.vm.shipping,
                    isreading: $scope.vm.reading,
                    isbn: isbn,
                    libraryids: $scope.vm.stringLibraryIds()
                }
            }).then(function(response) {
                $scope.results.inlibrary = response.data.books;
                $scope.vm.removeFromLoading(loadingName);
            }, function(response) {
                $mdToast.showSimple("Failed to search library.");
                $vm.removeFromLoading(loadingName);
            });
        };
        $scope.languages = {
            'ab': 'Abkhazian',
            'aa': 'Afar',
            'af': 'Afrikaans',
            'ak': 'Akan',
            'sq': 'Albanian',
            'am': 'Amharic',
            'ar': 'Arabic',
            'an': 'Aragonese',
            'hy': 'Armenian',
            'as': 'Assamese',
            'av': 'Avaric',
            'ae': 'Avestan',
            'ay': 'Aymara',
            'az': 'Azerbaijani',
            'bm': 'Bambara',
            'ba': 'Bashkir',
            'eu': 'Basque',
            'be': 'Belarusian',
            'bn': 'Bengali',
            'bh': 'Bihari languages',
            'bi': 'Bislama',
            'bs': 'Bosnian',
            'br': 'Breton',
            'bg': 'Bulgarian',
            'my': 'Burmese',
            'ca': 'Catalan, Valencian',
            'ch': 'Chamorro',
            'ce': 'Chechen',
            'ny': 'Chichewa, Chewa, Nyanja',
            'zh': 'Chinese',
            'cv': 'Chuvash',
            'kw': 'Cornish',
            'co': 'Corsican',
            'cr': 'Cree',
            'hr': 'Croatian',
            'cs': 'Czech',
            'da': 'Danish',
            'dv': 'Divehi, Dhivehi, Maldivian',
            'nl': 'Dutch, Flemish',
            'dz': 'Dzongkha',
            'en': 'English',
            'eo': 'Esperanto',
            'et': 'Estonian',
            'ee': 'Ewe',
            'fo': 'Faroese',
            'fj': 'Fijian',
            'fi': 'Finnish',
            'fr': 'French',
            'ff': 'Fulah',
            'gl': 'Galician',
            'ka': 'Georgian',
            'de': 'German',
            'el': 'Greek (modern)',
            'gn': 'Guaraní',
            'gu': 'Gujarati',
            'ht': 'Haitian, Haitian Creole',
            'ha': 'Hausa',
            'he': 'Hebrew (modern)',
            'hz': 'Herero',
            'hi': 'Hindi',
            'ho': 'Hiri Motu',
            'hu': 'Hungarian',
            'ia': 'Interlingua',
            'id': 'Indonesian',
            'ie': 'Interlingue',
            'ga': 'Irish',
            'ig': 'Igbo',
            'ik': 'Inupiaq',
            'io': 'Ido',
            'is': 'Icelandic',
            'it': 'Italian',
            'iu': 'Inuktitut',
            'ja': 'Japanese',
            'jv': 'Javanese',
            'kl': 'Kalaallisut, Greenlandic',
            'kn': 'Kannada',
            'kr': 'Kanuri',
            'ks': 'Kashmiri',
            'kk': 'Kazakh',
            'km': 'Central Khmer',
            'ki': 'Kikuyu, Gikuyu',
            'rw': 'Kinyarwanda',
            'ky': 'Kirghiz, Kyrgyz',
            'kv': 'Komi',
            'kg': 'Kongo',
            'ko': 'Korean',
            'ku': 'Kurdish',
            'kj': 'Kuanyama, Kwanyama',
            'la': 'Latin',
            'lb': 'Luxembourgish, Letzeburgesch',
            'lg': 'Ganda',
            'li': 'Limburgan, Limburger, Limburgish',
            'ln': 'Lingala',
            'lo': 'Lao',
            'lt': 'Lithuanian',
            'lu': 'Luba-Katanga',
            'lv': 'Latvian',
            'gv': 'Manx',
            'mk': 'Macedonian',
            'mg': 'Malagasy',
            'ms': 'Malay',
            'ml': 'Malayalam',
            'mt': 'Maltese',
            'mi': 'Maori',
            'mr': 'Marathi',
            'mh': 'Marshallese',
            'mn': 'Mongolian',
            'na': 'Nauru',
            'nv': 'Navajo, Navaho',
            'nd': 'North Ndebele',
            'ne': 'Nepali',
            'ng': 'Ndonga',
            'nb': 'Norwegian Bokmål',
            'nn': 'Norwegian Nynorsk',
            'no': 'Norwegian',
            'ii': 'Sichuan Yi, Nuosu',
            'nr': 'South Ndebele',
            'oc': 'Occitan',
            'oj': 'Ojibwa',
            'cu': 'Church Slavic, Church Slavonic, Old Church Slavonic, Old Slavonic, Old Bulgarian',
            'om': 'Oromo',
            'or': 'Oriya',
            'os': 'Ossetian, Ossetic',
            'pa': 'Panjabi, Punjabi',
            'pi': 'Pali',
            'fa': 'Persian',
            'pl': 'Polish',
            'ps': 'Pashto, Pushto',
            'pt': 'Portuguese',
            'qu': 'Quechua',
            'rm': 'Romansh',
            'rn': 'Rundi',
            'ro': 'Romanian, Moldavian, Moldovan',
            'ru': 'Russian',
            'sa': 'Sanskrit',
            'sc': 'Sardinian',
            'sd': 'Sindhi',
            'se': 'Northern Sami',
            'sm': 'Samoan',
            'sg': 'Sango',
            'sr': 'Serbian',
            'gd': 'Gaelic, Scottish Gaelic',
            'sn': 'Shona',
            'si': 'Sinhala, Sinhalese',
            'sk': 'Slovak',
            'sl': 'Slovenian',
            'so': 'Somali',
            'st': 'Southern Sotho',
            'es': 'Spanish, Castilian',
            'su': 'Sundanese',
            'sw': 'Swahili',
            'ss': 'Swati',
            'sv': 'Swedish',
            'ta': 'Tamil',
            'te': 'Telugu',
            'tg': 'Tajik',
            'th': 'Thai',
            'ti': 'Tigrinya',
            'bo': 'Tibetan',
            'tk': 'Turkmen',
            'tl': 'Tagalog',
            'tn': 'Tswana',
            'to': 'Tonga (Tonga Islands)',
            'tr': 'Turkish',
            'ts': 'Tsonga',
            'tt': 'Tatar',
            'tw': 'Twi',
            'ty': 'Tahitian',
            'ug': 'Uighur, Uyghur',
            'uk': 'Ukrainian',
            'ur': 'Urdu',
            'uz': 'Uzbek',
            've': 'Venda',
            'vi': 'Vietnamese',
            'vo': 'Volapük',
            'wa': 'Walloon',
            'cy': 'Welsh',
            'wo': 'Wolof',
            'fy': 'Western Frisian',
            'xh': 'Xhosa',
            'yi': 'Yiddish',
            'yo': 'Yoruba',
            'za': 'Zhuang, Chuang',
            'zu': 'Zulu'
        }
    });
