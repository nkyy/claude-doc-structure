# Claude Doc Structure 📋

Claude Code との効率的な協働のための包括的なドキュメントツールキット。整理されたプロジェクトコンテキスト、自動化ツール、再利用可能なテンプレートで開発ワークフローを効率化します。

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Go](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![CLI Tool](https://img.shields.io/badge/CLI-Tool-brightgreen.svg)](https://github.com/your-username/claude-doc-structure)
[![Zero Dependencies](https://img.shields.io/badge/Dependencies-Zero-green.svg)]()

## 🎯 目的

このテンプレートは、Claude Code の効果を最大化するプロジェクトドキュメントの構造化アプローチを提供します。これらのパターンに従うことで、以下が可能になります：

- Claude Code がプロジェクトコンテキストをより迅速に理解できる
- 整理されたドキュメントにより反復的な説明を削減
- 開発セッション間の一貫性維持
- プロジェクト成長に合わせたドキュメントの効率的なスケーリング

## 🚀 クイックスタート

### 方法 1：バイナリダウンロード（推奨）

1. **お使いのプラットフォーム用のバイナリをダウンロード** [Releases](https://github.com/your-username/claude-doc-structure/releases)から

2. **実行可能にしてPATHに追加：**
   ```bash
   chmod +x claude-docs
   sudo mv claude-docs /usr/local/bin/
   ```

3. **すぐに使用開始：**
   ```bash
   claude-docs init my-project
   claude-docs template api users
   claude-docs validate
   ```

### 方法 2：ソースからビルド

1. **クローンしてビルド：**
   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cd claude-doc-structure
   make build
   ```

2. **バイナリを使用：**
   ```bash
   ./bin/claude-docs init my-project
   ```

### 方法 3：手動セットアップ

1. **構造をコピー**してプロジェクトに追加：

   ```bash
   git clone https://github.com/your-username/claude-doc-structure.git
   cp -r claude-doc-structure/templates/* ./
   ```

2. **CLAUDE.md をカスタマイズ**してプロジェクト詳細を記載し、協働を開始！

## 📁 構造概要

```
your-project/
├── CLAUDE.md              # Claude Code用のメインプロジェクトコンテキスト
├── specs/                 # 詳細仕様書
│   ├── api.md            # APIドキュメント
│   └── screens.md        # UI/UX仕様書
└── .claude/              # Claude Code用アセット
    ├── prompts/          # 再利用可能なプロンプトテンプレート
    └── templates/        # ドキュメントテンプレート

claude-doc-structure/ (このリポジトリ)
├── main.go               # CLIエントリーポイント
├── go.mod               # Goモジュール定義
├── cmd/                 # CLIコマンド
├── internal/            # ドキュメント処理ロジック
├── bin/claude-docs      # ビルドされたバイナリ (make build後)
└── Makefile            # ビルド自動化
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
- 自動化された整理のために`claude-docs split`を使用
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

### CLI オプション & 機能

**ドキュメント分割：**
```bash
claude-docs split large-doc.md --by-headers --max-sections 8
claude-docs split large-doc.md --by-size --max-size-kb 50
claude-docs split large-doc.md --by-lines --lines-per-file 200
```

**ドキュメント統合：**
```bash
claude-docs merge specs/ --output combined.md
claude-docs merge docs/ --recursive --exclude "*.draft.md"
claude-docs merge . --pattern "*.md" --no-claude-optimization
```

**クロスプラットフォームビルド：**
```bash
make release    # Linux、macOS、Windows (x64 & ARM64) 用にビルド
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

### バイナリインストール（推奨）
```bash
# GitHub Releasesからダウンロード
curl -L https://github.com/your-username/claude-doc-structure/releases/latest/download/claude-docs-linux-amd64 -o claude-docs
chmod +x claude-docs
sudo mv claude-docs /usr/local/bin/
```

### ソースからビルド
```bash
git clone https://github.com/your-username/claude-doc-structure.git
cd claude-doc-structure
make build
# バイナリは ./bin/claude-docs に生成されます
```

### 要件
- **ランタイム**: なし（単一バイナリ、ゼロ依存）
- **ビルド**: Go 1.21+（ソースからビルドする場合のみ）

## 🆘 トラブルシューティング

**インストール後にコマンドが見つからない：**

```bash
# バイナリが実行可能でPATHにあるか確認
chmod +x claude-docs
which claude-docs  # パスが表示されるはず

# またはフルパスを使用
./bin/claude-docs --help
```

**権限エラー：**

```bash
# /usr/local/bin/ への書き込み権限があるか確認
sudo mv claude-docs /usr/local/bin/
```

**開発用：**

```bash
# ビルドとテスト
make build
./bin/claude-docs --help

# テスト実行
make test

# コードのフォーマットとリント
make fmt
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
| 🏗️ **クロスプラットフォーム** | Linux、macOS、Windows用の単一バイナリ          | `make release`         |
| ⚡ **ゼロ依存**           | ランタイム依存関係不要                             | ダウンロード & 実行    |
| 🎯 **Claude最適化**       | Claude Codeワークフロー専用設計                   | 全コマンド             |

---

_Claude Code で働く開発者のために ❤️ で作成_

_ドキュメントの混乱を Claude 最適化された明確さに変えましょう！_ ✨
