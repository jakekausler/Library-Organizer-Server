angular.module('libraryOrganizer')
.controller('settingsController', function($scope, $mdDialog, $http) {
	$scope.settings = [{
		group: 'Group1',
		settings: [{
			setting: 'Setting1',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}]
	}, {
		group: 'Group2',
		settings: [{
			setting: 'Setting1',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}]
	}, {
		group: 'Group3',
		settings: [{
			setting: 'Setting1',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting2',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}, {
			setting: 'Setting3',
			value: 'Value',
			valuetype: '',
			possiblevalues: []
		}]
	}];
	$scope.cancel = function() {
		$mdDialog.cancel();
	};
});