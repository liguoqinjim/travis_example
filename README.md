# travis_example [![Build Status](https://travis-ci.org/liguoqinjim/travis_example.svg?branch=master)](https://travis-ci.org/liguoqinjim/travis_example) [![codecov](https://codecov.io/gh/liguoqinjim/travis_example/branch/master/graph/badge.svg)](https://codecov.io/gh/liguoqinjim/travis_example)
travis-ci example

|文件夹|简介|说明|
|---|---|---|
|lab001|简答的单元测试||
|lab002|不做单元测试|要把这个里面的代码在代码覆盖率里面去除|

## 注意点
 - go1.10版本里面，`go test ./...`会自动去掉vendor文件夹，但是go1.6版本里面，`go test ./...`是会去包含vendor文件夹的。
 也就是vendor文件夹里面的也会去测试，所以在测试老版本的go的时候，要用脚本把vendor文件夹去掉
 - 点击travis-ci项目里面的status，可以得到status的链接。
 - 有的时候调用脚本的时候，会报错`Permission denied`，在`before_install`下加上`chmod +x test.sh`就可以了。
 - 在codecov的setting里面，可以得到markdown用的badge
 - 可以在test.sh脚本里面去掉不想要加入代码覆盖率的文件或者文件夹
 - 每个包里面都要有一个`xxx_test.go`文件，要是没有这个文件的话，go test会报`no test files`，那么代码覆盖率里面也不会去统计这个包了。