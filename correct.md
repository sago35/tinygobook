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
