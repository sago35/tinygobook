# 基礎から学ぶ TinyGoの組込み開発

高砂正哲が執筆した「基礎から学ぶ TinyGoの組込み開発」 (C&R研究所) のサポートサイトです。
質問や誤記などがある場合は本ページの Issue もしくは Twitter で受け付けています。

## 書籍情報

* C&R (紙) : https://www.c-r.com/book/detail/1477
* Amazon (紙、Kindle) : https://www.amazon.co.jp/dp/4863544006
* 本の森 (紙、PDF、EPUB) : https://book.mynavi.jp/manatee/c-r/books/detail/id=134168

![](./img/tinygobook.png)

各節まで含めた目次はこちら。

* [目次](./toc.md)

## Twitter

Twitter に投稿する時の hashtag は `#tinygo` と `#tinygobook` を使ってください。

* Twitter : [tinygo OR #tinygo OR @tinygolang OR #tinygobook](https://twitter.com/search?q=tinygo%20OR%20%23tinygo%20OR%20%40tinygolang%20OR%20%23tinygobook&src=typed_query&f=live)

## よくある質問と回答

### tinygo flash に失敗します

2 章 P.31 に従い (リセット x 2 で) ブートローダーに入れてから `tinygo flash` してみてください。
それでも改善しない場合は、 `tinygo build -o out.uf2` のようにして uf2 ファイルを作ってから手動で書き込みしてください。

## 正誤表

誤記等を見つけた場合は、 Issue もしくは Twitter で教えてください。


* [正誤表](./correct.md)

### 注意

ネットワーク部 (RTL8720DN) のファームウェアアップデート方法の修正があるため必ず確認してください。
具体的には、

```
$ git clone https://github.com/Seeed-Studio/ambd_flash_tool
```

ではなく

```
$ git clone https://github.com/Seeed-Studio/ambd_flash_tool --branch JP
```

を使うようにしてください。

### TinyGo と Go の組み合わせ

| TinyGo | Go | 備考 |
| --- | --- | --- |
| 0.28.1 | 1.19 - 1.20 | Wio Terminal の動作について一部問題あり(※1) |
| 0.27.0 | 1.19 - 1.20 | |
| 0.26.0 | 1.18 - 1.19 | 書籍執筆時 Version (脱稿直前のリリース) |
| 0.25.0 | 1.18 - 1.19 | 書籍執筆時 Version |

※1  
TinyGo 0.28.1 で Wio Terminal に搭載されている ATSAMD51 マイコンの Cache を有効化する変更が入り、 I2C などが一部動作不良となるケースがあります。
本件については [正誤表](./correct.md) に記載しています。


## Demos

デモアプリケーションはこちら。
Wio Terminal + TinyGo 0.26 で動作を確認しています。

### Wio Terminal Tracker

LIS3DH から得た情報を用いて、パソコン上の画像の傾きを制御する Demo です。

![](./img/tracker.png)

* [./wioterminal/tracker/](./wioterminal/tracker/)

```
$ tinygo flash --target wioterminal --size short ./wioterminal/tracker/
   code    data     bss |   flash     ram
  54496    1500    6260 |   55996    7760
```

注意) `-opt z` 以外でビルドすると、I2Cデータの受け取りに失敗します  

### Gopher福笑い

十字キーなどを使って目と口の位置を自由に動かすことができる Demo です。
面白い顔を作って Twitter に投稿してください。

![](./img/fukuwarai.png)

* [./wioterminal/fukuwarai](./wioterminal/fukuwarai/)

```shell
$ tinygo flash --target wioterminal --size short --opt 2 ./wioterminal/fukuwarai/
   code    data     bss |   flash     ram
 447876     356  180480 |  448232  180836
```

注意) 実行速度を高速化するため `-opt 2` でビルドすることを推奨します  

## Chapter 7 ネットワークに接続する

サポートサイトにて、 chap07 のコードを公開しています。
以下のようにして chap07 ディレクトリに移動してから書き込むことができます。
ssid や password が必要となるコードについては P.235 を参考に設定してください。
TinyGo 0.26 以降は `tinygo flash` 時に `--monitor` を指定することが出来ます。
多くの場合、 minicom や Tera Term を使わなくてもうまくやり取りできるはずです。

```
$ cd chap07/

$ tinygo flash --target wioterminal --size short --monitor ./update_test/
   code    data     bss |   flash     ram
  57452    1528    9004 |   58980   10532
Connected to COM5. Press Ctrl-C to exit.
RTL8270DN Firmware Version: 2.1.2
```


### 各種リンク

* https://tinygo.org/
    * https://tinygo.org/docs/reference/microcontrollers/wioterminal/
* Seeed
    * https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/ (TinyGo ではなく Arduino 情報)
    * https://wiki.seeedstudio.com/jp/Wio-Terminal-Getting-Started/ (TinyGo ではなく Arduino 情報)
* [sago35の日記 - Hatena Blog](https://sago35.hatenablog.com/)
    * [TinyGo 0.26 で遊べるマイコンボード一覧 - Hatena Blog](https://sago35.hatenablog.com/entry/2022/10/05/083000)
* [github.com/sago35/tinygo-workshop](https://github.com/sago35/tinygo-workshop)
    * Go Conference 2021 Autumn 内の Wio Terminal を使った TinyGo ハンズオン用の記事
    * ハンズオン時の動画: https://gocon.jp/2021autumn/sessions/workshop_a/
* [Wio Terminal で TinyGo プログラミングを始めよう - Qiita](https://qiita.com/sago35/items/92b22e8cbbf99d0cd3ef)
* Twitter : [tinygo OR #tinygo OR @tinygolang OR #tinygobook](https://twitter.com/search?q=tinygo%20OR%20%23tinygo%20OR%20%40tinygolang%20OR%20%23tinygobook&src=typed_query&f=live)

## 著者紹介

* 高砂正哲
    * GitHub: https://github.com/sago35
    * Twitter: [@sago35tk](https://twitter.com/sago35tk)

## about Gopher

The Gopher character is based on the Go mascot designed by [Renée French](https://reneefrench.blogspot.com/).
