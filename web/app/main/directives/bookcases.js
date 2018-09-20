angular.module('libraryOrganizer')
    .directive('bookcases', function() {
        var directive = {
            require: 'ngModel',
            template: '<div id="bookcases"></div>',
            link: function(vm, element, attrs) {
                vm.$watch(attrs.ngModel, function(cases) {
                    vm.cases = cases;
                    vm.container = document.getElementById('bookcases');
                    if (vm.cases) {
                        vm.drawShelf();
                    }
                    // document.getElementById('bookcase-canvas').removeEventListener('keypress', vm.zoomListener);
                    // document.getElementById('bookcase-canvas').addEventListener('keypress', vm.zoomListener);
                }, true);
                vm.zoomListener = function(e) {
                    // switch (e.key) {
                    // case 'w':
                    // 	vm.zoom *= 1.25;
                    // 	if (vm.zoom > 2.44140625) {
                    // 		vm.zoom = 2.44140625;
                    // 	}
                    // 	vm.drawShelf();
                    // 	break;
                    // case 's':
                    // 	vm.zoom /= 1.25;
                    // 	if (vm.zoom < 0.16777216) {
                    // 		vm.zoom = 0.16777216;
                    // 	}
                    // 	vm.drawShelf();
                    // 	break;
                    // }
                };
                vm.mouseclick = function(e, canvas, num) {
                    var rect = canvas.getBoundingClientRect();
                    var x = e.clientX - rect.left;
                    var y = e.clientY - rect.top;
                    var i = Math.floor(x / 100);
                    var j = Math.floor(y / 100);
                    if (vm.hashes[num]) {
                        for (var b in vm.hashes[num][i][j]) {
                            if (x < vm.hashes[num][i][j][b].x + vm.hashes[num][i][j][b].newwidth && x > vm.hashes[num][i][j][b].x && y < vm.hashes[num][i][j][b].y + vm.hashes[num][i][j][b].newheight && y > vm.hashes[num][i][j][b].y) {
                                vm.$parent.showBookDialog(e, vm.hashes[num][i][j][b], vm, 'shelves');
                            }
                        }
                    }
                };
                vm.cases = [];
                vm.hashes = [];
                vm.canvas = null;
                vm.zoom = 1;
                vm.doBoxesIntersect = function(a, b) {
                    return !(a.x > b.x + b.width || a.x + a.width < b.x || a.y > b.y + b.height || a.y + a.height < b.y);
                };
                vm.drawShelf = function() {
                    vm.container.innerHTML = "";
                    var margin = 50;
                    var x = margin;
                    var y = margin;
                    var width = margin;
                    var height = margin;
                    caseHeights = [];
                    for (var c in vm.cases) {
                        vm.cases[c].spacerheight *= vm.zoom;
                        vm.cases[c].width *= vm.zoom;
                        vm.cases[c].bookmargin *= vm.zoom;
                        var h = vm.cases[c].spacerheight;
                        for (var s in vm.cases[c].shelves) {
                            vm.cases[c].shelves[s].height *= vm.zoom;
                            h += vm.cases[c].spacerheight + vm.cases[c].shelves[s].height;
                        }
                        if (height < h) {
                            height = h;
                        }
                        caseHeights.push(h);
                        x += vm.cases[c].spacerheight + vm.cases[c].width + vm.cases[c].spacerheight + margin;
                        width = x;
                    }
                    vm.container.style.width = (width + margin) + "px";
                    vm.container.style.height = (height + margin) + "px";
                    vm.hashes = [];
                    vm.cases.forEach(function(_, c) {
                        var canvas = document.createElement('canvas');
                        vm.container.appendChild(canvas);
                        canvas.style.position = "inline-block";
                        canvas.width = vm.cases[c].width + vm.cases[c].spacerheight * 2;
                        canvas.height = height;
                        canvas.style.margin = (margin / 2) + "px";
                        canvas.addEventListener('click', function(e) {
                            vm.mouseclick(e, canvas, c);
                        });
                        hashes = [];
                        for (var i = 0; i < canvas.width / 100; i += 1) {
                            hashes.push([]);
                            for (var j = 0; j < canvas.height / 100; j += 1) {
                                hashes[i].push([]);
                                hashes[i][j] = [];
                            }
                        }
                        vm.hashes.push(hashes);
                        x = 0;
                        var ctx = canvas.getContext("2d");
                        ctx.font = (vm.zoom * 10) + "px Arial";
                        y = height - caseHeights[c];
                        var wood = document.getElementById('wood');
                        if (wood) {
                            ctx.drawImage(wood, x, y, vm.cases[c].width, caseHeights[c]);
                        }
                        for (var s in vm.cases[c].shelves) {
                            var ix = vm.cases[c].paddingleft + x + vm.cases[c].spacerheight;
                            ctx.fillRect(x, y, vm.cases[c].spacerheight, vm.cases[c].spacerheight + vm.cases[c].shelves[s].height);
                            ctx.fillRect(x + vm.cases[c].width, y, vm.cases[c].spacerheight, vm.cases[c].spacerheight + vm.cases[c].shelves[s].height);
                            ctx.fillRect(x + vm.cases[c].spacerheight, y, vm.cases[c].width, vm.cases[c].spacerheight);
                            for (var b in vm.cases[c].shelves[s].books) {
                                newWidth = vm.cases[c].shelves[s].books[b].width * vm.zoom;
                                newHeight = vm.cases[c].shelves[s].books[b].height * vm.zoom;
                                var spineColor = vm.cases[c].shelves[s].books[b].highlight === undefined ? vm.cases[c].shelves[s].books[b].spinecolor : (vm.cases[c].shelves[s].books[b].highlight ? "white" : "black");
                                var textColor;
                                if (vm.cases[c].shelves[s].books[b].highlight === undefined) {
                                    var converted = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(vm.cases[c].shelves[s].books[b].spinecolor);
                                    var rgb = converted ? [
                                        parseInt(converted[1], 16),
                                        parseInt(converted[2], 16),
                                        parseInt(converted[3], 16)
                                    ] : null;
                                    if (converted) {
                                        var o = Math.round(((parseInt(rgb[0], 10) * 299) + (parseInt(rgb[1], 10) * 587) + (parseInt(rgb[2], 10) * 114)) / 1000);
                                        textColor = (o > 125) ? 'black' : 'white';
                                    } else {
                                        textColor = 'white';
                                    }
                                } else {
                                    textColor = vm.cases[c].shelves[s].books[b].highlight ? "black" : "white";
                                }
                                ctx.fillStyle = 'black';
                                var bookwidth = newWidth <= 0 ? vm.cases[c].averagebookwidth * vm.zoom : newWidth;
                                var bookheight = newHeight <= 25 ? vm.cases[c].averagebookheight * vm.zoom : newHeight;
                                ctx.fillRect(ix, y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight, bookwidth, bookheight);
                                ctx.fillStyle = spineColor;
                                ctx.fillRect(ix + 1 * vm.zoom, y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight + 1 * vm.zoom, bookwidth - 2 * vm.zoom, bookheight - 2 * vm.zoom);
                                ctx.save();
                                ctx.translate(ix + bookwidth / 2, y + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight - 2 * vm.zoom);
                                ctx.rotate(-Math.PI / 2);
                                ctx.textAlign = "left";
                                ctx.textBaseline = "middle";
                                ctx.fillStyle = 'black';
                                var text = vm.cases[c].shelves[s].books[b].title;
                                if (ctx.measureText(text).width > bookheight - 4 * vm.zoom) {
                                    while (ctx.measureText(text + '...').width > bookheight - 4 * vm.zoom) {
                                        text = text.substring(0, text.length - 2);
                                    }
                                    text = text + '...';
                                }
                                ctx.fillStyle = textColor;
                                ctx.fillText(text, 0, 0);
                                ctx.restore();
                                vm.cases[c].shelves[s].books[b].x = ix;
                                vm.cases[c].shelves[s].books[b].y = y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight;
                                vm.cases[c].shelves[s].books[b].newwidth = bookwidth;
                                vm.cases[c].shelves[s].books[b].newheight = bookheight;
                                for (var i = Math.floor((ix - 1) / 100); i < Math.floor(((ix - 1) + (bookwidth + 2)) / 100) + 1; i += 1) {
                                    for (var j = Math.floor((y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight - 1) / 100); j < Math.floor(((y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight - 1) + (bookheight + 2)) / 100) + 1; j += 1) {
                                        if (i >= 0 && j >= 0 && vm.doBoxesIntersect({
                                                x: i * 100,
                                                y: j * 100,
                                                width: 100,
                                                height: 100
                                            }, {
                                                x: ix - 1,
                                                y: y - bookheight + vm.cases[c].shelves[s].height + vm.cases[c].spacerheight - 1,
                                                width: bookwidth + 2,
                                                height: bookheight + 2
                                            })) {
                                            vm.hashes[c][i][j].push(vm.cases[c].shelves[s].books[b]);
                                        }
                                    }
                                }
                                ix += bookwidth;
                            }
                            ctx.fillStyle = 'black';
                            y += vm.cases[c].spacerheight + vm.cases[c].shelves[s].height;
                        }
                        ctx.fillRect(x, y, vm.cases[c].width + vm.cases[c].spacerheight, vm.cases[c].spacerheight);
                        y = 0;
                        x += vm.cases[c].spacerheight + vm.cases[c].width + vm.cases[c].spacerheight;
                    });
                };
            }
        };
        return directive;
    });
