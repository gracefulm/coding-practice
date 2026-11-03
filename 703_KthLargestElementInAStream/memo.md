# 703. Kth Largest Element in a Stream

<https://leetcode.com/problems/kth-largest-element-in-a-stream/>

## Step1

- 考えたこと
  - KthLargestの要素としてスコアのスライスとkがあればあとは何とかなる
  - 降順ソートして先頭からk番目(k-1番目)の要素を返せば良い
  - とりあえず標準パッケージのソートを使う
  - Addするときはすでにソートされたスライスに要素を追加するので、より良いソートアルゴリズムがあるかもしれないがとりあえず標準パッケージを使う

## Step2

- Step1のレビュー
  - Constructor が呼び出し元の nums を書き換えてしまう
  - Add のたびに全体をソートし直している
  - k 番目より小さい要素は不要ないのに、すべて持ち続けている
  - sort.Sliceは非推奨でslice.SortFuncなどsliceパッケージを使うべき
    - ちなみにsort.Sliceやslice.SortFuncで使われるアルゴリズムはpdqsort(pattern-defeating quicksort)というアルゴリズム
      - 今回のように「ソート済みの配列の末尾に1件足しただけ」の入力なら、pdqsort はほぼ O(n) で終わる
      - <https://www.m3tech.blog/entry/pdqsort>
- 他の人のPRを読む
  - PR
    - <https://github.com/miyataka/coding-practice/pull/8>
    - <https://github.com/minsrd/arai60/pull/8/>
  - この手の問題は二分探索と最小ヒープを使うのが常識とのこと
    - とりあえずこの2種類の方法を書けるようになると良さそう

| 方法                  | 比較      | 移動                         | 全体の計算量 |
|-----------------------|-----------|------------------------------|--------------|
| 毎回ソート（pdqsort） | O(n)      | O(n)：1つずつ入れ替え        | O(n)         |
| 二分探索で挿入        | O(log n)  | O(n)：copy でまとめてずらす  | O(n)         |
| 最小ヒープ            | O(log k)  | O(log k)                     | O(log k)     |

## Step3

- 二分探索と最小ヒープを使った方法を実装
