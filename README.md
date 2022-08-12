# tinygobook

高砂正哲が執筆した「基礎から学ぶ組込TinyGo」 (C&R研究所) のサポートサイトです。  
質問や誤記などがある場合は本ページの Issue もしくは Twitter で受け付けています。  

## Twitter

Twitter の hashtag は `#tinygo` と `#tinygobook` を使ってください。  

* Twitter : [tinygo OR #tinygo OR @tinygolang OR #tinygobook](https://twitter.com/search?q=tinygo%20OR%20%23tinygo%20OR%20%40tinygolang%20OR%20%23tinygobook&src=typed_query&f=live)

## よくある質問と回答

## 正誤表

## Demos

デモアプリケーションはこちら。  
Wio Terminal + TinyGo 0.25 で動作を確認しています。  

### Wio Terminal Tracker

* [./wioterminal/tracker/](./wioterminal/tracker/)

```
$ tinygo flash --target wioterminal --size short ./wioterminal/tracker/
   code    data     bss |   flash     ram
  54496    1500    6260 |   55996    7760
```

注意) `-opt z` 以外でビルドすると、I2Cデータの受け取りに失敗します  

### Gopher福笑い

* [./wioterminal/fukuwarai](./wioterminal/fukuwarai/)

```shell
$ tinygo flash --target wioterminal --size short --opt 2 ./wioterminal/fukuwarai/
   code    data     bss |   flash     ram
 447876     356  180480 |  448232  180836
```

注意) 実行速度を高速化するため `-opt 2` でビルドすることを推奨します  

## 著者紹介

* 高砂正哲
    * GitHub: https://github.com/sago35
    * Twitter: [@sago35tk](https://twitter.com/sago35tk)
