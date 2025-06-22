# 🔧 汎用コードパターンと問題解決

## 📋 汎用パターン一覧

### 🔗 環境変数パターン
```{lang}
# パターン1: 直接取得
{var1} = {env_access}('{ENV_VAR1}')

# パターン2: デフォルト値付き
{var2} = {env_access}('{ENV_VAR2}', '{default_value}')

# パターン3: bool変換
{var3} = {env_access}('{ENV_VAR3}', 'false').lower() == 'true'

# パターン4: int変換
{var4} = {type_conversion}({env_access}('{ENV_VAR4}', '{default_number}'))
```

### 🏗️ 非同期初期化パターン
```{lang}
{async_keyword} {init_method}({self_param}):
    {try_keyword}:
        # 1. 外部サービス接続
        {await_keyword} {self_param}.{connect_method1}()
        {await_keyword} {self_param}.{connect_method2}()
        
        # 2. 設定読み込み・検証
        {self_param}.{load_config_method}()
        
        # 3. 接続テスト
        {await_keyword} {self_param}.{health_check_method}()
        
        {logger}.{info_method}("✅ Initialization successful")
        
    {except_keyword} {exception_type} {as_keyword} {error_var}:
        {logger}.{error_method}(f"❌ Initialization failed: {{error_var}}")
        {raise_keyword}
```

### 📊 HTTP API呼び出しパターン
```{lang}
{async_keyword} {make_request_method}({self_param}, {method_param}, {endpoint_param}, {params_param}={none_value}, {auth_param}={true_value}):
    {url_var} = f"{{self_param}.{base_url_attr}}}{{endpoint_param}}"
    
    {if_keyword} {auth_param}:
        {headers_var} = {self_param}.{get_auth_headers_method}()
    {else_keyword}:
        {headers_var} = {{"{content_type_header}": "{content_type_value}"}}
    
    {async_with_keyword} {self_param}.{session_attr}.{request_method}({method_param}, {url_var}, **{kwargs_var}) {as_keyword} {response_var}:
        {data_var} = {await_keyword} {response_var}.{json_method}()
        
        {if_keyword} {not_keyword} {response_var}.{ok_attr}:
            {raise_keyword} {api_error_class}(f"HTTP {{response_var}.{status_attr}}}: {{data_var}}")
            
        {return_keyword} {data_var}
```

### 🎯 抽象基底クラスパターン
```{lang}
{from_keyword} {abc_module} {import_keyword} {abc_class}, {abstract_method_decorator}

{class_keyword} {base_class_name}({abc_class}):
    {abstract_method_decorator}
    {async_keyword} {analyze_method}({self_param}, {data_param}) -> {signal_type}:
        """{docstring_text}"""
        {pass_keyword}
    
    {def_keyword} {common_method}({self_param}):
        """{common_logic_docstring}"""
        {pass_keyword}
```

## 🚨 汎用的な問題と解決パターン

### 1. 環境変数が読み込まれない
```{lang}
# 問題: {env_access}() が {none_value} を返す
# 原因: {env_loader} が読み込まれていない

# 解決: メインファイルで {load_env_method}()
{from_keyword} {env_loader_module} {import_keyword} {load_env_method}
{load_env_method}()  # これを忘れがち

# 確認コマンド
{env_check_command}
```

### 2. データベース/キャッシュ接続エラー
```{lang}
# 問題: "{auth_required_error}" / "{connection_refused_error}"
# 原因1: 認証設定ミスマッチ
# 原因2: サービス未起動

# 解決パターン:
# 1. サービス起動確認: {service_status_command}
# 2. 接続文字列確認: {connection_string_pattern}
# 3. ネットワーク確認: {network_test_command}
```

### 3. HTTP API認証エラー
```{lang}
# 問題: "{unauthorized_error}"
# 原因: パブリックAPIに認証を付けている、認証情報不正

# 解決: APIの性質に応じた認証設定
# パブリック: {auth_param}={false_value}
# プライベート: 正しいAPIキー・署名
```

### 4. ポート競合
```bash
# 問題: "{port_in_use_error}"
# 解決手順:
{port_check_command}  # 使用プロセス確認
{kill_process_command}  # プロセス終了
# または別ポート使用
```

### 5. 非同期処理エラー
```{lang}
# 問題: "{async_loop_error}"
# 原因: 既存ループ内で{run_async_method}()を呼び出し

# 解決:
# 新規ループ: {run_async_method}({func_call}())
# 既存ループ内: {await_keyword} {func_call}()
```

## 🔍 汎用デバッグコマンド

### 環境・設定確認
```bash
# 環境変数確認
{env_debug_command}

# データベース接続テスト
{db_test_command}

# キャッシュ接続テスト
{cache_test_command}

# API接続テスト
{api_test_command}
```

### システム状態確認
```bash
# プロセス確認
{process_check_command1}
{process_check_command2}

# ポート使用確認
{port_check_command1}
{port_check_command2}

# コンテナ状態確認
{container_status_command}
{container_logs_command}
```

### ログ・ファイル確認
```bash
# アプリケーションログ
{app_log_command}
{service_log_command}

# システムリソース
{resource_monitor_command}
{disk_usage_command}
{memory_usage_command}
```

## 🏃‍♂️ 汎用的な調査フロー

### 新機能実装時
```
1. 既存の類似機能を{search_command}検索
2. 既存パターンを参考にした実装
3. 設定・環境変数の確認と追加
4. エラーハンドリングの実装
5. ログ出力とテストの追加
```

### 問題調査時
```
1. エラーメッセージでコード検索
2. 該当箇所の詳細確認
3. 設定・環境の確認
4. 実際の動作状況確認
5. ドキュメント・過去事例確認
```

### 設定変更時
```
1. 設定項目の使用箇所を特定
2. 影響範囲とデフォルト値確認
3. バックアップ取得
4. 段階的変更とテスト実施
```

## 💡 汎用開発効率化のヒント

### 効果的なコード検索
```bash
# 機能別検索
{search_command} "{pattern1}\|{pattern2}" {search_path}
{search_command} "{pattern3}\|{pattern4}" . --include="*.{file_ext}"
{search_command} "{pattern5}\|{pattern6}" {log_path}

# パターン検索
{search_command} "{async_pattern}" {source_path}  # 非同期処理
{search_command} "{env_pattern}" {source_path}    # 環境変数使用箇所
```

### 素早い動作確認
```bash
# 設定確認スクリプト
{config_check_script}

# テスト実行
{test_command}

# ヘルスチェック
{health_check_command}

# ログ監視
{log_monitor_command}
```