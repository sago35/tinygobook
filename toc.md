# 基礎から学ぶ TinyGoの組込み開発

* C&R (紙) : https://www.c-r.com/book/detail/1477
* Amazon (紙、Kindle) : https://www.amazon.co.jp/dp/4863544006
* 本の森 (紙、PDF、EPUB) : https://book.mynavi.jp/manatee/c-r/books/detail/id=134168

## 目次

* Chapter 1 TinyGoとは
    * TinyGoはGoの新しいコンパイラー
* Chapter 2 開発環境のセットアップ
    * ハードウェアの準備
    * TinyGoのセットアップ
    * 関連ツールのインストール
    * ビルドして動かしてみよう
    * Hello Wio Terminal
* Chapter 3 Goの基本
    * Hello World
    * 変数、定数、列挙型
    * 型
    * 制御フロー
    * 関数
    * メソッド
    * goroutineとchannelとselect
    * インターフェース
    * unsafe
    * cgo
* Chapter 4 TinyGo Internals
    * GOROOTとTINYGOROOT
    * ビルドタグ
    * TinyGoの標準package
    * TinyGoのビルドの流れ
    * TinyGoの実行
* Chapter 5 各ペリフェラルの使い方
    * ペリフェラルとは
    * GPIO
    * USB CDC
    * UART
    * ADC
    * DAC
    * PWM
    * I2C
    * SPI
    * USB HID
    * USB MIDI
    * TRNG
    * その他
* Chapter 6 ディスプレイに表示する
    * Display interface
    * tinydisplayでパソコン上の画面を作りこむ
    * 基本図形、フォント、画像の描画
    * ディスプレイの仕様
    * tinygo.org/x/drivers/ili9341の使い方
* Chapter 7 ネットワークに接続する
    * TinyGoのネットワーク機能
    * 本章でのソースコード構成
    * ネットワーク内のアクセスポイントに接続する
    * ネットワーク接続のセットアップ
    * TCP
    * UDP
    * NTP
    * MQTT
    * HTTP／HTTPS
* Chapter 8 アプリケーション作成
    * Wio Terminal Tracker
    * Gopher福笑い
    * 最後に
* 付録 デバッグ
    * マイコンをデバッグしよう
