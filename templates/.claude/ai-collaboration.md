# 🤖 AI 協働開発のベストプラクティス

## ⚠️ 重要な行動指針

### 📁 .claude 配下ドキュメント作成時の注意

```
✅ 汎用的内容（他プロジェクトでも使える）
- 検索手法、デバッグ方法
- 一般的なコードパターン
- AI効率化のための手法

❌ プロジェクト固有内容（{name}特有）
- 具体的なファイル名・クラス名
- プロジェクト固有の設定・機能
- 特定APIやビジネスロジック

👉 プロジェクト固有内容は:
- `.claude/project-specifics.md` に記載
- または `specs/` 配下に仕様として記載
```

### 🔍 事実確認の徹底

```
❌ 絶対にやってはいけないこと：
- 設定や状態を推測して回答する
- 実際の確認なしに「動作している」と報告する
- 憶測に基づいて機能の説明をする

✅ 必ず行うこと：
- 実際のファイル内容を読んで確認
- コマンド実行結果を確認してから回答
- 設定の動作状況を具体的にテスト
- 不明な点は「確認します」と明言する
```

### 📝 正確性の原則

```
1. 実際の設定確認: ファイルを読み、環境変数を確認
2. 動作テスト: APIエンドポイントやコマンドを実際に実行
3. 証拠に基づく回答: ログ出力や結果を示して説明
4. 推測の排除: 「おそらく」「多分」は使わない
```

## 🎯 効果的な AI 活用方法

### 📋 プロンプト設計の原則

#### 具体的で段階的な指示

```markdown
# ✅ 良い例

「{feature_name}を実装してください。以下の仕様に従って：

1. {requirement1}
2. {requirement2}
3. {requirement3}
4. {requirement4}
5. 戻り値は {return_type}」

# ❌ 悪い例

「{vague_request}」
```

#### コンテキストの事前提供

```markdown
# ✅ 効果的なコンテキスト設定

「このプロジェクトは{tech_stack}で書かれた{project_type}です。
既存の {base_class} クラスを継承して、{feature_name}を実装してください。

- 基底クラス: {base_class_path}
- 参考実装: {example_path}
- テスト例: {test_path}」
```

### 🔄 反復開発のパターン

#### 段階的実装アプローチ

```
Phase 1: スケルトン実装 → 基本構造とインターフェース
Phase 2: コア機能実装 → 主要ロジックの実装
Phase 3: エラーハンドリング → 例外処理と検証
Phase 4: テスト追加 → 単体・統合テスト
Phase 5: 最適化 → パフォーマンス改善
```

#### フィードバックループの活用

```markdown
# 実装後の改善サイクル

1. 「この実装をレビューして、改善点を提案してください」
2. 「パフォーマンスの問題はありますか？」
3. 「セキュリティの観点で問題はありますか？」
4. 「テスト可能性を高めるには？」
5. 「保守性を向上させるには？」
```

## 🛠️ コード品質向上のテクニック

### 📝 コードレビュー依頼の方法

#### 効果的なレビュー依頼

```markdown
# ✅ 構造化されたレビュー依頼

「以下の {class_name} クラスをレビューしてください：

**確認してほしい点：**

1. {review_point1}
2. {review_point2}
3. {review_point3}
4. {review_point4}

**特に心配な部分：**

- line {line_range1} の{concern1}
- line {line_range2} の{concern2}

**コード：**
[実装コードを貼り付け]
```

#### パフォーマンス最適化の相談

```markdown
# ✅ 具体的なパフォーマンス相談

「この{process_name}が遅い（{current_time}→ 目標 {target_time}）：

**現在の処理フロー：**

1. {step1} ({time1})
2. {step2} ({time2}) ← ボトルネック
3. {step3} ({time3})

**制約条件：**

- {constraint1}
- {constraint2}
- {constraint3}

最適化のアイデアを教えてください。」
```

### 🧪 テスト戦略の相談

#### テスト設計の依頼

```markdown
# ✅ 包括的なテスト設計相談

「{class_name} の包括的なテストを設計してください：

**テスト対象の機能：**

- {function1}
- {function2}
- {function3}
- {function4}

**テストシナリオ：**

- 正常系: {normal_case}
- 異常系: {error_case}
- 境界値: {boundary_case}

**期待するテスト形式：**

- {test_framework}
- {test_approach}
- {test_structure}」
```

## 📚 ドキュメント作成の活用

### 📖 技術文書の段階的作成

#### アーキテクチャ設計書の作成依頼

```markdown
# ✅ 段階的なドキュメント作成

「システムアーキテクチャ設計書を作成してください：

**Phase 1: 概要レベル**

- {overview_item1}
- {overview_item2}
- {overview_item3}

**Phase 2: 詳細設計**

- {detail_item1}
- {detail_item2}
- {detail_item3}

**Phase 3: 運用観点**

- {operation_item1}
- {operation_item2}
- {operation_item3}

**対象読者：** {target_audience}
**フォーマット：** {document_format}」
```

#### API 仕様書の自動生成

```markdown
# ✅ 実装からドキュメント生成

「以下の{framework}コードから完全な API 仕様書を生成してください：

**含めてほしい内容：**

- {api_content1}
- {api_content2}
- {api_content3}
- {api_content4}
- {api_content5}

**形式：** {api_format}」
```

## 🐛 デバッグとトラブルシューティング

### 🔍 問題解決のアプローチ

#### エラー分析の依頼

```markdown
# ✅ 構造化されたエラー分析

「以下のエラーの原因と解決策を教えてください：

**エラー内容：**
```
{error_message}
{stack_trace}
```

**発生条件：**
- {condition1}
- {condition2}
- {condition3}

**環境情報：**
- {env_info1}
- {env_info2}
- {env_info3}

**実行中の関連コード：**
[関連コードを貼り付け]

**求める回答：**
1. 根本原因の特定
2. 具体的な修正方法
3. 再発防止策
```

#### パフォーマンス問題の診断

```markdown
# ✅ パフォーマンス診断依頼

「システムが徐々に遅くなる問題を調査してください：

**症状：**

- {symptom1}
- {symptom2}
- {symptom3}

**監視データ：**

- {monitoring_data1}
- {monitoring_data2}
- {monitoring_data3}

**疑わしい箇所：**

1. {suspect1}
2. {suspect2}
3. {suspect3}

プロファイリング方法と修正案を教えてください。」
```

## ⚙️ 設定と DevOps

### 🚀 デプロイメント自動化

#### CI/CD パイプライン設計

```markdown
# ✅ 段階的な CI/CD 設計

「{tech_stack}アプリの CI/CD パイプラインを設計してください：

**要件：**

1. {cicd_requirement1}
2. {cicd_requirement2}
3. {cicd_requirement3}
4. {cicd_requirement4}
5. {cicd_requirement5}

**現在の構成：**

- アプリ: {app_stack}
- DB: {database}
- インフラ: {infrastructure}
- 監視: {monitoring}

**特別な要件：**

- {special_requirement1}
- {special_requirement2}
- {special_requirement3}

段階的に設計してください。」
```

### 🔧 環境構築の自動化

#### Docker イメージ最適化

```markdown
# ✅ Docker 最適化相談

「以下の Dockerfile を最適化してください：

**現在の問題：**

- {docker_problem1}
- {docker_problem2}
- {docker_problem3}

**要件：**

- {docker_requirement1}
- {docker_requirement2}
- {docker_requirement3}
- {docker_requirement4}

**現在の Dockerfile：**
[Dockerfile を貼り付け]

{optimization_techniques}を含めて改善してください。」
```

## 🔐 セキュリティ強化

### 🛡️ セキュリティレビュー

#### セキュリティ監査の依頼

```markdown
# ✅ 包括的セキュリティ監査

「{project_type}のセキュリティを監査してください：

**監査対象：**

1. {security_target1}
2. {security_target2}
3. {security_target3}
4. {security_target4}
5. {security_target5}

**特に重要な資産：**

- {critical_asset1}
- {critical_asset2}
- {critical_asset3}

**コンプライアンス要件：**

- {compliance1}
- {compliance2}

**現在のセキュリティ実装：**
[実装詳細を貼り付け]

脆弱性と改善提案をお願いします。」
```

## 📊 データ分析・最適化

### 📈 パフォーマンス分析

#### 最適化の相談

```markdown
# ✅ データ駆動の最適化

「{feature_name}のパフォーマンスを改善してください：

**現在の成績：**

- {metric1}: {current_value1}（目標: {target_value1}）
- {metric2}: {current_value2}（目標: {target_value2}）
- {metric3}: {current_value3}（目標: {target_value3}）

**利用可能なデータ：**

- {available_data1}
- {available_data2}
- {available_data3}

**現在のパラメータ：**

- {param1}: {param_value1}
- {param2}: {param_value2}
- {param3}: {param_value3}

**分析依頼：**

1. {analysis_request1}
2. {analysis_request2}
3. {analysis_request3}
4. {analysis_request4}

{analysis_methods}も含めて検討してください。」
```

## 🎯 プロジェクト管理・計画

### 📅 開発計画の立案

#### マイルストーン設計

```markdown
# ✅ リアルな開発計画

「{project_type}の MVP 開発計画を立ててください：

**目標：**

- {goal_timeline}
- {goal_team_size}
- {goal_constraint}

**必須機能：**

1. {required_feature1}
2. {required_feature2}
3. {required_feature3}
4. {required_feature4}
5. {required_feature5}

**技術制約：**

- {tech_constraint1}
- {tech_constraint2}
- {tech_constraint3}

**リスク要因：**

- {risk_factor1}
- {risk_factor2}
- {risk_factor3}

週単位のマイルストーンで計画してください。」
```

---

## 🎓 学習効果を高めるコツ

### 💡 知識の定着方法

1. **段階的学習**: 複雑な概念は小さな部分に分けて質問
2. **実践的応用**: 学んだ概念を実際のコードに適用
3. **多角的視点**: 同じ問題を異なる観点から検討
4. **継続的改善**: 実装 → レビュー → 改善のサイクル

### 🔄 効率的な協働パターン

1. **事前準備**: 質問前に既存情報を整理
2. **具体的指示**: 曖昧な依頼より具体的な要求
3. **文脈共有**: プロジェクトの背景と制約を明示
4. **段階的発展**: 簡単な実装から徐々に複雑化

**これらの方法論により、AI 協働開発の効率性と品質を大幅に向上させることができます。**