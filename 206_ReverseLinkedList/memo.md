# 206. Reverse Linked List

<https://leetcode.com/problems/reverse-linked-list/>

## Step1

- 20分くらいで何回か提出してPass
- 考えたこと
  - ループで回しながらNextを前のノードに入れ替えていく
  - リストのNextを動的に繋ぎ変えていくには、Nextを入れ替える前に保持しておく必要がある
  - 現在のノードがnilになったらループを抜ける
  - 現在のノードがnilになったということは、一つ前のノードが先頭(戻り値)になる
  - 次の場合は何もしないので早期リターンする
    - headがnil
    - head.Nextがnil
  - 引数に破壊的変更を加えてしまうがStep1としてはとりあえずOKとする

## Step2

- 他の人のPRを見て気づいたこと
  - 確認したPR
    - <https://github.com/minsrd/arai60/pull/7>
    - <https://github.com/miyataka/coding-practice/pull/7>
  - Step1の`prev`という変数名は戻り値に使うため`reversedHead`などの方が情報量が多くて良い
  - Step1で`if head == nil || head.Next == nil {..}`の早期リターンをしているが、もっとシンプルに描ける
  - 再帰でも実装できる
- 3種類の方法で実装をし直し
  - `reverseList`
    - Step1の方針で書き方を改善
  - `reverseListWithoutSideEffect`
    - 引数の破壊的変更がない実装
    - 1番シンプルに書けたのと、相当な理由がない限り引数の破壊的な変更はすべきではないので、特に制約がない限りこの実装が一番良さそう
  - `reverseListRecursive`
    - 再帰を使った実装
      - 他の人のPRを見て実装
    - 再帰は苦手
    - 書き方がうまく思いつかない

## Step3

- Step2の`reverseListWithoutSideEffect`の実装で3回
