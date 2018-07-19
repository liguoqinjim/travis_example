# travis_example [![Build Status](https://travis-ci.org/liguoqinjim/travis_example.svg?branch=master)](https://travis-ci.org/liguoqinjim/travis_example)
travis-ci example

|文件夹|简介|说明|
|---|---|---|
|lab001|简答的单元测试||
|lab002|不做单元测试|要把这个里面的代码在测试覆盖率里面去除|

## 注意点
 - go1.10版本里面，`go test ./...`会自动去掉vendor文件夹，但是go1.6版本里面，`go test ./...`是会去包含vendor文件夹的。
 也就是vendor文件夹里面的也会去测试。
 - 点击travis-ci项目里面的status，可以得到status的链接。
 - 我们有的时候调用脚本的时候，会报错`Permission denied`，这个时候就在`before_install`下加上`chmod +x test.sh`就可以了。