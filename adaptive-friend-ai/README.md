# Python音声アプリ + Mastra適応型友人AI統合プロトタイプ

## プロジェクト概要

既存のPython音声会話アプリと連携するMastraベースの適応型友人AIを実装するマイクロサービス構成のプロトタイプです。

このAIは時間をかけてユーザーの性格や好みを学習し、それに合わせて自分の話し方や性格を変化させる「成長する友人関係」を提供します。

## アーキテクチャ

```
[Python音声アプリ] ←→ HTTP/WebSocket ←→ [Node.js + Mastra AIエージェント]
     ↓                                        ↓
[ローカル音声処理]                          [適応型メモリー + 性格変化]
[音声認識・合成]                           [Working Memory + Tools]
```

## プロジェクト構造

```
adaptive-friend-ai/
├── mastra-backend/          # Node.js + Mastra AIエージェント
│   ├── src/
│   │   ├── agents/         # 適応型エージェント定義
│   │   ├── tools/          # カスタムツール
│   │   ├── memory/         # メモリー設定
│   │   ├── api/            # HTTP APIサーバー
│   │   └── types/          # TypeScript型定義
│   ├── package.json
│   └── tsconfig.json
├── python-client/          # Python連携クライアント
│   ├── mastra_client.py    # HTTP APIクライアント
│   ├── voice_app_example.py # 音声アプリ例
│   └── requirements.txt
├── docs/                   # ドキュメント
│   ├── API.md             # API仕様
│   ├── ARCHITECTURE.md    # アーキテクチャ詳細
│   └── TESTING.md         # テスト手順
└── README.md              # このファイル
```

## クイックスタート

### 1. Mastraバックエンド起動

```bash
cd mastra-backend
npm install
npm run dev
```

### 2. Python音声アプリの統合

```bash
cd python-client
pip install -r requirements.txt
python voice_app_example.py
```

### 3. API テスト

```bash
curl -X POST http://localhost:4111/chat \
  -H "Content-Type: application/json" \
  -d '{"userId": "user123", "message": "こんにちは！"}'
```

## 主要機能

### 適応型学習
- ユーザーの会話パターンと好みの自動学習
- 関係性レベルに応じた話し方の変化（丁寧→カジュアル→親友風）
- 個人化されたワーキングメモリーの構築

### マイクロサービス構成
- Python: ローカル音声処理（音声認識・合成）
- Node.js + Mastra: AIエージェント、適応学習、メモリー管理
- HTTP API: 疎結合な連携インターフェース

### パフォーマンス
- API応答時間: 2秒以内
- 同時ユーザー: 10名程度（プロトタイプ）
- 自動メモリー圧縮とセマンティック検索

## API エンドポイント

- `POST /chat` - メイン会話処理
- `GET /user/:userId/memory` - ユーザーメモリー確認
- `POST /user/:userId/reset` - メモリーリセット
- `GET /health` - ヘルスチェック

## 開発状況

- [x] プロジェクト構造設計
- [ ] Node.js + Mastraセットアップ
- [ ] 適応型メモリーシステム
- [ ] 適応型エージェントとツール
- [ ] HTTP APIサーバー
- [ ] Python連携クライアント
- [ ] テスト環境

## 技術スタック

**バックエンド:**
- Node.js + TypeScript
- Mastra Framework
- Express.js
- LibSQL (SQLite)

**クライアント:**
- Python 3.8+
- Requests
- Speech Recognition (音声処理例)

**AI/ML:**
- OpenAI GPT-4o
- Mastra Memory & Tools
- セマンティック検索

## ライセンス

MIT License

---

*適応型友人AI - 時間をかけて関係性を深めるAIパートナー*