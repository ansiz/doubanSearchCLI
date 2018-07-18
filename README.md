# doubanSearchCLI

Search Douban data in command line.

## Introduction

Search Douban data in command line. （being developed）

### Features

- Search and list books
- Sort by rate num

### TODO

- Add cache
- Search more

## Quick Start

### Prerequisite

- Golang 1.8+ installed

### Installation

```bash
git clone https://github.com/ansiz/doubanSearchCLI
cd doubanSearchCLI
make install
```

### Usage

```txt
NAME:
   Douban searcher - Search data from Douban

USAGE:
   dsearch [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     list     search items by specified keyword, output the list
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose      run in verbose mode
   --help, -h     show help
   --version, -v  print the version
```

#### Example

```txt
dsearch list book --verbose -k "计算机" -p 1

深入理解计算机系统（原书第2版）》-- （美）Randal E.Bryant / 龚奕利 / 机械工业出版社 / 2011		评分：9.7(2081人评价)

《深入理解计算机系统》-- （美）Randal E.Bryant, David R.O'Hallaron / 电子工业出版社 / 2006		评分：9.7(282人评价)

《具体数学》-- Ronald L.Graham / 张凡 / 人民邮电出版社 / 2013		评分：9.6(216人评价)

《深入理解计算机系统》-- Randal E.Bryant / 龚奕利 / 中国电力出版社 / 2004		评分：9.5(2560人评价)

《计算机程序的构造和解释》-- Harold Abelson / 裘宗燕 / 机械工业出版社 / 2004		评分：9.5(1935人评价)

《具体数学（英文版第2版）》-- [美] Ronald L. Graham / 机械工业出版社 / 2002		评分：9.5(802人评价)

《计算机程序设计艺术（第1卷）》-- [美] Donald E. Knuth / 清华大学出版社 / 2002		评分：9.4(445人评价)

《编码》-- [美] Charles Petzold / 左飞 / 电子工业出版社 / 2010		评分：9.2(2134人评价)

《计算机科学概论（第11版）》-- J. Glenn Brookshear / 刘艺 / 人民邮电出版社 / 2011		评分：9.2(254人评价)

《计算机系统要素》-- Noam Nisan / 周维 / 电子工业出版社 / 2007		评分：9.1(131人评价)

《黑客与画家》-- [美] Paul Graham / 阮一峰 / 人民邮电出版社 / 2013		评分：8.9(476人评价)

《计算机网络（第4版）》-- [美] James F. Kurose / 陈鸣 / 机械工业出版社 / 2009		评分：8.8(699人评价)

《计算机网络》-- Andrew S. Tanenbaum / 潘爱民 / 清华大学出版社 / 2004		评分：8.7(857人评价)

《黑客》-- Steven Levy / 赵俐 / 机械工业出版社华章公司 / 2011		评分：8.3(841人评价)

《奇思妙想》-- Dennis E. Shasha / 向怡宁 / 人民邮电出版社 / 2012		评分：8.1(304人评价)

《灵魂机器的时代》-- （美）雷·库兹韦尔 / 沈志彦 / 上海译文出版社 / 2006		评分：8.1(216人评价)

《奇点临近》-- Ray Kurzweil / 李庆诚 / 机械工业出版社 / 2011		评分：7.7(1198人评价)

《设计原本》-- Frederick P. Brooks, Jr. / InfoQ中文站 / 机械工业出版社 / 2011		评分：7.7(368人评价)

《C程序设计》-- 谭浩强 / 清华大学出版社 / 2005		评分：7.0(1429人评价)

《计算机》-- 于今昌 编 / 吉林出版集团有限责任公司 / 2007		评分：0.0（评价人数不足）
```

## Support and Bug Reports