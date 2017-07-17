angular.module('libraryOrganizer')
.controller('settingsController', function($scope, $mdDialog, $http, vm) {
    $scope.settings = [];
	$scope.pushSettings = function () {
		var settings = [];
		for (group in $scope.settings) {
			for (s in $scope.settings[group].settings) {
				setting = {
					group: $scope.settings[group].group,
					name: $scope.settings[group].settings[s].setting,
					value: $scope.settings[group].settings[s].value,
					valuetype: $scope.settings[group].settings[s].valuetype,
					possiblevalues: $scope.settings[group].settings[s].possiblevalues
				}
				settings.push(setting)
			}
		}
		$http({
			url: '/updatesettings',
			method: 'POST',
			data: JSON.stringify(settings)
		});
	}
    $scope.updateSettings = function() {
        $scope.settings = [];
        $http.get('/settings').then(function(response) {
            var settinggroups = {}
            response.data.forEach(function(v, i) {
                if (!settinggroups[v.group]) {
                    settinggroups[v.group] = [];
                }
                settinggroups[v.group].push({
                    setting: v.name,
                    value: v.value,
                    valuetype: v.valuetype,
                    possiblevalues: v.possiblevalues?v.possiblevalues:[]
                })
            })
            for (group in settinggroups) {
                var settings = [];
                for (setting in settinggroups[group]) {
                    settings.push({
                        setting: settinggroups[group][setting].setting,
                        value: settinggroups[group][setting].value,
                        valuetype: settinggroups[group][setting].valuetype,
                        possiblevalues: settinggroups[group][setting].possiblevalues
                    })
                }
                $scope.settings.push({
                    group: group,
                    settings: settings
                })
            }
        })
    }
    $scope.updateSettings();
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
});