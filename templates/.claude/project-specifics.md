# 🎯 {name} プロジェクト固有情報

*このファイルは{name}プロジェクト特有の情報のみを記載*

## 📁 プロジェクト固有のファイル構造

### 核心コンポーネント
```
src/
├── main.{ext}              # メインエントリーポイント
├── {component1}.{ext}      # 主要コンポーネント1
├── {component2}.{ext}      # 主要コンポーネント2
└── {component3}.{ext}      # 主要コンポーネント3
```

### 設定・インフラ
```
config/
├── {config1}.{config_ext}  # 設定ファイル1
├── {config2}.{config_ext}  # 設定ファイル2
└── {config3}.{config_ext}  # 設定ファイル3

.env                        # 環境変数設定
{build_file}               # ビルド設定
```

## 🔧 プロジェクト固有の設定パターン

### 環境変数設定
```bash
# .env でのプロジェクト設定
{ENV_KEY1}={env_value1}
{ENV_KEY2}={env_value2}
{ENV_KEY3}={env_value3}

# プロジェクト固有パラメータ
{PROJECT_PARAM1}={param_value1}
{PROJECT_PARAM2}={param_value2}
```

### 外部サービス接続設定
```{lang}
# {service_name}接続設定の自動判定
def _get_{service}_config(self):
    # 1. 環境判定
    # 2. 設定読み込み
    # 3. 接続文字列生成
```

## 🎯 {name}固有の機能

### 主要機能1
```{lang}
class {MainClass}:
    """主要機能の説明"""
    
    # 独自の処理ロジック
    def {method_name}(self):
        # 処理内容
        pass
```

### 主要機能2
```{lang}
class {SecondaryClass}:
    """副次機能の説明"""
    
    # 処理フロー
    - {step1}
    - {step2}
    - {step3}
```

## 🔍 {name} 固有の検索パターン

### 機能関連
```bash
# 主要機能検索
grep -r "{keyword1}\|{keyword2}" src/

# 設定関連
grep -r "config\|setting\|{config_keyword}" src/

# API関連
grep -r "{api_keyword}\|{endpoint_pattern}" src/
```

### 外部サービス関連
```bash
# サービス固有機能
grep -r "{service_name}\|{service_keyword}" src/

# 認証関連
grep -r "auth\|token\|{auth_keyword}" src/

# データ処理
grep -r "{data_keyword}\|{process_keyword}" src/
```

## 🔧 {name} 固有のコードパターン

### 環境変数パターン
```{lang}
# {name}固有の環境変数
{var1} = os.getenv('{ENV_VAR1}')
{var2} = os.getenv('{ENV_VAR2}', '{default_value}')
{var3} = os.getenv('{ENV_VAR3}', '{default_value}').lower() == 'true'

# 場所: .env → {env_loader} → {env_access}
```

### API呼び出しパターン
```{lang}
# {service_name} API呼び出し
response = await self._make_request("{method}", "{endpoint}", params, require_auth={auth_required})

# 場所: src/{api_file}
```

### 処理パターン
```{lang}
# {name}固有の処理パターン
class {ProcessClass}({BaseClass}):
    async def {process_method}(self, {param}: {param_type}) -> {return_type}:
        # 1. {step1}
        # 2. {step2}
        # 3. {step3}
        # 4. {step4}

# 場所: src/{process_files}
```

## 🚨 {name} 固有のよくある問題

### 1. {problem1_title}
```
# 症状: {problem1_symptom}
# 原因: {problem1_cause}
# 解決: {problem1_solution}
# ファイル場所: {problem1_location}
```

### 2. {problem2_title}
```
# 症状: {problem2_symptom}
# 原因: {problem2_cause}
# 解決: 
# 1. {problem2_solution1}
# 2. {problem2_solution2}
# ファイル場所: {problem2_location}
```

### 3. {problem3_title}
```
# 症状: {problem3_symptom}
# 原因: {problem3_cause}
# 解決: {problem3_solution}
# ファイル場所: {problem3_location}
```

## 📊 プロジェクト固有のデータフロー

### システム起動シーケンス
```
1. {startup_step1}
   ↓ {startup_action1}
2. {startup_step2}
   ↓ {startup_action2}
3. {startup_step3}
   ↓ {startup_action3}
4. {startup_step4}
```

### 処理フロー
```
1. {process_step1}: {process_action1}
2. {process_step2}: {process_action2}
3. {process_step3}: {process_action3}
4. {process_step4}: {process_action4}
```

## 🎛️ プロジェクト固有の監視・制御

### ヘルスチェック
```bash
curl http://localhost:{port}/{health_endpoint}
# 戻り値: {health_response}
```

### 主要ログパターン
```
# 正常動作
"{success_log1}"
"{success_log2}"

# 警告・エラー
"{warning_log1}"
"{error_log1}"
```

### {name} 固有のデバッグコマンド
```bash
# 環境変数確認
{env_check_command}

# 接続テスト
{connection_test_command}

# ヘルスチェック
curl http://localhost:{port}/{health_endpoint}

# 設定確認
{config_check_command}
```

### {deployment_tool}管理コマンド
```bash
# 開発用構成
{dev_start_command}

# 本格運用構成
{prod_start_command}

# ログ確認
{log_check_command}

# システム起動
{system_start_command}
```