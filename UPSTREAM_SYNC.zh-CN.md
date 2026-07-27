# 跟进上游更新

汉化分支以 `vxcontrol/pentagi:main` 为上游基准，保留英文词典作为回退和差异参照。

## 本地同步

请先提交当前改动，再在 `zh-CN` 分支执行：

```bash
node scripts/sync-upstream.mjs
```

脚本会拉取 `origin/main`，以非快进方式合并，并运行前端英文基线、提供商测试标签、安装器汉化门禁以及 REST/GraphQL API 错误响应门禁。若上游新增了用户可见英文，合并会停在提交前：前端文案需接入 `en.ts` 和 `zh-CN.ts`，安装器文案需接入集中语言文件，新增的稳定 REST API 错误码需在 `backend/pkg/server/response/locale.go` 中提供中文响应。GraphQL 未知内部错误会在中文界面使用统一中文提示，原始英文只保留在技术日志中。品牌、URL、模型 ID、命令与枚举实际值保持原样。共享组件中已被实际调用方覆盖的英文回退、AI 系统提示词与工具协议、服务端技术日志、自动生成的 Swagger 文档，以及从第三方项目引入的原版 Grafana 看板也不纳入汉化，避免改变程序契约或制造无意义的上游冲突。只检查是否有更新时使用：

```bash
node scripts/sync-upstream.mjs --check
```

如果上游使用其他远程名称，可显式指定：

```bash
node scripts/sync-upstream.mjs --remote=upstream --branch=main
```

## 自动检查

`.github/workflows/upstream-sync.yml` 每小时检查一次官方主线。有新提交时，它会在个人仓库创建或更新同步 PR；主 CI 会运行词典测试、前端英文基线、提供商测试标签、安装器汉化门禁和 REST/GraphQL API 错误响应门禁。GitHub 只会从仓库默认分支加载计划任务，因此请确保该工作流已存在于默认分支（个人汉化仓库建议将 `zh-CN` 设为默认分支）。

## 自动构建汉化镜像

向 `zh-CN` 分支推送提交后，`.github/workflows/zh-cn-image.yml` 会使用当前源码构建 Linux amd64 镜像，并发布以下标签：

- `ghcr.io/fulaoaz/pentagi:zh-cn`：始终指向最新的汉化版本。
- `ghcr.io/fulaoaz/pentagi:zh-cn-<commit>`：用于锁定和回滚到具体提交。

如果 GHCR 包尚未设为公开，请先在 WSL 中使用具有 `read:packages` 权限的 GitHub 令牌登录，然后拉取并通过 Compose 启动：

```bash
gh auth token | docker login ghcr.io -u YOUR_GITHUB_USER --password-stdin
docker pull ghcr.io/fulaoaz/pentagi:zh-cn
PENTAGI_IMAGE=ghcr.io/fulaoaz/pentagi:zh-cn docker compose up -d
```

上游同步 PR 合并到 `zh-CN` 后会自动触发同一镜像工作流，因此本地部署只需再次执行 `docker compose pull pentagi` 和 `docker compose up -d pentagi`。

界面英文基线记录尚未汉化的已知文案。新增或消失的条目都会使检查失败：

```bash
cd frontend
pnpm i18n:check
```

完成一批汉化并逐项确认后，更新基线：

```bash
pnpm i18n:baseline
```

提供商测试标签门禁会读取后端测试注册表，确保设置页显示的固定测试名称都有中文映射；上游新增测试时需要同步补充标签：

```bash
pnpm exec vitest run src/pages/settings/provider-test-labels.test.ts
```

安装器门禁会检查集中语言文件、安装器启动入口以及直接写入 TUI 和进度终端的英文句子。技术标识有明确放行清单；新增英文必须先人工判断是否需要汉化：

```bash
cd backend
go test ./cmd/installer/wizard/locale
```

API 错误响应门禁会解析 REST 稳定错误目录，确保每个错误码都有简体中文文案，并验证 GraphQL 的 HTTP/WebSocket 语言传递与中文错误响应。服务端内部错误和日志仍保持英文；只有客户端发送 `Accept-Language: zh-CN` 时，REST 响应中的 `msg` 和 GraphQL 响应中的用户提示才会切换为中文：

```bash
cd backend
go test ./pkg/server/response
go test ./pkg/server/services -run '^TestLocalizedGraphQLErrorPresenter$'
```
