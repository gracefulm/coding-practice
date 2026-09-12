# 82. Remove Duplicates from Sorted List II

<https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/>

## Step1

- x 引数のリンクトリストを繋ぎかえる方法で考えたが、5分くらい考えて実装方法が思いつかずに断念
  - 空間計算量: O(1)
- 新しheadをループで回しながら、重複した要素をスキップしながら新しいリストを作成する方針で実装
  - 空間計算量: O(n)
  - 以下のような細かいパターンを考慮できずに何回も失敗
    - 要素が1つのパターン: [1]
    - 最後の要素のみ残るパターン: [1,1,1,2]
  - また、以下の要因によりネストが深くなりコードがごちゃついてしまった
    - リストの先頭要素がforの中で決定するので1ループ目に特別な処理が必要
    - forの`current.Next!=nil`の条件により、最後の要素が抜けてしまう回避策として、ループの最後に特別な処理が必要
- 直感的にも読みにくく納得のいっていないコードだが、とりあえずPassすることができたのでStep1としては完了

## Step2

- step1をAIにレビューさせた結果**番兵(sentinel)パターン**を使うことでネストした分岐がなくなりスッキリかけることがわかった
  - 番兵パターンの使用例はGoの標準ライブラリの`container/list`などがある
    - <https://cs.opensource.google/go/go/+/refs/tags/go1.26.0:src/container/list/list.go;l=46-51>
- 引数のリストを繋ぎ直す`deleteDuplicates_2_withSideEffect`と新しいリストを作成する`deleteDuplicates_2_withoutSideEffect`を作成
  - 空間計算量的に、「元のリストを書き換えたくない」などの要件がない限り`deleteDuplicates_2_withSideEffect`を採用するのが良さそう
  - とりあえずStep3ではこれらの実装をスラスラかけるようにする

## Step3

- 時間を空けて5回くらい書いてようやく一発でPassできるようになった
