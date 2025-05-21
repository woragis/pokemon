package config

var Admins = []string{"admin"}
var Moderators = append([]string{"moderators"}, Admins...)
var Writers = append([]string{"writer"}, Moderators...)

