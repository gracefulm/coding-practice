# 347. Top K Frequent Elements

<https://leetcode.com/problems/top-k-frequent-elements/>

## Step1

- ナイーブに実装する
  - mapを使って集計
  - その数値(num)と出現回数(frequency)の組み合わせを出現回数で降順ソート
  - ソートした要素の先頭からk番目までのnumを返す
- ソートを入れると計4回ループしており、ループ回数を減らす方法がある気がしているが一旦通ったのでOK

## Step2

- 2種類の方法で実装
  - `step2/sort`: Step1の方針(map集計 → ソート → 先頭k個)のままAIレビューを受けてブラッシュアップ
  - `step2/heap`: サイズkの最小ヒープを使った実装

### sort版(`step2/sort`)

- Step1からの変更点
  - バグ修正: 先頭k個に切り詰める条件
    - Step1: `if len(nums) > k { sortedNFs = sortedNFs[:k] }`
    - Step2: `nfs = nfs[:min(len(nfs), k)]`
    - Step1は`nums`の長さで判定しており、異なる値の数がk未満だと結果にゼロ値が混ざっていた
      - 例: `nums=[1,1,1,2], k=3` → `[1 2 0]`
      - スライス式`s[low:high]`の上限は`len(s)`ではなく`cap(s)`なので、範囲外エラーにならずcapの領域(ゼロ値)が見えてしまう
    - LeetCodeの制約(kは異なる値の数以下)内では発生しない
  - 入力の前提: 防御的に書く方を選択
    - `len(nums) == 0 || k <= 0`の場合は早期リターンする
      - Step1は`len(nums) == 0`だけを見ていたが、実際にpanicする可能性があるのは`k < 0`の方(`make`のcapに負の値を渡すとpanic)
    - kが異なる値の数より大きい場合は、上の`min`で全件を返す
  - カウント処理をシンプルにした
    - Step1: `c, ok := frequency[n]`で存在チェックしてから代入
    - Step2: `frequency[n]++`
    - mapは存在しないキーを読むとゼロ値を返すので存在チェックは不要
    - 出現回数に0が入ることはないので、「キーがない」と「値が0」を区別する必要もない
  - スライスのcapの見積もりを見直した
    - Step1: `make([]numFrequency, 0, k)`
    - Step2: `make([]numFrequency, 0, len(frequency))`
    - 入る要素数は異なる値の数なので`len(frequency)`がちょうど良い
      - `len(map)`はキーの数を返す
  - ソートの比較関数を`cmp.Compare`に変えた
    - Step1の`j.frequency - i.frequency`は値によってはオーバーフローし得る
    - 意図も読み取りやすい
  - 変数名を`sortedNFs`から`nfs`に変えた
    - appendしている時点ではまだソートされていないため
- 計算量(mは異なる値の数)
  - 時間: O(n + m log m)
    - ソートのO(m log m)が支配的
  - 空間: O(m)

### heap版(`step2/heap`)

- Step1で気になっていた「ループ回数を減らせないか」への対応
  - ループ回数そのものより、支配的なソートのO(m log m)を減らす方が効く
- 方針
  - mapで集計するところはsort版と同じ
  - mapの要素を1つずつ最小ヒープにPushし、k個を超えたら根(出現回数が一番少ないもの)をPopする
  - 最後にヒープに残ったk個が答え
- なぜ最大ヒープではなく最小ヒープなのか
  - 上位k個を残したいので、「次に追い出す候補 = 今の上位k個の中で一番少ないもの」をすぐ取り出せる必要がある
  - 最大ヒープに全件入れてk回Popする方法もあるが、ヒープのサイズがmになるので計算量はO(m + k log m)になり、ソートとの差は小さい
- 計算量
  - 時間: O(n + m log k)
    - mapの構築: O(n)
    - ヒープへのPush/Pop: 1回O(log k)をm回
  - 空間: O(m)
    - map: O(m)、ヒープ: O(k)
- `container/heap`の使い方
  - `heap.Interface`(`Len`/`Less`/`Swap`/`Push`/`Pop`)を満たす型を自分で定義する
    - `Less`を`<`にすると最小ヒープになる
    - `Push`/`Pop`はスライスの長さを変えるのでレシーバはポインタにする
  - 要素の出し入れは`mh.Push(...)`ではなく`heap.Push(&mh, ...)`/`heap.Pop(&mh)`で行う
    - 自分で定義した`Push`/`Pop`は末尾の追加・削除をするだけで、ヒープの並びを保つのは`container/heap`側の関数
  - `heap.Init`は既に要素が入っているスライスをヒープの形に並べ替えるためのもので、空のヒープには不要
  - ヒープは最大k+1個になるので`make(minHeap, 0, k+1)`で事前に確保した
- 結果の取り出し
  - ヒープが空になるまでPopしてappendしている
  - 最小ヒープなので結果は出現回数の少ない順になる(問題文では順序は問われていない)
  - kが異なる値の数より大きい場合もヒープの要素数分だけPopするのでpanicしない
    - 最初はk回Popしていて、ヒープが空の状態で`heap.Pop`を呼んでpanicしていた
- 実装中に踏んだバグ
  - `make([]int, n)`で作ったスライスに`append`して、先頭にゼロ値が混ざった
    - `make([]int, n)`はlenもnになる(中身はゼロ値n個)ので、appendするとその後ろに追加される
    - appendするなら`make([]int, 0, n)`、インデックスで代入するなら`make([]int, n)`
    - Step1の`[1 2 0]`と同じくlenとcapの取り違え

### その他メモ

- 毎回調べている気がする...スライス式は半開区間`[low, high)`
  - `s[:n]`は「先頭からn個」なので、全件は`s[:len(s)]`(`len(s)-1`ではない)
  - インデックスアクセス`s[i]`の上限は`len(s)-1`なので混同しやすい
- lenとcapの取り違えによるバグを2回出しているので要注意

### 残課題

- 今後の復習: バケットソート(O(n))での実装
  - 時間O(n)、空間O(n)でsort版・heap版より計算量が小さい
  - 方針のヒント
    - 1つの値の出現回数は最大でもnなので、出現回数をインデックスにした配列(長さn+1)に振り分けられる
    - `buckets[freq]`に出現回数がfreqの値を入れておき、出現回数の大きい方から見てk個集まったら終わり
  - 確認したいこと
    - 計算量がO(n)になる理由と、sort版・heap版との比較
    - 同じ出現回数の値が複数ある場合の扱い
    - kが異なる値の数より大きい場合にpanicしないか
- sort版とheap版でパッケージ名が揃っていない(`package step2`/`package heap`)
  - 型名・変数名はレビュー対応で揃えた
- テストコードを用意して、両方の実装に同じケースを流せるようにする

## Step3

- min heapを使った実装で練習

## レビュー対応

- PR: <https://github.com/gracefulm/coding-practice/pull/9>

### 変数名・型名の見直し

- 指摘1: mapの変数名は`(キー)To(値)`にすると読みやすい → `numToFrequency`
  - `frequency`だけだと、mapなのか、ある1つの値の出現回数(int)なのかが名前から分からない
  - キーと値の両方が名前に入っていると、`numToFrequency[num]++`のように使う側のコードも読みやすい
  - 調べたこと: map の命名は`keyToValue`以外にもある
    - `valueByKey`(例: `frequencyByNum`)もよく使われる
    - 用途が明らかなら`counts`や`seen`のような複数形・役割名も使われる
    - Goの公式スタイルにmapの命名規則はなく、「スコープが狭いほど短く、広いほど説明的に」が基本
    - どれを選ぶにしても、ファイル間・Step間で揃えるのが大事
- 指摘2: 単語の文字を削った省略形は読み手の負担になるので、原則フルスペルにする
  - 参考: <https://google.github.io/styleguide/cppguide.html#General_Naming_Rules>
    - > Do not abbreviate by deleting letters within a word.
  - API、JSONのように広く知られた略語や、`num`/`max`/`min`のようによく見るprefixは使ってよい
  - 最終的にはチームで一般的な書き方に合わせる
- 対応内容(Step1〜3すべて)

  | 変更前 | 変更後 | 対象 |
  |---|---|---|
  | `frequency`/`numsFreq`/`numFreqMap` (map) | `numToFrequency` | Step1, Step2(sort/heap), Step3 |
  | `sortedNFs`/`nfs` | `numFrequencies` | Step1, Step2(sort) |
  | `numFreq`型、`freq`フィールド | `numFrequency`型、`frequency`フィールド | Step2(heap), Step3 |
  | `mh` (ヒープの変数) | `topKHeap` | Step2(heap), Step3 |
  | `topk` | `topK` | Step2(heap), Step3 |
  | `n`, `f` (mapのrangeの変数) | `num`, `frequency` | Step1, Step2(sort), Step3 |
  | Step3のレシーバ名`mh` | `h` | Step3 (Step2(heap)に揃えた) |

  - 1行で終わる短いスコープの変数(`nf`など)やメソッドのレシーバ(`h`)は短いままにした
    - Goではこれらを短く書くのが慣習
  - 上のStep1/Step2のメモに出てくるコード片は、修正前の変数名のまま残している
