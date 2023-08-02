# 正誤表

誤りや、誤りに見えるものがある場合、 Issue に記載お願いします。

## 重要な正誤情報

- Chapter 7 : ネットワークに接続する
  - P.229 - P.230 : JP リージョンのファームウェアを使う形に修正
    - 誤 : git clone https://github.com/Seeed-Studio/ambd_flash_tool
    - 正 : git clone https://github.com/Seeed-Studio/ambd_flash_tool --branch JP

## 些細な正誤情報

- Chapter 2 : 開発環境のセットアップ
  - P.32, P.34, P.36 : tinygo のオプションは `-` でも `--` でも問題ないが、 P.118 より前のページは `-` を使うべき
    - 誤の例 : tinygo flash --target wioterminal --size short examples/blinky1
    - 正の例 : tinygo flash -target wioterminal -size short examples/blinky1

## TinyGo の Version によるもの

- Chapter 5
  - (TinyGo 0.28.1 以降) P.185, P.189 のソースコード内で `"machine/usb/midi"` を import している箇所で import エラーになる
    - package path を `"machine/usb/adc/midi"` に書き換えてください

## その他雑多な情報

- TinyGo 0.28.1 での動作チェック結果
  - [Isssues #8](https://github.com/sago35/tinygobook/issues/8) に詳細情報あり
    - USB MIDI の package path は machine/usb/adc/midi に変更となりました
    - I2C が動作しないため、上記 Issue に記載の enableCache() 削除を行ってください
    - ディスプレイの PNG 表示が出来ないので、上記 Issue に記載の enableCache() 削除を行ってください
    - QSPI / SDcard が動かない場合は enableCache() 削除を行ってください
    - Chapter 8 の Wio Terminal Tracker が I2C に依存していて動かないので enbaleCache() 削除を行ってください
