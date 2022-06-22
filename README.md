[![sunset.exe](https://img.shields.io/github/v/release/apxxxxxxe/sunset-sunrise-saori?color=%23e3aa40&label=sunset.exe&logo=github)](https://github.com/apxxxxxxe/sunset-sunrise-saori/releases/latest/download/sunset.exe)
[![commits](https://img.shields.io/github/last-commit/apxxxxxxe/sunset-sunrise-saori?color=%23e3aa40&label=%E6%9C%80%E7%B5%82%E6%9B%B4%E6%96%B0&logo=github)](https://github.com/apxxxxxxe/sunset-sunrise-saori/commits/main)
[![commits](https://img.shields.io/tokei/lines/github/apxxxxxxe/sunset-sunrise-saori?color=%23e3aa40)](https://github.com/apxxxxxxe/sunset-sunrise-saori/commits/main)

# sunset-sunrise-saori
日没・日の出の時間を計算するSAORI-basic

### 既定値
|緯度|経度|標準時|
|--|--|--|
|35.000|135.000|+0900|

第２引数以降で指定がなければこの値になる。  
日本の特定地域の厳密な時間や海外の時間がほしい、などでない限りはこれでよいはず  
ちなみにこれは兵庫県明石市を基準とした値

### つかいかた
#### 第１引数
「YYYYMMDD」の形式で記述された日付(例:2022年7月1日→20220701)
#### 第２引数
緯度(実数、省略可)。
#### 第３引数
経度(実数、省略可)。
#### 第４引数
標準時(実数、省略可)。
#### 返り値
次の形式で日の出・日没の時刻がカンマ（,）区切りで返る  
例:日の出が4時49分29秒、日没が19時18分0秒→「04,49,29,19,18,00」

## 使用例(里々)
里々での使用例です。上記の意味がわからなくても以下の手順でゴーストに組み込むことができます。  

### 目標
この例では、里々における（現在時）（現在分）（現在秒）のように使える
- （日の出時）
- （日の出分）
- （日の出秒）
- （日没時）
- （日没分）
- （日没秒）

という変数（ゴースト起動毎に更新され、その日の日の出、日没の時刻を反映する）  
を用意し、使うことを目指します。

### 準備
- ゴーストのフォルダにsunset.exeを配置する。
  - この例では[ゴースト名]/ghost/master/saori/sunset.exeに配置する。
- satori_conf.txtに以下の記述を追加（すでに＠SAORIがあるならその最後尾にsunset以降を追記するだけでよい）
```
＠SAORI
sunset,saori/sunset.exe
```
これは本SAORIを関数として呼び出しできるように登録しています。  
ここに登録してあげることでsunsetという関数が何をするものなのかを里々に教えています

- 起動毎に読み込まれるイベント(OnSatoriLoad等)に以下の記述を追加
```
（nop、（split、（sunset、（sprintf、%04d%02d%02d、（現在年）、（現在月）、（現在日）））、,））φ
（set,日の出時,（sprintf,%d,（S0）））φ
（set,日の出分,（sprintf,%d,（S1）））φ
（set,日の出秒,（sprintf,%d,（S2）））φ
（set,日没時,（sprintf,%d,（S3）））φ
（set,日没分,（sprintf,%d,（S4）））φ
（set,日没秒,（sprintf,%d,（S5）））φ
```
この一連の処理をすることで、その日付時点での日の出、日没の時刻が各変数(日の出時、日の出分、日の出秒、日没時、日没分、日没秒)に代入されます。  
これで日の出/日没の時間を扱う準備ができました。
### 実際につかってみる
- （例）日の出の時刻になったら特別なトークをする
```
＊OnMinuteChange
＞日の出イベント【タブ】（日の出時）==（現在時）&&（日の出分）==（現在分）

＊日の出イベント
：もう日が昇っちゃたよ！はやくねなさーい！
```
この例では、毎分起こるイベント(OnMinuteChange)で日の出の時刻と現在時刻を比較し、  
それらが等しければ(=日の出の時刻になったとき)日の出イベントを起こしています。  
このように日の出/日没のタイミングでのトークが実現できます。

## 最新版ダウンロード
[![sunset.exe](https://img.shields.io/github/v/release/apxxxxxxe/sunset-sunrise-saori?color=%23e3aa40&label=sunset.exe&logo=github)](https://github.com/apxxxxxxe/sunset-sunrise-saori/releases/latest/download/sunset.exe)

