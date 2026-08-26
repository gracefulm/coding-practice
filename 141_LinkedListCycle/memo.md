# 141. Linked List Cycle

<https://leetcode.com/problems/linked-list-cycle/>

## Step1

- Passするのに30分以上かかった
- この方法しか思いつかなかったため、ループをネストさせナイーブに実装
- 良い変数名が思いつかずにcurrent/current2という変数名になってしまった
  - イケていないとは思いつつ...

## Step2

- 別の実装方法が思いつかなかったためAIで学習
- 以下のコメントの通りFloydの実装は、アルゴリズムを知っていないと基本的に実装できないため思いつかなくても問題ない
  - <https://github.com/tk-hirom/Arai60/pull/1#discussion_r1641231416>
  - Brentの実装はFloydの実装の改良版のためこれも同じ
  - ※Floyd/Brentの実装はちゃんと理解できてない
- 標準的な実装方法はHashMapを使った方法で、これは普通に書けなければ危機感を持った方が良いらしい

## Step3

- tab補完や、コードハイライトのないエディタで実装
- 1回目
  - Mapを使った方法で実装を試みて5分適度で実装完了するも、細かな構文エラーや実装ミスが複数あり
- 2回目
  - OK
