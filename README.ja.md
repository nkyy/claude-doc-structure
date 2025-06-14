# Claude Doc Structure 📋

Claude Code との効率的な協働のための包括的なドキュメントツールキット。整理されたプロジェクトコンテキスト、自動化ツール、再利用可能なテンプレートで開発ワークフローを効率化します。

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Python](https://img.shields.io/badge/Python-3.8%2B-blue.svg)](https://www.python.org/)
[![CLI Tool](https://img.shields.io/badge/CLI-Tool-brightgreen.svg)](https://github.com/your-username/claude-doc-structure)

## 🎯 目的

このテンプレートは、Claude Code の効果を最大化するプロジェクトドキュメントの構造化アプローチを提供します。これらのパターンに従うことで、以下が可能になります：

- Claude Code がプロジェクトコンテキストをより迅速に理解できる
- 整理されたドキュメントにより反復的な説明を削減
- 開発セッション間の一貫性維持
- プロジェクト成長に合わせたドキュメントの効率的なスケーリング

## 🚀 クイックスタート

### 方法 1：CLI ツールを使用（推奨）

1. **ツールをインストール：**

   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cd claude-doc-structure
   make install
   ```

2. **プロジェクトを初期化：**

   ```bash
   claude-docs init my-project
   ```

3. **ツールを使い始める：**

   ```bash
   # テンプレート生成
   claude-docs template api users

   # 大きなドキュメントを分割
   claude-docs split README.md --by-headers

   # 複数のドキュメントを統合
   claude-docs merge specs/ --output combined.md

   # 構造を検証
   claude-docs validate
   ```

### 方法 2：手動セットアップ

1. **構造をコピー**してプロジェクトに追加：

   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cp -r claude-doc-structure/.claude ./
   cp claude-doc-structure/CLAUDE.md ./
   ```

2. **CLAUDE.md をカスタマイズ**してプロジェクト詳細を記載し、協働を開始！

## 📁 構造概要

```
your-project/
├── CLAUDE.md              # Claude Code用のメインプロジェクトコンテキスト
├── specs/                 # 詳細仕様書
│   ├── api.md            # APIドキュメント
│   └── screens.md        # UI/UX仕様書
├── .claude/              # Claude Code用アセット
│   ├── prompts/          # 再利用可能なプロンプトテンプレート
│   └── templates/        # ドキュメントテンプレート
├── scripts/              # ドキュメント用ユーティリティ
│   ├── split_docs.py     # 大きなドキュメントを小さなファイルに分割
│   └── merge_docs.py     # 包括的なビューのためのドキュメント統合
├── claude-docs           # 統合CLIツール
├── Makefile             # 便利なショートカット
└── setup.py             # パッケージ設定
```

## 📖 ドキュメントアプローチ

### 段階的ドキュメント戦略

**段階 1：単一ファイル（小規模プロジェクト）**

- `CLAUDE.md`のみを使用
- 迅速な参照のためにすべてを一箇所に保持
- 10,000 行未満のプロジェクトに最適

**段階 2：ドメイン別分割（中規模プロジェクト）**

- `specs/api.md`、`specs/screens.md`を作成
- `CLAUDE.md`から相対パスで参照
- 複数コンポーネントを持つプロジェクトに理想的

**段階 3：階層構造（大規模プロジェクト）**

- 機能とモジュール別に整理
- 自動化された整理のために`scripts/split_docs.py`を使用
- ドキュメント間の相互参照を維持

### ベストプラクティス

**Claude Code 統合のために：**

- 明確で説明的なヘッダーを使用
- 行番号付きのファイルパスを含める: `src/api/users.js:42`
- ドキュメント全体で一貫した用語を維持
- 最適なパフォーマンスのためファイルあたり 200KB 未満にコンテキストを保持

**ドキュメント作成：**

- 目的とコンテキストを先頭に記載
- 抽象的な説明よりも例を使用
- 更新に「何が変わったか」と「なぜか」を含める
- 実際のコード場所を参照

## 🛠️ ツール & コマンド

### 統合 CLI ツール

`claude-docs`コマンドですべての機能を一箇所で提供：

```bash
# プロジェクト初期化
claude-docs init [project-name]           # ドキュメント構造を作成
claude-docs validate [directory]          # ドキュメント構造を検証

# ドキュメント管理
claude-docs split <file> [options]        # 大きなドキュメントを分割
claude-docs merge <directory> [options]   # 複数のドキュメントを統合

# テンプレート生成
claude-docs template <type> [name]        # ドキュメントテンプレートを生成
```

### Makefile ショートカット

より簡単な使用のために提供された Makefile を使用：

```bash
make help                                  # すべての利用可能なコマンドを表示
make init                                  # ドキュメント構造を初期化
make split FILE=large-doc.md              # ドキュメントを分割
make merge DIR=specs/                     # ドキュメントを統合
make template TYPE=api NAME=users         # テンプレートを生成
make validate                             # 構造を検証
```

### ドキュメント管理スクリプト

**直接スクリプト使用（お好みで）：**

```bash
# 大きなドキュメントを分割
python scripts/split_docs.py specs/large-spec.md --by-headers --max-sections 5

# 包括的なレビューのためのマージ
python scripts/merge_docs.py specs/ --output combined.md
```

### テンプレートシステム

プロフェッショナルなドキュメントテンプレートを生成：

```bash
# APIドキュメントテンプレート
claude-docs template api user-management

# 画面仕様書テンプレート
claude-docs template screen dashboard

# 機能仕様書テンプレート
claude-docs template feature authentication
```

## 🌟 例 & ワークフロー

### 一般的なワークフロー

**1. 新しいプロジェクトの開始：**

```bash
# 構造を初期化
claude-docs init my-awesome-app

# 初期テンプレートを生成
claude-docs template api authentication
claude-docs template screen login
claude-docs template feature user-management

# すべてが正しく設定されているか検証
claude-docs validate
```

**2. 大きなドキュメントの管理：**

```bash
# 大きな仕様書を管理しやすいチャンクに分割
claude-docs split specs/massive-api-spec.md --by-headers --max-sections 8

# 後で包括的なレビューのために統合
claude-docs merge specs/ --output complete-api-docs.md
```

**3. チームコラボレーション：**

```bash
# 各チームメンバーが自分のセクションで作業
claude-docs merge specs/ --exclude work-in-progress.md --output team-review.md

# チームのドキュメント構造を検証
claude-docs validate
```

### 小規模プロジェクトの例

```markdown
# TodoApp 用の CLAUDE.md

## プロジェクト概要

ローカルストレージを使ったシンプルな React todo アプリケーション。

## 主要ファイル

- `src/App.js:1` - メインコンポーネント
- `src/components/TodoList.js:15` - Todo リストロジック
- `src/utils/storage.js:8` - ローカルストレージユーティリティ

## 現在の焦点

Firebase を使ったユーザー認証の追加。
```

### 生成されるテンプレートの例

`claude-docs template api users`を実行すると以下が生成されます：

```markdown
# Users API エンドポイント

## 概要

このエンドポイントの目的についての簡単な説明。

## HTTP メソッドと URL
```

GET/POST/PUT/DELETE /api/users

```

## パラメータ
### パスパラメータ
- `id` (string): ユーザー識別子

### クエリパラメータ
- `limit` (number, optional): ページあたりの結果数
- `offset` (number, optional): ページネーションオフセット
...
```

### 大規模プロジェクトの例

包括的なマルチサービスアプリケーション構造については`examples/large-project/`を参照してください。

## 🚀 インストール

### クイックインストール

```bash
git clone https://github.com/your-username/claude-doc-structure.git
cd claude-doc-structure
make install
```

### 手動インストール

```bash
pip install -e .
```

### 要件

- Python 3.8+
- 外部依存関係なし（標準ライブラリのみ使用）

## 🆘 トラブルシューティング

**インストール後にコマンドが見つからない：**

```bash
# スクリプトが実行可能か確認
chmod +x claude-docs

# PATHに追加するかフルパスを使用
export PATH=$PATH:$(pwd)
```

**権限エラー：**

```bash
# 適切な権限で実行
sudo make install
```

**開発用：**

```bash
# 開発モードでインストール
pip install -e ".[dev]"

# テスト実行
make test

# リント実行
make lint
```

## 🤝 貢献

貢献を歓迎します！始め方：

1. リポジトリをフォーク
2. 機能ブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add amazing feature'`)
4. ブランチにプッシュ (`git push origin feature/amazing-feature`)
5. プルリクエストを作成

詳細なガイドラインについては[CONTRIBUTING.md](.github/CONTRIBUTING.md)をご覧ください。

## 📝 ライセンス

MIT ライセンス - 詳細は[LICENSE](LICENSE)をご覧ください。

## 🔗 関連リソース

- [Claude Code ドキュメント](https://docs.anthropic.com/claude-code)
- [AI 支援開発のベストプラクティス](https://github.com/anthropics/claude-code)
- [Markdown ガイド](https://www.markdownguide.org/)

## 📊 機能一覧

| 機能                      | 説明                                               | コマンド               |
| ------------------------- | -------------------------------------------------- | ---------------------- |
| 🚀 **プロジェクト初期化** | 最適なドキュメント構造を作成                       | `claude-docs init`     |
| ✂️ **ドキュメント分割**   | 大きなドキュメントを管理しやすいピースに分割       | `claude-docs split`    |
| 🔄 **ドキュメント統合**   | レビュー用に複数のドキュメントを結合               | `claude-docs merge`    |
| 📝 **テンプレート生成**   | プロフェッショナルなドキュメントテンプレートを生成 | `claude-docs template` |
| ✅ **構造検証**           | ドキュメントのベストプラクティスを検証             | `claude-docs validate` |
| 🛠️ **Makefile サポート**  | すべての操作の便利なショートカット                 | `make help`            |
| 📦 **簡単インストール**   | シンプルなインストールプロセス                     | `make install`         |

---

_Claude Code で働く開発者のために ❤️ で作成_

_ドキュメントの混乱を Claude 最適化された明確さに変えましょう！_ ✨
