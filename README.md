# Competitive Programming Information Bot

[![CI](https://github.com/granddaifuku/Contests-Line-Bot/actions/workflows/test.yml/badge.svg)](https://github.com/granddaifuku/Contests-Line-Bot/actions)

## About
This is the [Line](https://line.me/ja/) bot that informs you about the [AtCoder](https://atcoder.jp/?lang=ja) and [Codeforces](https://codeforces.com/), [yukicoder](https://yukicoder.me/) contests.

## Information Contents
- AtCoder
  - Contest Name
  - Contest Start and End time
  - Rated Range
- Codeforces
  - Contest Name
  - Contest Start and End time
- yukicoder
  - Contest Name
  - Contest Start and End time

## How to use
1. Install the Line app.
2. Scan the QR code below.
3. Invite the bot into the group or the one-on-one talk room.
4. The bot recognizes the word "コンテスト" (now only Katakana is available). Then it sends you (or group) the pieces of information above.

## QR Code
<img width="213" alt="Screen Shot 2020-07-11 at 23 37 43" src="https://user-images.githubusercontent.com/49578068/87226596-cf319080-c3cf-11ea-950a-d0d25f76c805.png">

## 日本語版 README
[競プロの日程を通知する Line Bot を作った話](https://granddaifuku.hatenablog.com/entry/2020/01/22/210601)  

## For Developers
### Coding environment
- Golang
- Docker

### Testing
- `make test` to run all the tests.

### Stop Docker Things
- `make down` to stop containers and to remove networks, volumes, and images.

### Chores
- You can do miscellaneous work by `manage.sh`.
