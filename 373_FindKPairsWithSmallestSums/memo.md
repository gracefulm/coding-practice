# 373. Find K Pairs with Smallest Sums

<https://leetcode.com/problems/find-k-pairs-with-smallest-sums/>

## Step1

- ナイーブに全ペアを最小ヒープにいれてK回Popして配列化して返す実装をするもOOMでNA
  - LeetCode 373 の制約は nums1.length, nums2.length <= 10^5 なので、最大で 10^5 × 10^5 = 10^10 個のペアを作ることになり、1要素あたり16バイト（int 2つ）で約160GB
- 2つの配列がソートされている前提をうまく活用するために試行錯誤したが1時間くらい考えて結局実装できなかったので諦めた
- 仮にメモリが問題なくても、時間計算量がO(n·m·log(n·m))でTLEになる可能性が高い
  - 制約の上限 n = m = 10^5 を当てはめると、次のようになります。
  - ペア数：n·m = 10^10
  - log₂(10^10) ≈ 33
  - 合計で約 3.3 × 10^11 回の操作
  - Go で単純な操作がおおよそ 1秒に 10^8〜10^9 回程度とすると、数百秒〜数千秒かかります。LeetCode の制限は通常数秒程度です。
  - ヒープを使わなくても、10^10 回ループを回すだけで十数秒以上かかるので、全ペアを列挙する方針自体が成り立ちません。

## Step2

- AIレビューと他の人のPRを確認する
- 見たPR
  - <https://github.com/miyataka/coding-practice/pull/10>
  - <https://github.com/n6o/leetcode_arai60/pull/10>

[レビューコメント1](https://github.com/n6o/leetcode_arai60/pull/10/changes#r2781856580)

> 実装前に時間計算量は求めましたでしょうか。また、時間計算量から実行時間は推定されましたでしょうか。

なんとなく頭で、「計算量多そう」程度しか考えてなかったので次から意識する

[レビューコメント2](https://github.com/n6o/leetcode_arai60/pull/10/changes#r2781867653)

> pair は通常 2 つの要素を持つデータ構造だと思います。 IndexesAndSum や IndexPairAndSum はいかがでしょうか？

### 計算量

- n = `len(nums1)`、m = `len(nums2)`、ヒープの最大サイズを h = min(n, k) とする
- 時間計算量: O(k · log min(n, k))
  - 初期化: h 回 Push、1回 O(log h) で O(h · log h)
  - メインループ: 1周で結果が1つ増えるので最大 k 回。1周で Pop 1回 + Push 最大1回、どちらも O(log h) で O(k · log h)
  - h ≤ k なので初期化はメインループ以下になり、全体は O(k · log h)
  - 見積もり: k = h = 10^4 で 10^4 × log₂(10^4) ≈ 1.4 × 10^5 回。Push + Pop 1組 約 73 ns（[実測値](../notes/complexity-and-runtime.md)）× 10^4 ≈ 1ms 未満
- 空間計算量: O(k)。返り値を除けば O(min(n, k))
  - ヒープ: O(h)。最初に h 個 Push したあとは Pop 1回につき Push 最大1回なので、サイズは h から増えない
  - 返り値: ペア k 個で O(k)
  - 計算量を問われたら、返り値を含めるかどうかを明確にする
- Step1 との比較

| | 時間 | 空間（返り値を除く） |
|---|---|---|
| Step1 | O(n·m · log(n·m)) | O(n·m) |
| Step2 | O(k · log min(n, k)) | O(min(n, k)) |

- `nums2` の長さ m に依存しなくなった

## Step3
