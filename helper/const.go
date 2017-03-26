package helper

var VERSION = byte('\x05')
var AUTH_NO = byte('\x00')
var AUTH_USERNAME_PASSWORD = byte('\x02')
var NO_ACCEPTABLE_METHODS = byte('\xFF')

var CMD_CONNECT = byte('\x01')

var ATYP_IPV4 = byte('\x01')
var APTY_DOMAINNAME = byte('\x03')
var APTY_IPV6 = byte('\x04')

var REPLY_SUCCESS = byte('\x00')

var RSV = byte('\x00')
