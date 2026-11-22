---
name: create-pr
description: Create a GitHub pull request for a LeetCode problem in this coding-practice repo, filling .github/PULL_REQUEST_TEMPLATE.md and asking the user (via AskUserQuestion) which problem they will solve next. Use whenever the user asks to open/create a PR here — "PRを作成して", "PR出して", "プルリク作って", "create a PR" — including when combined with a commit request ("コミットしてPR作って").
---

# Create PR (coding-practice)

Open a pull request for the problem worked on in the current branch, using the
repository's PR template. The "次に解く問題" section must always be confirmed
with the user through **AskUserQuestion** — never guess it or leave it as TBD
without asking.

## Step 1 — Inspect the branch

Run in parallel:

```bash
git status --short
git branch --show-current
git log origin/main..HEAD --oneline
gh pr list --head "$(git branch --show-current)" --state open
```

- Uncommitted changes → commit them first with the `conventional-commit` skill.
- No commits ahead of `main` → nothing to open; tell the user and stop.
- An open PR already exists for this branch → show its URL and ask whether to
  update its body instead of creating a new one.

## Step 2 — Identify the problem

Problem directories are named `<number>_<PascalCaseTitle>` (e.g.
`373_FindKPairsWithSmallestSums`). Find the directory touched by the branch:

```bash
git diff --name-only origin/main...HEAD
```

Read the first lines of its `memo.md`; it starts with the title and URL:

```markdown
# 373. Find K Pairs with Smallest Sums

<https://leetcode.com/problems/find-k-pairs-with-smallest-sums/>
```

Take `<number>. <Title>` and the URL from there. If `memo.md` lacks them, derive
the title from the directory name and the URL slug from the title (lowercase,
spaces → `-`), and mention the guess to the user.

If the branch touches no problem directory (e.g. notes or tooling), this skill
does not apply as-is: write a normal title/body and skip the template sections
that do not fit.

## Step 3 — Ask for the next problem (required)

Call **AskUserQuestion** with one question:

- question: `次に解く問題は何ですか？（Other に「番号. タイトル」または LeetCode の URL を入力してください）`
- header: `次の問題`
- options:
  - `未定` — 「次に解く問題の予告など」を空欄のままにする
  - `あとで追記` — `TBD` と書いておく

The user will normally type the problem through **Other**. Turn the answer into a
link `[<number>. <Title>](https://leetcode.com/problems/<slug>/)`:

- A URL given → take the slug from it; use the official title for the link text.
- `番号. タイトル` given → build the slug from the title.
- Only a number or only a title → fill in the other part from your knowledge of
  LeetCode and state it in the final message so the user can catch mistakes.

## Step 4 — Build the body from the template

Read `.github/PULL_REQUEST_TEMPLATE.md` (do not hard-code it — the template may
change) and fill it:

- Replace each `<!-- ... -->` guidance comment with the content for that
  section; remove the comment itself.
- `## 問題` → `[<number>. <Title>](<url>)`
- `## 問題文概要(プレミアムの場合)` → keep `nop` unless the problem is a premium
  one; then write a short summary of the problem statement.
- `## 次に解く問題の予告など` → the link from Step 3 (empty for `未定`, `TBD` for
  `あとで追記`).
- Keep any other sections from the template, filled sensibly.
- End the body with the attribution line required by the current session, if
  any (e.g. `Generated with Claude Code`).

Write the body to `$TMPDIR/pr-body.md`.

## Step 5 — Push and create

- Title: `<number>. <Title>` (e.g. `373. Find K Pairs with Smallest Sums`).
- Base: `main`.

```bash
git push -u origin "$(git branch --show-current)"
gh pr create --base main --title "<number>. <Title>" --body-file "$TMPDIR/pr-body.md"
```

If `git push -u` reports that it could not write `.git/config` (sandbox), the
push itself still succeeded; mention that the upstream was not recorded.

Finish by giving the user the PR URL and, in one line, the next-problem link
that went into the body.
