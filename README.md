log
===

Leveled log implementation in Go for uniqush system

Checking uses of this log implementation
----------------------------------------

go tool vet -printfuncs debugf,infof,configf,warnf,errorf,alertf,fatalf .
