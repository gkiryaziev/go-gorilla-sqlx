
var crudApp = angular.module('crudApp', []);

crudApp.controller('crudCtrl', ['$scope', '$http', function($scope, $http) {

    // get all users
    $http.get('/api/v1/users').success(function(data, status, headers, config) {
    		$scope.users = data.data;
    }).error(function() {

    });

    // add new user
    $scope.userNew = function() {
        var newUser = {};
        newUser.first_name = $scope.first_name;
        newUser.last_name = $scope.last_name;
        newUser.middle_name = $scope.middle_name;
        newUser.dob = $scope.dob;
        newUser.address = $scope.address;
        newUser.phone = $scope.phone;
        newUser.login = $scope.login;
        newUser.password = $scope.password;

        $http({
            method : 'POST',
            url: "/api/v1/users",
            data: angular.toJson(newUser),
            headers : {'Content-Type': 'application/x-www-form-urlencoded'}
        })
        .success(function(data) {
            $scope.users.push({
                id: $scope.id,
                first_name: $scope.first_name,
                last_name: $scope.last_name,
                middle_name: $scope.middle_name,
                dob: $scope.dob,
                address: $scope.address,
                phone: $scope.phone,
                login: $scope.login,
                password: $scope.password,
            });
            $scope.id = '';
            $scope.first_name = '';
            $scope.last_name = '';
            $scope.middle_name = '';
            $scope.dob = '';
            $scope.address = '';
            $scope.phone = '';
            $scope.login = '';
            $scope.password = '';
        });
    };

    // save edited user
    $scope.userSave = function() {
        var index = getIndex($scope.id);
        $scope.users[index].first_name = $scope.first_name;
        $scope.users[index].last_name = $scope.last_name;
        $scope.users[index].middle_name = $scope.middle_name;
        $scope.users[index].dob = $scope.dob;
        $scope.users[index].address = $scope.address;
        $scope.users[index].phone = $scope.phone;
        $scope.users[index].login = $scope.login;
        $scope.users[index].password = $scope.password;

        $http({
            method : 'PUT',
            url: '/api/v1/users',
            data: angular.toJson($scope.users[index]),
            headers : {'Content-Type': 'application/x-www-form-urlencoded'}
        })
        .success(function(data) {

        });
    };

    // edit user by id
    $scope.userEdit = function(id) {
        var index = getIndex(id);
        var user = $scope.users[index];
        $scope.id = user.id;
        $scope.first_name = user.first_name;
        $scope.last_name = user.last_name;
        $scope.middle_name = user.middle_name;
        $scope.dob = user.dob;
        $scope.address = user.address;
        $scope.phone = user.phone;
        $scope.login = user.login;
        $scope.password = user.password;
    };

    // delete user by id
    $scope.userDelete = function(id) {
        var result = confirm('Are you sure?');
        if (result === true) {
            $http({
                method  : 'DELETE',
                url: "/api/v1/users/" + id,
            })
            .success(function(data) {
                var index = getIndex(id);
                $scope.users.splice(index, 1);
            });
        }
    };

    // get array index by id
    function getIndex(id) {
        for (var i = 0; $scope.users.length; i++) {
            if ($scope.users[i].id == id) {
                return i;
            }
        }
        return -1;
    }
}]);